// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ip

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
	ToIPAddressListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing ipaddresses.
type ListOpts struct {
	// IPAddresses will only return ipaddresses with the specified UUIDs.
	IPAddresses []string `q:"ip"`
	// CIDR will only return ipaddresses with the specified CIDR.
	CIDR []string `q:"cidr"`
	// DomainID will only return ipaddresses with the specified DomainID.
	DomainID string `q:"domain"`
	// ProjectID will only return ipaddresses with the specified ProjectID.
	ProjectID string `q:"project"`
	// Limit will limit the number of results returned per page.
	Limit int `q:"limit"`
}

// ToIPAddressListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToIPAddressListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of ipadresses.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	serviceURL := listURL(client)
	if opts != nil {
		query, err := opts.ToIPAddressListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		serviceURL += query
	}
	return pagination.NewPager(client, serviceURL, v1.CreatePage())
}

// Get retrieves a specific ipaddress.
func Get(ctx context.Context, c *gophercloud.ServiceClient, ipaddress string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, ipaddress), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
