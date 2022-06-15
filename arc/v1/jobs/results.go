package jobs

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"

	"github.com/sapcc/gophercloud-sapcc/util"
)

const (
	invalidMarker = "-1"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a job resource.
func (r commonResult) Extract() (*Job, error) {
	var s Job
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// InitResult represents the result of a create operation. Call its Extract
// method to interpret it as a Job.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Job.
type GetResult struct {
	commonResult
}

// GetLogHeader represents the headers returned in the response from a
// GetLog request.
type GetLogHeader struct {
	ContentType string `json:"Content-Type"`
}

// InitResult represents the result of a get job log operation. Call its
// ExtractHeaders method to interpret it as an job log response headers or
// ExtractContent method to interpret it as a response content.
type GetLogResult struct {
	gophercloud.HeaderResult
	Body []byte
}

// Extract will return a struct of headers returned from a call to GetLog.
func (r GetLogResult) ExtractHeaders() (*GetLogHeader, error) {
	var s *GetLogHeader
	err := r.ExtractInto(&s)
	return s, err
}

// ExtractContent is a function that takes a GetLogResult's io.Reader body
// and reads all available data into a slice of bytes. Please be aware that due
// the nature of io.Reader is forward-only - meaning that it can only be read
// once and not rewound. You can recreate a reader from the output of this
// function by using bytes.NewReader(initBytes)
func (r *GetLogResult) ExtractContent() ([]byte, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	return r.Body, nil
}

// User represents an Arc Job User.
type User struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	DomainID   string   `json:"domain_id"`
	DomainName string   `json:"domain_name"`
	Roles      []string `json:"roles"`
}

// Job represents an Arc Job.
type Job struct {
	Version int    `json:"version"`
	Sender  string `json:"sender"`
	// RequestID represents the JobID
	RequestID string `json:"request_id"`
	// To represents the AgentID
	To      string `json:"to"`
	Timeout int    `json:"timeout"`
	// agent can be: chef (zero), execute (script, tarball)
	Agent string `json:"agent"`
	// action can be: script, zero, tarball
	Action  string `json:"action"`
	Payload string `json:"payload"`
	// Status can be: queued, executing, failed, complete
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Project   string    `json:"project"`
	User      User      `json:"user"`
}

func (r *Job) UnmarshalJSON(b []byte) error {
	type tmp Job
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339Milli `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339Milli `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Job(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return nil
}

// JobPage is the page returned by a pager when traversing over a collection
// of jobs.
type JobPage struct {
	pagination.MarkerPageBase
}

// NextPageURL is invoked when a paginated collection of jobs has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r JobPage) NextPageURL() (string, error) {
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
func (r JobPage) LastMarker() (string, error) {
	currentPage, totalPages, err := util.GetCurrentAndTotalPages(r.MarkerPageBase)
	if err != nil || currentPage >= totalPages {
		return invalidMarker, err
	}
	return strconv.Itoa(currentPage + 1), nil
}

// IsEmpty checks whether a JobPage struct is empty.
func (r JobPage) IsEmpty() (bool, error) {
	jobs, err := ExtractJobs(r)
	return len(jobs) == 0, err
}

// ExtractJobs accepts a Page struct, specifically a JobPage struct,
// and extracts the elements into a slice of Job structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractJobs(r pagination.Page) ([]Job, error) {
	var s []Job
	err := ExtractJobsInto(r, &s)
	return s, err
}

func ExtractJobsInto(r pagination.Page, v interface{}) error {
	return r.(JobPage).Result.ExtractIntoSlicePtr(v, "")
}
