package price

import "github.com/gophercloud/gophercloud"

func listURL(c *gophercloud.ServiceClient, opts ListOpts) string {
	if opts.Region == "" {
		return c.ServiceURL("masterdata", "pricelist")
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
