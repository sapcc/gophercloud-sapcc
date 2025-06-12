// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

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
