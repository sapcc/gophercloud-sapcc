// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package resources

import "github.com/gophercloud/gophercloud/v2"

func resourceURL(c *gophercloud.ServiceClient, projectID, resourceType string) string {
	return c.ServiceURL("projects", projectID, "resources", resourceType)
}

func rootURL(c *gophercloud.ServiceClient, projectID string) string {
	return c.ServiceURL("projects", projectID)
}

func listURL(c *gophercloud.ServiceClient, projectID string) string {
	return rootURL(c, projectID)
}

func getURL(c *gophercloud.ServiceClient, projectID, resourceType string) string {
	return resourceURL(c, projectID, resourceType)
}

func deleteURL(c *gophercloud.ServiceClient, projectID, resourceType string) string {
	return resourceURL(c, projectID, resourceType)
}

func createURL(c *gophercloud.ServiceClient, projectID, resourceType string) string {
	return resourceURL(c, projectID, resourceType)
}
