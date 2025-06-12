// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package projects

import "github.com/gophercloud/gophercloud/v2"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("identity", "project")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("identity", "project", id)
}
