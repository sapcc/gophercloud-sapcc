package automation

import (
	"github.com/gophercloud/gophercloud"
)

// NewAutomationV1 creates a ServiceClient that may be used with the v1 automation package.
func NewAutomationV1(client *gophercloud.ProviderClient, endpointOpts gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	endpointOpts.ApplyDefaults("automation")
	url, err := client.EndpointLocator(endpointOpts)
	if err != nil {
		return sc, err
	}

	resourceBase := url + "api/v1/"
	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       url,
		Type:           "automation",
		ResourceBase:   resourceBase,
	}, nil
}
