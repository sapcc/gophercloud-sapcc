// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package operations

import "github.com/gophercloud/gophercloud/v2"

func pendingURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("operations", "pending")
}

func recentlyFailedURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("operations", "recently-failed")
}

func recentlySucceededURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("operations", "recently-succeeded")
}

func projectPendingURL(c *gophercloud.ServiceClient, projectID, assetType string) string {
	return c.ServiceURL("projects", projectID, "resources", assetType, "operations", "pending")
}

func projectRecentlyFailedURL(c *gophercloud.ServiceClient, projectID, assetType string) string {
	return c.ServiceURL("projects", projectID, "resources", assetType, "operations", "recently-failed")
}

func projectRecentlySucceededURL(c *gophercloud.ServiceClient, projectID, assetType string) string {
	return c.ServiceURL("projects", projectID, "resources", assetType, "operations", "recently-succeeded")
}
