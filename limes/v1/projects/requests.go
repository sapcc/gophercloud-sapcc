package projects

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

type ListOpts struct {
	Detail bool `q:"detail"`
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
