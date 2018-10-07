// Package projects provides interaction with Limes at the project hierarchical level.
package projects

import (
	"github.com/gophercloud/gophercloud"
	"github.com/sapcc/limes/pkg/api"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

// ListOpts enables filtering of a List request.
type ListOpts struct {
	Detail   bool   `q:"detail"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToProjectListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProjectListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the projects in a specific domain.
func List(c *gophercloud.ServiceClient, domainID string, opts ListOptsBuilder) (r CommonResult) {
	url := listURL(c, domainID)
	if opts != nil {
		query, err := opts.ToProjectListQuery()
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
	ToProjectGetQuery() (string, error)
}

// GetOpts enables filtering of a Get request.
type GetOpts struct {
	Detail   bool   `q:"detail"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToProjectGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToProjectGetQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single project, by ID.
func Get(c *gophercloud.ServiceClient, domainID string, projectID string, opts GetOptsBuilder) (r CommonResult) {
	url := getURL(c, domainID, projectID)
	if opts != nil {
		query, err := opts.ToProjectGetQuery()
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
	ToProjectUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a project.
type UpdateOpts struct {
	Services api.ServiceQuotas `json:"services"`
}

// ToProjectUpdateMap formats a UpdateOpts into an Update request.
func (opts UpdateOpts) ToProjectUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "project")
}

// Update modifies the attributes of a project.
func Update(c *gophercloud.ServiceClient, domainID string, projectID string, opts UpdateOptsBuilder) (r CommonResult) {
	url := updateURL(c, domainID, projectID)
	b, err := opts.ToProjectUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(url, b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Sync schedules a sync task that pulls a project's data from the backing services
// into Limes' local database.
func Sync(c *gophercloud.ServiceClient, domainID string, projectID string) error {
	url := syncURL(c, domainID, projectID)

	_, err := c.Post(url, nil, nil, nil)
	return err
}
