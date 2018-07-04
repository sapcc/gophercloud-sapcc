package projects

import "github.com/gophercloud/gophercloud"

func listURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("domains", id, "projects")
}

func getURL(client *gophercloud.ServiceClient, domainID, projectID string) string {
	return client.ServiceURL("domains", domainID, "projects", projectID)
}

func putURL(client *gophercloud.ServiceClient, domainID, projectID string) string {
	return client.ServiceURL("domains", domainID, "projects", projectID)
}
