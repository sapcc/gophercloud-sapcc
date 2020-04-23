package agents

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

const (
	invalidMarker = "-1"
)

type GetResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts an agent resource.
func (r GetResult) Extract() (*Agent, error) {
	var s Agent
	err := r.ExtractInto(&s)
	return &s, err
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// InitHeader represents the headers returned in the response from an Init
// request.
type InitHeader struct {
	ContentType string `json:"Content-Type"`
}

// InitJSON represents the structure returned in the response from an Init
// request.
type InitJSON struct {
	Token        string `json:"token"`
	URL          string `json:"url"`
	EndpointURL  string `json:"endpoint_url"`
	UpdateURL    string `json:"update_url"`
	RenewCertURL string `json:"renew_cert_url"`
}

// InitResult represents the result of an init operation. Call its
// ExtractHeaders method to interpret it as an init agent response headers or
// ExtractContent method to interpret it as a response content.
type InitResult struct {
	gophercloud.HeaderResult
	Body []byte
}

// ExtractHeaders will return a struct of headers returned from a call to Init.
func (r InitResult) ExtractHeaders() (*InitHeader, error) {
	var s *InitHeader
	err := r.ExtractInto(&s)
	return s, err
}

// ExtractContent is a function that takes a InitResult's io.Reader body
// and reads all available data into a slice of bytes. Please be aware that due
// the nature of io.Reader is forward-only - meaning that it can only be read
// once and not rewound. You can recreate a reader from the output of this
// function by using bytes.NewReader(initBytes)
func (r *InitResult) ExtractContent() ([]byte, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	return r.Body, nil
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Agent represents an Arc Agent.
type Agent struct {
	DisplayName  string                 `json:"display_name"`
	AgentID      string                 `json:"agent_id"`
	Project      string                 `json:"project"`
	Organization string                 `json:"organization"`
	Facts        map[string]interface{} `json:"facts"`
	Tags         map[string]string      `json:"tags"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	UpdatedWith  string                 `json:"updated_with"`
	UpdatedBy    string                 `json:"updated_by"`
}

func (r *Agent) UnmarshalJSON(b []byte) error {
	type tmp Agent
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339Milli `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339Milli `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Agent(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return nil
}

func (r *Agent) MarshalJSON() ([]byte, error) {
	type ext struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	type tmp struct {
		ext
		Agent
	}

	s := tmp{
		ext{
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
			UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
		},
		Agent{
			DisplayName:  r.DisplayName,
			AgentID:      r.AgentID,
			Project:      r.Project,
			Organization: r.Organization,
			Facts:        r.Facts,
			Tags:         r.Tags,
			UpdatedWith:  r.UpdatedWith,
			UpdatedBy:    r.UpdatedBy,
		},
	}

	return json.Marshal(s)
}

// AgentPage is the page returned by a pager when traversing over a collection
// of agents.
type AgentPage struct {
	pagination.MarkerPageBase
}

// NextPageURL is invoked when a paginated collection of agents has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r AgentPage) NextPageURL() (string, error) {
	currentURL := r.URL
	mark, err := r.Owner.LastMarker()
	if err != nil {
		return "", err
	}
	if mark == invalidMarker {
		return "", nil
	}

	q := currentURL.Query()
	q.Set("page", mark)
	currentURL.RawQuery = q.Encode()
	return currentURL.String(), nil
}

// LastMarker returns the next page in a ListResult.
func (r AgentPage) LastMarker() (string, error) {
	totalPages := -1
	currentPage := -1
	var err error

	page := r.URL.Query().Get("page")
	if page == "" {
		currentPage = 1
	} else {
		currentPage, err = strconv.Atoi(page)
		if err != nil {
			return invalidMarker, err
		}
		if currentPage < 1 {
			currentPage = 1
		}
	}

	if pages, ok := r.Header["Pagination-Pages"]; ok {
		for _, p := range pages {
			totalPages, err = strconv.Atoi(p)
			if err != nil {
				return invalidMarker, err
			}
			break
		}
	}

	if currentPage >= totalPages {
		return invalidMarker, nil
	}

	return strconv.Itoa(currentPage + 1), nil
}

// IsEmpty checks whether a AgentPage struct is empty.
func (r AgentPage) IsEmpty() (bool, error) {
	agents, err := ExtractAgents(r)
	return len(agents) == 0, err
}

// ExtractAgents accepts a Page struct, specifically an AgentPage struct,
// and extracts the elements into a slice of Agent structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractAgents(r pagination.Page) ([]Agent, error) {
	var s []Agent
	err := ExtractAgentsInto(r, &s)
	return s, err
}

func ExtractAgentsInto(r pagination.Page, v interface{}) error {
	return r.(AgentPage).Result.ExtractIntoSlicePtr(v, "")
}

type Tags map[string]string

// TagsResult is the result of a tags request. Call its Extract method
// to interpret it as a map[string]string.
type TagsResult struct {
	gophercloud.Result
}

// TagsDeleteResult is the response from a tags Delete operation. Call
// its ExtractErr to determine if the request succeeded or failed.
type TagsErrResult struct {
	gophercloud.ErrResult
}

// Extract interprets any TagsResult as map[string]string.
func (r TagsResult) Extract() (map[string]string, error) {
	var s Tags
	err := r.ExtractInto(&s)
	return s, err
}

type Facts map[string]interface{}

// FactsResult is the result of a facts request. Call its Extract method
// to interpret it as a map[string]string.
type FactsResult struct {
	gophercloud.Result
}

// Extract interprets any FactsResult as map[string]string.
func (r FactsResult) Extract() (map[string]interface{}, error) {
	var s Facts
	err := r.ExtractInto(&s)
	return s, err
}
