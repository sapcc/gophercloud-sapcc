// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package costing

import "github.com/gophercloud/gophercloud/v2"

func listURL(c *gophercloud.ServiceClient, v string) string {
	return c.ServiceURL("services", "costing", v)
}
