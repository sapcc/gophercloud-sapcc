// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package jobs

import (
	"context"
	"io"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToJobListQuery() (string, error)
}

// ListOpts allows the filtering of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the job
// attributes you want to see returned. Page and PerPage are used for
// pagination.
type ListOpts struct {
	Page    int    `q:"page"`
	PerPage int    `q:"per_page"`
	AgentID string `q:"agent_id"`
}

// ToJobListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToJobListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// jobs. It accepts a ListOpts struct, which allows you to filter the returned
// collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToJobListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := JobPage{pagination.MarkerPageBase{PageResult: r}}
		p.Owner = p
		return p
	})
}

// Get retrieves a specific job based on its unique ID.
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
	ToJobCreateMap() (map[string]any, error)
}

// CreateOpts represents the attributes used when creating a new job.
type CreateOpts struct {
	// To represents the AgentID
	To string `json:"to" required:"true"`
	// 1-86400
	Timeout int `json:"timeout" required:"true"`
	// agent can be: chef (zero), execute (script, tarball)
	Agent string `json:"agent" required:"true"`
	// action can be: script, zero, tarball
	Action  string `json:"action" required:"true"`
	Payload string `json:"payload" required:"true"`
}

// ToJobCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToJobCreateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and creates a new job using the values
// provided.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToJobCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Post(ctx, createURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetLog retrieves a log for a Job ID.
func GetLog(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetLogResult) {
	resp, err := c.Request(ctx, http.MethodGet, getLogURL(c, id), &gophercloud.RequestOpts{
		OkCodes:          []int{http.StatusOK},
		KeepResponseBody: true,
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	if r.Err != nil {
		return
	}
	defer resp.Body.Close()
	r.Body, r.Err = io.ReadAll(resp.Body)
	return
}
