// Copyright 2024 SAP SE
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

package interconnections

import (
	"net/http"
	"net/url"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

const (
	invalidMarker = "-1"
)

// commonResult is an internal base type for all operation results.
type commonResult struct {
	gophercloud.Result
}

// Extract will extract a response object from a result.
func (r commonResult) Extract() (*Interconnection, error) {
	var s Interconnection
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v any) error {
	return r.Result.ExtractIntoStructPtr(v, "interconnection")
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	commonResult
}

// Interconnection represents the structure of an interconnection object.
type Interconnection struct {
	ID                      string     `json:"id"`
	Name                    string     `json:"name"`
	ProjectID               string     `json:"project_id"`
	TenantID                string     `json:"tenant_id"`
	Type                    string     `json:"type"`
	State                   string     `json:"state"`
	LocalResourceID         string     `json:"local_resource_id"`
	RemoteResourceID        string     `json:"remote_resource_id"`
	RemoteRegion            string     `json:"remote_region"`
	RemoteInterconnectionID string     `json:"remote_interconnection_id"`
	LocalParameters         Parameters `json:"local_parameters"`
	RemoteParameters        Parameters `json:"remote_parameters"`
}

// Parameter represents the structure of a parameter object.
type Parameters struct {
	ProjectID []string `json:"project_id"`
}

// InterconnectionPage
type InterconnectionPage struct {
	pagination.MarkerPageBase
}

// NextPageURL is invoked when a paginated collection of interconnections has
// reached the end of a page and the pager seeks to traverse to the next page.
func (r InterconnectionPage) NextPageURL() (string, error) {
	currentURL := r.URL
	mark, err := r.Owner.LastMarker()
	if err != nil {
		return "", err
	}
	if mark == invalidMarker {
		return "", nil
	}

	q := currentURL.Query()
	q.Set("marker", mark)
	currentURL.RawQuery = q.Encode()
	return currentURL.String(), nil
}

// LastMarker returns the last offset in a ListResult.
func (r InterconnectionPage) LastMarker() (string, error) {
	results, err := ExtractInterconnections(r)
	if err != nil {
		return invalidMarker, err
	}
	if len(results) == 0 {
		return invalidMarker, nil
	}

	u, err := url.Parse(r.URL.String())
	if err != nil {
		return invalidMarker, err
	}
	queryParams := u.Query()
	limit := queryParams.Get("limit")

	// Limit is not present, only one page required
	if limit == "" {
		return invalidMarker, nil
	}

	return results[len(results)-1].ID, nil
}

// IsEmpty checks whether a Interconnection struct is empty.
func (r InterconnectionPage) IsEmpty() (bool, error) {
	if r.StatusCode == http.StatusNoContent {
		return true, nil
	}

	is, err := ExtractInterconnections(r)
	return len(is) == 0, err
}

// ExtractInterconnections accepts a Page struct, specifically a
// InterconnectionPage struct, and extracts the elements into a slice of
// Interconnection structs. In other words, a generic collection is mapped into
// a relevant slice.
func ExtractInterconnections(r pagination.Page) ([]Interconnection, error) {
	var s []Interconnection
	err := ExtractInterconnectionssInto(r, &s)
	return s, err
}

func ExtractInterconnectionssInto(r pagination.Page, v any) error {
	return r.(InterconnectionPage).Result.ExtractIntoSlicePtr(v, "interconnections")
}
