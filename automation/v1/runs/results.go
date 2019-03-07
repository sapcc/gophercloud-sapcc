package runs

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

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a run resource.
func (r commonResult) Extract() (*Run, error) {
	var s Run
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Run.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract method
// to interpret it as a Run.
type GetResult struct {
	commonResult
}

// Owner represents a Lyra Run Owner.
type Owner struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DomainID   string `json:"domain_id"`
	DomainName string `json:"domain_name"`
}

// Run represents a Lyra Run.
type Run struct {
	ID                   string      `json:"id"`
	AutomationID         string      `json:"automation_id"`
	AutomationName       string      `json:"automation_name"`
	Selector             string      `json:"selector"`
	RepositoryRevision   string      `json:"repository_revision"`
	AutomationAttributes interface{} `json:"automation_attributes"`
	// State could be: preparing, executing, failed, completed
	State     string    `json:"state"`
	Log       string    `json:"log"`
	Jobs      []string  `json:"jobs"`
	Owner     Owner     `json:"owner"`
	ProjectID string    `json:"project_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (r *Run) UnmarshalJSON(b []byte) error {
	type tmp Run
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339Milli `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339Milli `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Run(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return nil
}

func (r *Run) MarshalJSON() ([]byte, error) {
	type ext struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	type tmp struct {
		ext
		Run
	}

	s := tmp{
		ext{
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
			UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
		},
		Run{
			ID:                   r.ID,
			AutomationID:         r.AutomationID,
			AutomationName:       r.AutomationName,
			Selector:             r.Selector,
			RepositoryRevision:   r.RepositoryRevision,
			AutomationAttributes: r.AutomationAttributes,
			State:                r.State,
			Log:                  r.Log,
			Jobs:                 r.Jobs,
			Owner:                r.Owner,
			ProjectID:            r.ProjectID,
		},
	}

	return json.Marshal(s)
}

// RunPage is the page returned by a pager when traversing over a collection of
// runs.
type RunPage struct {
	pagination.MarkerPageBase
}

// NextPageURL is invoked when a paginated collection of runs has reached the
// end of a page and the pager seeks to traverse over a new one. In order to do
// this, it needs to construct the next page's URL.
func (r RunPage) NextPageURL() (string, error) {
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
func (r RunPage) LastMarker() (string, error) {
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

// IsEmpty checks whether a RunPage struct is empty.
func (r RunPage) IsEmpty() (bool, error) {
	runs, err := ExtractRuns(r)
	return len(runs) == 0, err
}

// ExtractRuns accepts a Page struct, specifically a RunPage struct,
// and extracts the elements into a slice of Run structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractRuns(r pagination.Page) ([]Run, error) {
	var s []Run
	err := ExtractRunsInto(r, &s)
	return s, err
}

func ExtractRunsInto(r pagination.Page, v interface{}) error {
	return r.(RunPage).Result.ExtractIntoSlicePtr(v, "")
}
