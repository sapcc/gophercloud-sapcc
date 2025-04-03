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
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToRunListQuery() (string, error)
}

// ListOpts allows the listing of paginated collections through the API. Page
// and PerPage are used for pagination.
type ListOpts struct {
	Page    int `q:"page"`
	PerPage int `q:"per_page"`
}

// ToRunListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRunListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// runs. It accepts a ListOpts struct.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToRunListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := RunPage{pagination.MarkerPageBase{PageResult: r}}
		p.Owner = p
		return p
	})
}

// Get retrieves a specific run based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToRunCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents the attributes used when creating a new run.
type CreateOpts struct {
	AutomationID string `json:"automation_id" required:"true"`
	Selector     string `json:"selector" required:"true"`
}

// ToRunCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToRunCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and creates a new run using the values
// provided.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRunCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Post(ctx, createURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
