package limes

import "github.com/gophercloud/gophercloud"

func NewLimesV1(client *gophercloud.ProviderClient, endpointOpts gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	endpointOpts.ApplyDefaults("resources")
	endpoint, err := client.EndpointLocator(endpointOpts)
	if err != nil {
		return nil, err
	}

	endpoint += "v1/"

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           "resources",
	}, nil
}
