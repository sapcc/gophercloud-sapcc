package agents

import "github.com/gophercloud/gophercloud"

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

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
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

func deleteTagURL(c *gophercloud.ServiceClient, id string, key string) string {
	return c.ServiceURL("agents", id, "tags", key)
}

func factsURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id, "facts")
}
