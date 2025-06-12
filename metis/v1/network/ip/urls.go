// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ip

import "github.com/gophercloud/gophercloud/v2"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("network", "ip")
}

func getURL(c *gophercloud.ServiceClient, ipaddress string) string {
	return c.ServiceURL("network", "ip", ipaddress)
}
