package arc

import (
	"github.com/gophercloud/gophercloud"
)

// NewArcV1 creates a ServiceClient that may be used with the v1 arc package.
func NewArcV1(client *gophercloud.ProviderClient, endpointOpts gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	endpointOpts.ApplyDefaults("arc")
	url, err := client.EndpointLocator(endpointOpts)
	if err != nil {
		return sc, err
	}

	resourceBase := url + "api/v1/"
	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       url,
		Type:           "arc",
		ResourceBase:   resourceBase,
	}, nil
}
