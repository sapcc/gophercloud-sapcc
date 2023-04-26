// Copyright 2020 SAP SE
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
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/billing/masterdata/projects"
)

var projectsList = []projects.Project{
	{
		ProjectID:   "e9141fb24eee4b3e9f25ae69cda31132",
		ProjectName: "project",
		Description: "Demos and Tests",
		ParentID:    "2bac466eed364d8a92e477459e908736",
		DomainID:    "2bac466eed364d8a92e477459e908736",
		DomainName:  "domain",
		CostObject: projects.CostObject{
			Name:      "123456789",
			Type:      "IO",
			Inherited: false,
		},
		ProjectType:                    "quota",
		ResponsiblePrimaryContactID:    "D123456",
		ResponsiblePrimaryContactEmail: "example@mail.com",
		RevenueRelevance:               "generating",
		BusinessCriticality:            "dev",
		NumberOfEndusers:               100,
		AdditionalInformation:          "info",
		ChangedBy:                      "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
		ChangedAt:                      time.Date(2019, time.August, 20, 14, 39, 39, 786000000, time.UTC),
		Collector:                      "billing.region.local",
		Region:                         "region",
		IsComplete:                     true,
	},
}

var updateResponse = projects.Project{
	ProjectID:   "e9141fb24eee4b3e9f25ae69cda31132",
	ProjectName: "project",
	Description: "Demos and Tests",
	ParentID:    "2bac466eed364d8a92e477459e908736",
	DomainID:    "2bac466eed364d8a92e477459e908736",
	DomainName:  "domain",
	CostObject: projects.CostObject{
		Inherited: true,
	},
	ProjectType:                               "quota",
	ResponsiblePrimaryContactID:               "D123456",
	ResponsiblePrimaryContactEmail:            "example@mail.com",
	ResponsibleInventoryRoleID:                "D123456",
	ResponsibleInventoryRoleEmail:             "123@mail.com",
	ResponsibleInfrastructureCoordinatorID:    "D123456",
	ResponsibleInfrastructureCoordinatorEmail: "123@mail.com",
	Customer:                     "123ABC",
	GPUEnabled:                   false,
	ContainsPIIDPPHR:             true,
	ContainsExternalCustomerData: false,
	TypeOfData:                   "",
	ExtCertification: &projects.ExtCertification{
		C5:   false,
		ISO:  true,
		PCI:  false,
		SOC1: false,
		SOC2: false,
		SOX:  false,
	},
	RevenueRelevance:    "generating",
	BusinessCriticality: "dev",
	NumberOfEndusers:    99,
	ChangedBy:           "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
	ChangedAt:           time.Date(2019, time.August, 26, 9, 9, 5, 457000000, time.UTC),
	Collector:           "billing.region.local",
	Region:              "region",
	IsComplete:          true,
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	allProjects, err := projects.List(fake.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := projects.ExtractProjects(allProjects)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, projectsList, actual)
}

func TestListOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.CheckEquals(t, r.URL.Query().Get("checkCOValidity"), "true")
		th.CheckEquals(t, r.URL.Query().Get("excludeDeleted"), "true")
		th.CheckEquals(t, r.URL.Query().Get("from"), "2023-04-26T12:31:42.1337")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	listOpts := projects.ListOpts{
		CheckCOValidity: true,
		ExcludeDeleted:  true,
		From:            time.Date(2023, 04, 26, 12, 31, 42, 133700000, time.UTC),
	}

	allProjects, err := projects.ListWithOpts(fake.ServiceClient(), listOpts).AllPages()
	th.AssertNoErr(t, err)

	actual, err := projects.ExtractProjects(allProjects)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, projectsList, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/projects/e9141fb24eee4b3e9f25ae69cda31132", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	n, err := projects.Get(fake.ServiceClient(), "e9141fb24eee4b3e9f25ae69cda31132").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, projectsList[0])
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/projects/e9141fb24eee4b3e9f25ae69cda31132", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})

	options := projects.UpdateOpts{
		RevenueRelevance:                          "generating",
		Description:                               "Demos and Tests",
		ResponsiblePrimaryContactID:               "D123456",
		ResponsiblePrimaryContactEmail:            "example@mail.com",
		ResponsibleInventoryRoleID:                "D123456",
		ResponsibleInventoryRoleEmail:             "123@mail.com",
		ResponsibleInfrastructureCoordinatorID:    "D123456",
		ResponsibleInfrastructureCoordinatorEmail: "123@mail.com",
		CostObject: projects.CostObject{
			Inherited: true,
		},
		ProjectID:                    "e9141fb24eee4b3e9f25ae69cda31132",
		DomainID:                     "2bac466eed364d8a92e477459e908736",
		ProjectName:                  "project",
		NumberOfEndusers:             99,
		Customer:                     "123ABC",
		GPUEnabled:                   false,
		ContainsPIIDPPHR:             true,
		ContainsExternalCustomerData: false,
		TypeOfData:                   "",
		ExtCertification: &projects.ExtCertification{
			ISO: true,
		},
		ParentID:            "2bac466eed364d8a92e477459e908736",
		BusinessCriticality: "dev",
	}

	s, err := projects.Update(fake.ServiceClient(), "e9141fb24eee4b3e9f25ae69cda31132", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *s, updateResponse)
}
