package costing

import "github.com/gophercloud/gophercloud"

func listURL(c *gophercloud.ServiceClient, v string) string {
	return c.ServiceURL("services", "costing", v)
}
