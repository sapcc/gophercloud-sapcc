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

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToInterconnectionListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	ID                      []string `q:"id"`
	State                   []string `q:"state"`
	ProjectID               []string `q:"project_id"`
	Name                    []string `q:"name"`
	Type                    []string `q:"type"`
	LocalResourceID         []string `q:"local_resource_id"`
	RemoteResourceID        []string `q:"remote_resource_id"`
	RemoteRegion            []string `q:"remote_region"`
	RemoteInterconnectionID []string `q:"remote_interconnection_id"`
	Fields                  []string `q:"fields"`
	Limit                   int      `q:"limit"`
	Marker                  string   `q:"marker"`
}

// ToInterconnectionListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToInterconnectionListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List makes a request against the API to list interconnections accessible to
// you.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToInterconnectionListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := InterconnectionPage{pagination.MarkerPageBase{PageResult: r}}
		p.MarkerPageBase.Owner = p
		return p
	})
}

// Get retrieves a specific interconnection based on its ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToInterconnectionCreateMap() (map[string]any, error)
}

// CreateOpts represents options used to create an interconnection.
type CreateOpts struct {
	Name                    string `json:"name,omitempty"`
	ProjectID               string `json:"project_id,omitempty"`
	Type                    string `json:"type,omitempty"`
	LocalResourceID         string `json:"local_resource_id"`
	RemoteResourceID        string `json:"remote_resource_id"`
	RemoteRegion            string `json:"remote_region"`
	RemoteInterconnectionID string `json:"remote_interconnection_id,omitempty"`
}

// ToInterconnectionCreateMap formats a CreateOpts into a map.
func (opts CreateOpts) ToInterconnectionCreateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "interconnection")
}

// Create creates a new interconnection.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToInterconnectionCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Post(createURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes an interconnection.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Delete(deleteURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToInterconnectionUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update an interconnection.
type UpdateOpts struct {
	Name                    *string `json:"name,omitempty"`
	State                   *string `json:"state,omitempty"`
	RemoteInterconnectionID *string `json:"remote_interconnection_id,omitempty"`
}

// ToInterconnectionUpdateMap formats an UpdateOpts into a map.
func (opts UpdateOpts) ToInterconnectionUpdateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "interconnection")
}

// Update allows interconnections to be updated.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToInterconnectionUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Put(updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
