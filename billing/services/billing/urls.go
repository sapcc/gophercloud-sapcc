// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package billing

import "github.com/gophercloud/gophercloud/v2"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("services", "billing")
}
