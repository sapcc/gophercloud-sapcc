// Package domains provides interaction with Limes at the domain hierarchical level.
package domains

import (
	"github.com/gophercloud/gophercloud"
	"github.com/sapcc/limes/pkg/api"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToDomainListQuery() (string, error)
}

// ListOpts enables filtering of a List request.
type ListOpts struct {
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToDomainListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToDomainListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the domains to which the current token has access.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) (r CommonResult) {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToDomainListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = c.Get(url, &r.Body, nil)
	return
}

// GetOptsBuilder allows extensions to add additional parameters to the Get request.
type GetOptsBuilder interface {
	ToDomainGetQuery() (string, error)
}

// GetOpts enables filtering of a Get request.
type GetOpts struct {
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToDomainGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToDomainGetQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single domain, by ID.
func Get(c *gophercloud.ServiceClient, domainID string, opts GetOptsBuilder) (r CommonResult) {
	url := getURL(c, domainID)
	if opts != nil {
		query, err := opts.ToDomainGetQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = c.Get(url, &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToDomainUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a domain.
type UpdateOpts struct {
	Services api.ServiceQuotas `json:"services"`
}

// ToDomainUpdateMap formats a UpdateOpts into an Update request.
func (opts UpdateOpts) ToDomainUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "domain")
}

// Update modifies the attributes of a domain.
func Update(c *gophercloud.ServiceClient, domainID string, opts UpdateOptsBuilder) error {
	url := updateURL(c, domainID)
	b, err := opts.ToDomainUpdateMap()
	if err != nil {
		return err
	}
	_, err = c.Put(url, b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
