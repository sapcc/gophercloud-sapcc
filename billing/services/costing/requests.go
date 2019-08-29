package costing

import (
	"net/url"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToCostingListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing costings.
type ListOpts struct {
	CostObject        string    `q:"cost_object"`
	ProjectID         string    `q:"project_id"`
	DomainID          string    `q:"domain_id"`
	Service           string    `q:"service"`
	Measure           string    `q:"measure"`
	ExcludeInternalCO bool      `q:"exclude_internal_co"`
	Format            string    `q:"format"`
	Language          string    `q:"language"`
	Last              int       `q:"last"`
	Start             time.Time `q:"-"`
	End               time.Time `q:"-"`
}

// ToCostingListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCostingListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	params := q.Query()

	if opts.Start != (time.Time{}) {
		params.Add("start", opts.Start.Format(gophercloud.RFC3339MilliNoZ))
	}

	if opts.End != (time.Time{}) {
		params.Add("end", opts.End.Format(gophercloud.RFC3339MilliNoZ))
	}

	q = &url.URL{RawQuery: params.Encode()}

	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// costing clusters.
func ListCluster(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	return list(c, opts, listURL(c, "cluster"))
}

// List returns a Pager which allows you to iterate over a collection of
// costing domains.
func ListDomains(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	return list(c, opts, listURL(c, "domains"))
}

// List returns a Pager which allows you to iterate over a collection of
// costing projects.
func ListProjects(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	return list(c, opts, listURL(c, "projects"))
}

// List returns a Pager which allows you to iterate over a collection of
// costing objects.
func ListObjects(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	return list(c, opts, listURL(c, "objects"))
}

func list(c *gophercloud.ServiceClient, opts ListOptsBuilder, url string) pagination.Pager {
	if opts != nil {
		query, err := opts.ToCostingListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return CostingPage{pagination.SinglePageBase(r)}
	})
}
