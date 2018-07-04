package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/pagination"

	"github.com/sapcc/gophercloud-limes/limes"
	"github.com/sapcc/gophercloud-limes/limes/v1/projects"
)

func main() {
	provider := NewAuthenticatedClient()
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

func NewAuthenticatedClient() *gophercloud.ProviderClient {
	authOpts := &tokens.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		Username:         os.Getenv("OS_USERNAME"),
		Password:         os.Getenv("OS_PASSWORD"),
		DomainName:       os.Getenv("OS_USER_DOMAIN_NAME"),
		AllowReauth:      true,
		Scope: tokens.Scope{
			ProjectName: os.Getenv("OS_PROJECT_NAME"),
			DomainName:  os.Getenv("OS_PROJECT_DOMAIN_NAME"),
		},
	}

	provider, err := openstack.NewClient(os.Getenv("OS_AUTH_URL"))
	if err != nil {
		log.Fatalf("could not initialize openstack client: %v", err)
	}

	provider.UseTokenLock()

	err = openstack.AuthenticateV3(provider, authOpts, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("could not authenticat provider client: %v", err)
	}

	return provider
}

func NewIdentity(provider *gophercloud.ProviderClient) *gophercloud.ServiceClient {
	identity, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("could not initialize identity client: %v", err)
	}
	return identity
}

func NewLimes(provider *gophercloud.ProviderClient) *gophercloud.ServiceClient {
	limesClient, err := limes.NewLimesV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("could not initialize Limes client: %v", err)
	}
	return limesClient
}
