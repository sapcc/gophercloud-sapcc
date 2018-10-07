package main

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
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
		log.Fatalf("could not get project from token: %v", err)
	}

	result := projects.List(limesClient, project.Domain.ID, projects.ListOpts{Detail: true})
	if result.Err != nil {
		log.Fatalf("could not get projects: %v", result.Err)
	}

	projectList, err := result.ExtractProjects()
	if err != nil {
		log.Fatalf("could not get projects: %v", err)
	}
	for _, project := range projectList {
		fmt.Printf("%+v\n", project.Services)
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
