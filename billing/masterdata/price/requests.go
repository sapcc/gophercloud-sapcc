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

package price

import (
	"errors"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOpts is a structure that holds options for listing prices.
type ListOpts struct {
	OnlyActive bool      `q:"onlyActive"`
	MetricType string    `q:"METRIC_TYPE"`
	Region     string    `q:"-"`
	From       time.Time `q:"-"`
	To         time.Time `q:"-"`
}

// ToPriceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPriceListQuery() (string, error) {
	if opts.Region != "" && opts.MetricType == "" {
		return "", errors.New("option MetricType is required, when Region is set")
	}

	if (opts.From != (time.Time{}) || opts.To != (time.Time{})) && opts.Region == "" {
		return "", errors.New("option Region is required, when From or To are set")
	}

	if opts.OnlyActive && opts.Region != "" {
		return "", errors.New("cannot use OnlyActive, when Region is set")
	}

	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}

	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// price.
func List(c *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(c, opts)
	if opts != (ListOpts{}) {
		query, err := opts.ToPriceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PricePage{pagination.SinglePageBase(r)}
	})
}
