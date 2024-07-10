// Copyright 2023 SAP SE
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

package costobjects

import (
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"

	v1 "github.com/sapcc/gophercloud-sapcc/v2/metis/v1"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToCostObjectListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing costobjects.
type ListOpts struct {
	// Project will only return costobjects for the specified project uuid.
	Project string `q:"project"`
	// Domain will only return costobjects for the specified domain uuid.
	Domain string `q:"domain"`
	// Limit will limit the number of results returned per page.
	Limit int `q:"limit"`
	// UUIDs will only return costobjects with the specified UUIDs.
	UUIDs []string `q:"uuids"`
}

// ToCostObjectListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCostObjectListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// costobjects.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	serviceURL := listURL(client)
	if opts != nil {
		query, err := opts.ToCostObjectListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		serviceURL += query
	}
	return pagination.NewPager(client, serviceURL, v1.CreatePage())
}

// Get retrieves a specific costobject based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
