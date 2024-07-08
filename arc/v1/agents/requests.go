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

package agents

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
	ToAgentListQuery() (string, error)
}

// ListOpts allows the filtering of paginated collections through the API.
// Filtering is achieved by passing in filter value. Page and PerPage are used
// for pagination.
type ListOpts struct {
	Page    int `q:"page"`
	PerPage int `q:"per_page"`
	// E.g. '@os = "darwin" OR (landscape = "staging" AND pool = "green")'
	// where:
	// @fact - fact
	// tag - tag
	Filter string `q:"q"`
}

// ToAgentListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAgentListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// agents. It accepts a ListOpts struct, which allows you to filter the
// returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToAgentListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := AgentPage{pagination.MarkerPageBase{PageResult: r}}
		p.MarkerPageBase.Owner = p
		return p
	})
}

// Get retrieves a specific agent based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// InitOptsBuilder allows extensions to add additional parameters to the
// Init request.
type InitOptsBuilder interface {
	ToAgentInitMap() (map[string]string, error)
}

// InitOpts represents the attributes used when initializing a new agent.
type InitOpts struct {
	// Valid options:
	// * application/json
	// * text/x-shellscript
	// * text/x-powershellscript
	// * text/cloud-config
	Accept string `h:"Accept" required:"true"`
}

// ToAgentInitMap formats a InitOpts into a map of headers.
func (opts InitOpts) ToAgentInitMap() (map[string]string, error) {
	return gophercloud.BuildHeaders(opts)
}

// Init accepts an InitOpts struct and initializes a new agent using the values
// provided.
func Init(ctx context.Context, c *gophercloud.ServiceClient, opts InitOptsBuilder) (r InitResult) {
	h, err := opts.ToAgentInitMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Request(ctx, http.MethodPost, initURL(c), &gophercloud.RequestOpts{
		MoreHeaders:      h,
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

// Delete accepts a unique ID and deletes the agent associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Delete(ctx, deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateTagsBuilder allows extensions to add additional parameters to
// the CreateTags request.
type CreateTagsBuilder interface {
	ToTagsCreateMap() (map[string]string, error)
}

// ToTagsCreateMap converts a Tags into a request body.
func (opts Tags) ToTagsCreateMap() (map[string]string, error) {
	return opts, nil
}

// CreateTags adds/updates tags for a given agent.
func CreateTags(ctx context.Context, client *gophercloud.ServiceClient, agentID string, opts Tags) (r TagsErrResult) {
	b, err := opts.ToTagsCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := client.Post(ctx, tagsURL(client, agentID), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetTags lists tags for a given agent.
func GetTags(ctx context.Context, client *gophercloud.ServiceClient, agentID string) (r TagsResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := client.Get(ctx, tagsURL(client, agentID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteTag deletes an individual tag from an agent.
func DeleteTag(ctx context.Context, client *gophercloud.ServiceClient, agentID, key string) (r TagsErrResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := client.Delete(ctx, deleteTagURL(client, agentID, key), &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetFacts lists tags for a given agent.
func GetFacts(ctx context.Context, client *gophercloud.ServiceClient, agentID string) (r FactsResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := client.Get(ctx, factsURL(client, agentID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
