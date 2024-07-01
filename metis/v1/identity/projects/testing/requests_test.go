// Copyright 2023 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testing

import (
	"context"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/metis/v1/identity/projects"
)

func TestGetProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(context.TODO(), fakeclient.ServiceClient(), "project-1").Extract()
	th.AssertNoErr(t, err)

	expected := &projects.Project{
		Name:        "project1",
		UUID:        "project-1",
		Description: "project1 descr",
		DomainName:  "domain1",
		DomainUUID:  "domain-1",
		CBRMasterdata: projects.CBRMasterdata{
			CostObjectName:                  "my-cost-object",
			CostObjectType:                  "IO",
			CostObjectInherited:             true,
			BusinessCriticality:             "test",
			RevenueRelevance:                "true",
			NumberOfEndusers:                0,
			PrimaryContactUserID:            "i001337",
			PrimaryContactEmail:             "max.mustermann@sample.com",
			OperatorUserID:                  "i001337",
			OperatorEmail:                   "max.mustermann@sample.com",
			InventoryRoleUserID:             "i001337",
			InventoryRoleEmail:              "max.mustermann@sample.com",
			InfrastructureCoordinatorUserID: "i001337",
			InfrastructureCoordinatorEmail:  "max.mustermann@sample.com",
			ExternalCertifications: projects.ExternalCertifications{
				ISO:  false,
				PCI:  true,
				SOC1: false,
				SOC2: false,
				C5:   false,
				SOX:  false,
			},
			GPUEnabled:                   false,
			ContainsPIIDPPHR:             true,
			ContainsExternalCustomerData: true,
		},
		Users: []projects.User{
			{
				UUID:        "1234abcd",
				Name:        "Max Mustermann",
				Email:       "max.mustermann@sample.com",
				Description: "muster",
			},
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)
	opts := projects.ListOpts{
		Limit: 1,
	}

	p, err := projects.List(fakeclient.ServiceClient(), opts).AllPages(context.TODO())
	th.AssertNoErr(t, err)
	actual, err := projects.Extract(p)
	th.AssertNoErr(t, err)

	expected := []projects.Project{
		{
			Name:        "project1",
			UUID:        "project-1",
			Description: "project1 descr",
			DomainName:  "domain1",
			DomainUUID:  "domain-1",
			CBRMasterdata: projects.CBRMasterdata{
				CostObjectName:                  "my-cost-object",
				CostObjectType:                  "IO",
				CostObjectInherited:             true,
				BusinessCriticality:             "test",
				RevenueRelevance:                "true",
				NumberOfEndusers:                0,
				PrimaryContactUserID:            "i001337",
				PrimaryContactEmail:             "max.mustermann@sample.com",
				OperatorUserID:                  "i001337",
				OperatorEmail:                   "max.mustermann@sample.com",
				InventoryRoleUserID:             "i001337",
				InventoryRoleEmail:              "max.mustermann@sample.com",
				InfrastructureCoordinatorUserID: "i001337",
				InfrastructureCoordinatorEmail:  "max.mustermann@sample.com",
				ExternalCertifications: projects.ExternalCertifications{
					ISO:  false,
					PCI:  true,
					SOC1: false,
					SOC2: false,
					C5:   false,
					SOX:  false,
				},
				GPUEnabled:                   false,
				ContainsPIIDPPHR:             true,
				ContainsExternalCustomerData: true,
			},
			Users: []projects.User{
				{
					UUID:        "1234abcd",
					Name:        "Max Mustermann",
					Email:       "max.mustermann@sample.com",
					Description: "muster",
				},
			},
		},
		{
			Name:        "project2",
			UUID:        "project-2",
			Description: "project2 descr",
			DomainName:  "domain2",
			DomainUUID:  "domain-2",
			CBRMasterdata: projects.CBRMasterdata{
				CostObjectName:                  "my-cost-object",
				CostObjectType:                  "IO",
				CostObjectInherited:             false,
				BusinessCriticality:             "test",
				RevenueRelevance:                "true",
				NumberOfEndusers:                0,
				PrimaryContactUserID:            "i001337",
				PrimaryContactEmail:             "max.mustermann@sample.com",
				OperatorUserID:                  "i001337",
				OperatorEmail:                   "max.mustermann@sample.com",
				InventoryRoleUserID:             "i001337",
				InventoryRoleEmail:              "max.mustermann@sample.com",
				InfrastructureCoordinatorUserID: "i001337",
				InfrastructureCoordinatorEmail:  "max.mustermann@sample.com",
				ExternalCertifications: projects.ExternalCertifications{
					ISO:  false,
					PCI:  false,
					SOC1: false,
					SOC2: false,
					C5:   false,
					SOX:  false,
				},
				GPUEnabled:                   false,
				ContainsPIIDPPHR:             true,
				ContainsExternalCustomerData: true,
			},
			Users: []projects.User{
				{
					UUID: "abcd1234",
					Name: "TU1234",
				},
			},
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}
