// Copyright 2020 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runs

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"

	"github.com/sapcc/gophercloud-sapcc/util"
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
	currentPage, totalPages, err := util.GetCurrentAndTotalPages(r.MarkerPageBase)
	if err != nil || currentPage >= totalPages {
		return invalidMarker, err
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
