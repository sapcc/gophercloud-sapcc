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

package billing

import (
	"net/url"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToBillingListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing billings.
type ListOpts struct {
	CostObject       string    `q:"cost_object"`
	ProjectID        string    `q:"project_id"`
	ExcludeMissingCO bool      `q:"exclude_missing_co"`
	Format           string    `q:"format"`
	Language         string    `q:"language"`
	Year             int       `q:"year"`
	Month            int       `q:"month"`
	From             time.Time `q:"-"`
	To               time.Time `q:"-"`
}

// ToBillingListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToBillingListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	params := q.Query()

	if opts.From != (time.Time{}) {
		params.Add("from", opts.From.Format(gophercloud.RFC3339MilliNoZ))
	}

	if opts.To != (time.Time{}) {
		params.Add("to", opts.To.Format(gophercloud.RFC3339MilliNoZ))
	}

	q = &url.URL{RawQuery: params.Encode()}

	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// billing.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	serviceURL := listURL(c)
	if opts != nil {
		query, err := opts.ToBillingListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		serviceURL += query
	}
	return pagination.NewPager(c, serviceURL, func(r pagination.PageResult) pagination.Page {
		return BillingPage{pagination.SinglePageBase(r)}
	})
}
