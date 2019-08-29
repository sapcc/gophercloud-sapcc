package billing

import (
	"net/url"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToBillingListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing billings.
type ListOpts struct {
	CostObject       string    `q:"cost_object"`
	ProjectID        string    `q:"project_id"`
	ExcludeMissingCO bool      `q:"exclude_missing_co"`
	Format           string    `q:"format"`
	Language         string    `q:"language"`
	Year             int       `q:"year"`
	Month            int       `q:"month"`
	From             time.Time `q:"-"`
	To               time.Time `q:"-"`
}

// ToBillingListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToBillingListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	params := q.Query()

	if opts.From != (time.Time{}) {
		params.Add("from", opts.From.Format(gophercloud.RFC3339MilliNoZ))
	}

	if opts.To != (time.Time{}) {
		params.Add("to", opts.To.Format(gophercloud.RFC3339MilliNoZ))
	}

	q = &url.URL{RawQuery: params.Encode()}

	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// billing.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToBillingListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return BillingPage{pagination.SinglePageBase(r)}
	})
}
