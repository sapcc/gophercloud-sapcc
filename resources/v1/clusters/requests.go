// Package clusters provides interaction with Limes at the cluster hierarchical level.
package clusters

import (
	"github.com/gophercloud/gophercloud"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToClusterListQuery() (string, error)
}

// ListOpts contains parameters for filtering a List request
type ListOpts struct {
	Detail   bool   `q:"detail"`
	Local    bool   `q:"local"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToClusterListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToClusterListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the clusters to which the current token has access.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) (r CommonResult) {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToClusterListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	resp, err := c.Get(url, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetOptsBuilder allows extensions to add additional parameters to the Get request.
type GetOptsBuilder interface {
	ToClusterGetQuery() (string, error)
}

// GetOpts contains parameters for filtering a Get request.
type GetOpts struct {
	Detail   bool   `q:"detail"`
	Local    bool   `q:"local"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToClusterGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToClusterGetQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single cluster, by ID.
func Get(c *gophercloud.ServiceClient, clusterID string, opts GetOptsBuilder) (r CommonResult) {
	url := getURL(c, clusterID)
	if opts != nil {
		query, err := opts.ToClusterGetQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	resp, err := c.Get(url, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
