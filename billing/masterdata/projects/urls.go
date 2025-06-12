// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package projects

import "github.com/gophercloud/gophercloud/v2"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("masterdata", "projects", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("masterdata", "projects")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
