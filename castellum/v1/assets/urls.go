// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package assets

import "github.com/gophercloud/gophercloud/v2"

func assetsURL(c *gophercloud.ServiceClient, projectID, assetType string) string {
	return c.ServiceURL("projects", projectID, "assets", assetType)
}

func assetURL(c *gophercloud.ServiceClient, projectID, assetType, assetID string) string {
	return c.ServiceURL("projects", projectID, "assets", assetType, assetID)
}

func errorResolvedURL(c *gophercloud.ServiceClient, projectID, assetType, assetID string) string {
	return c.ServiceURL("projects", projectID, "assets", assetType, assetID, "error-resolved")
}
