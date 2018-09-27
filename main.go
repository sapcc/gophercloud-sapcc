package main

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/utils/openstack/clientconfig"

	"github.com/sapcc/gophercloud-limes/resources"
	"github.com/sapcc/gophercloud-limes/resources/v1/projects"
)

func main() {
	provider, err := clientconfig.AuthenticatedClient(nil)
	if err != nil {
		log.Fatalf("could not initialize openstack client: %v", err)
	}
	limesClient := NewLimes(provider)
	identityClient := NewIdentity(provider)

	project, err := tokens.Get(identityClient, provider.Token()).ExtractProject()
	if err != nil {
		log.Fatalf("could get project from token: %v", err)
	}

	err = projects.List(limesClient, project.Domain.ID, projects.ListOpts{Detail: true}).EachPage(func(page pagination.Page) (bool, error) {
		if list, err := projects.ExtractProjects(page); err != nil {
			return false, err
		} else {
			for _, project := range list {
				fmt.Printf("%+v\n", project.Services)
			}
		}
		return true, nil
	})
	if err != nil {
		log.Fatalf("couldn't get projects: %v", err)
	}
}

func NewIdentity(provider *gophercloud.ProviderClient) *gophercloud.ServiceClient {
	identity, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("could not initialize identity client: %v", err)
	}
	return identity
}

func NewLimes(provider *gophercloud.ProviderClient) *gophercloud.ServiceClient {
	limesClient, err := resources.NewLimesV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("could not initialize Limes client: %v", err)
	}
	return limesClient
}
