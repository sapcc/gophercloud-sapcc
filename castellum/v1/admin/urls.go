// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
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
