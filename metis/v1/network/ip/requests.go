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

package ip

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"

	v1 "github.com/sapcc/gophercloud-sapcc/metis/v1"
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
func Get(c *gophercloud.ServiceClient, ipaddress string) (r GetResult) {
	resp, err := c.Get(getURL(c, ipaddress), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
