// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package domains

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request
type ListOptsBuilder interface {
	ToDomainListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing domain masterdata.
type ListOpts struct {
	CheckCOValidity bool      `q:"checkCOValidity"`
	ExcludeDeleted  bool      `q:"excludeDeleted"`
	From            time.Time `q:"-"`
}

// ToDomainListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToDomainListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	params := q.Query()

	if opts.From != (time.Time{}) {
		params.Add("from", opts.From.Format(gophercloud.RFC3339MilliNoZ))
	}

	q = &url.URL{RawQuery: params.Encode()}

	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// domains.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	serviceURL := listURL(c)
	if opts != (ListOpts{}) {
		query, err := opts.ToDomainListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		serviceURL += query
	}
	return pagination.NewPager(c, serviceURL, func(r pagination.PageResult) pagination.Page {
		return DomainPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific domain based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToDomainUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents the attributes used when updating an existing
// domain.
type UpdateOpts struct {
	// ID of the domain
	DomainID string `json:"domain_id,omitempty"`
	// Name of the domain
	DomainName string `json:"domain_name,omitempty"`
	// Description of the domain
	Description string `json:"description"`
	// SAP-User-Id of primary contact for the domain
	ResponsiblePrimaryContactID string `json:"responsible_primary_contact_id"`
	// Email-address of primary contact for the domain
	ResponsiblePrimaryContactEmail string `json:"responsible_primary_contact_email"`
	// Freetext field for additional information for domain
	AdditionalInformation string `json:"additional_information"`
	// The cost object structure
	CostObject CostObject `json:"cost_object" required:"true"`
	// Collector of the domain
	Collector string `json:"collector"`
	// Region of the domain
	Region string `json:"region"`
}

// ToDomainUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToDomainUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update accepts a UpdateOpts struct and updates an existing domain using
// the values provided.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToDomainUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Put(ctx, updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func DomainToUpdateOpts(domain *Domain) UpdateOpts {
	return UpdateOpts{
		DomainID:                       domain.DomainID,
		DomainName:                     domain.DomainName,
		Description:                    domain.Description,
		ResponsiblePrimaryContactID:    domain.ResponsiblePrimaryContactID,
		ResponsiblePrimaryContactEmail: domain.ResponsiblePrimaryContactEmail,
		AdditionalInformation:          domain.AdditionalInformation,
		CostObject:                     domain.CostObject,
		Collector:                      domain.Collector,
		Region:                         domain.Region,
	}
}
