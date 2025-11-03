// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package admin

import "github.com/gophercloud/gophercloud/v2"

func resourceScrapeErrorsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("admin", "resource-scrape-errors")
}

func assetScrapeErrorsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("admin", "asset-scrape-errors")
}

func assetResizeErrorsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("admin", "asset-resize-errors")
}
