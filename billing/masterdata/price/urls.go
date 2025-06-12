// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package price

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

func listURL(c *gophercloud.ServiceClient, opts ListOpts) string {
	if opts.Region == "" {
		return c.ServiceURL("masterdata", "pricelist")
	}

	if opts.To.Equal((time.Time{})) {
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
