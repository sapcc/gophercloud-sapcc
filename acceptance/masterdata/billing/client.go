package billing

import (
	"net/http"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/sapcc/gophercloud-billing/billing"
)

// configureDebug will configure the provider client to print the API
// requests and responses if OS_DEBUG is enabled.
func configureDebug(client *gophercloud.ProviderClient) *gophercloud.ProviderClient {
	if os.Getenv("OS_DEBUG") != "" {
		client.HTTPClient = http.Client{
			Transport: &clients.LogRoundTripper{
				Rt: &http.Transport{},
			},
		}
	}

	return client
}

// NewBillingClient returns a *ServiceClient for making calls
// to the OpenStack Billing API. An error will be returned if
// authentication or client creation was not possible.
func NewBillingClient() (*gophercloud.ServiceClient, error) {
	ao, err := clientconfig.AuthOptions(nil)
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(*ao)
	if err != nil {
		return nil, err
	}

	client = configureDebug(client)

	return billing.NewBilling(client, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
}
