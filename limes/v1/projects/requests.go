package projects

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/sapcc/limes/pkg/reports"
)

type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

type ListOpts struct {
	Detail   bool   `q:"detail"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProjectListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, domainID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client, domainID)
	if opts != nil {
		query, err := opts.ToProjectListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ProjectPage{pagination.SinglePageBase(r)}
	})
}

type GetOptsBuilder interface {
	ToProjectGetQuery() (string, error)
}

type GetOpts struct {
	Detail bool `q:"detail"`
}

func (opts GetOpts) ToProjectGetQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func Get(c *gophercloud.ServiceClient, domainID string, projectID string, opts GetOptsBuilder) (r GetResult) {
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

type UpdateOptsBuilder interface {
	ToProjectUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Services reports.ProjectServices
}

func (opts UpdateOpts) ToProjectUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "project")
}

func Update(c *gophercloud.ServiceClient, domainID string, projectID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToProjectUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(putURL(c, domainID, projectID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}
