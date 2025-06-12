// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package billing

import (
	"context"
	"net/http"
	"os"

	"github.com/sapcc/gophercloud-sapcc/v2/internal/acceptance/clients"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/utils/v2/openstack/clientconfig"

	cc_clients "github.com/sapcc/gophercloud-sapcc/v2/clients"
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
func NewBillingClient(ctx context.Context) (*gophercloud.ServiceClient, error) {
	ao, err := clientconfig.AuthOptions(nil)
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ctx, *ao)
	if err != nil {
		return nil, err
	}

	client = configureDebug(client)

	return cc_clients.NewBilling(client, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
}
