package projects

import "github.com/gophercloud/gophercloud"

func listURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("domains", id, "projects")
}
