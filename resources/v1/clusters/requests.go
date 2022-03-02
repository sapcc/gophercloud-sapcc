// Package clusters provides interaction with Limes at the cluster hierarchical level.
package clusters

import (
	"github.com/gophercloud/gophercloud"
)

// GetOptsBuilder allows extensions to add additional parameters to the Get request.
type GetOptsBuilder interface {
	ToClusterGetQuery() (string, error)
}

// GetOpts contains parameters for filtering a Get request.
type GetOpts struct {
	Detail    bool     `q:"detail"`
	Areas     []string `q:"area"`
	Services  []string `q:"service"`
	Resources []string `q:"resource"`
}

// ToClusterGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToClusterGetQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single cluster, by ID.
func Get(c *gophercloud.ServiceClient, opts GetOptsBuilder) (r CommonResult) {
	url := getURL(c)
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
