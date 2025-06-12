// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package agents

import "github.com/gophercloud/gophercloud/v2"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("agents")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func initURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("agents", "init")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func tagsURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id, "tags")
}

func deleteTagURL(c *gophercloud.ServiceClient, id, key string) string {
	return c.ServiceURL("agents", id, "tags", key)
}

func factsURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id, "facts")
}
