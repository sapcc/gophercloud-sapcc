package price

import (
	"time"

	"github.com/gophercloud/gophercloud"
)

func listURL(c *gophercloud.ServiceClient, opts ListOpts) string {
	if opts.Region == "" {
		return c.ServiceURL("masterdata", "pricelist")
	}

	if opts.To == (time.Time{}) {
		opts.To = time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)
	}

	return c.ServiceURL(
		"masterdata",
		"price",
		opts.Region,
		opts.MetricType,
		opts.From.Format(gophercloud.RFC3339NoZ),
		opts.To.Format(gophercloud.RFC3339NoZ),
	)
}
