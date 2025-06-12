// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package interconnections

import "github.com/gophercloud/gophercloud/v2"

const urlBase = "interconnection/interconnections"

// return /v2.0/interconnection/interconnections/{interconnection-id}
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(urlBase, id)
}

// return /v2.0/interconnection/interconnections
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(urlBase)
}

// return /v2.0/interconnection/interconnections/{interconnection-id}
func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/interconnection/interconnections
func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/interconnection/interconnections
func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/interconnection/interconnections/{interconnection-id}
func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/interconnection/interconnections/{interconnection-id}
func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
