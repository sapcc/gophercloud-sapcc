package clusters

import "github.com/gophercloud/gophercloud"

func getURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("clusters", "current")
}
