// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package automations

import "github.com/gophercloud/gophercloud/v2"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("automations", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("automations")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
