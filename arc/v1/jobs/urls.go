package jobs

import "github.com/gophercloud/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("jobs", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("jobs")
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

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func getLogURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("jobs", id, "log")
}
