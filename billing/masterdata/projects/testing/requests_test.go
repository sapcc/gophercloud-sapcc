package testing

import (
	"encoding/json"
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
	ProjectType:                    "quota",
	ResponsiblePrimaryContactID:    "D123456",
	ResponsiblePrimaryContactEmail: "example@mail.com",
	RevenueRelevance:               "generating",
	BusinessCriticality:            "dev",
	NumberOfEndusers:               99,
	ChangedBy:                      "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
	ChangedAt:                      time.Date(2019, time.August, 26, 9, 9, 5, 457000000, time.UTC),
	Collector:                      "billing.region.local",
	Region:                         "region",
	IsComplete:                     true,
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	allProjects, err := projects.List(fake.ServiceClient()).AllPages()
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

		fmt.Fprintf(w, GetResponse)
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

		fmt.Fprintf(w, UpdateResponse)
	})

	options := projects.UpdateOpts{
		RevenueRelevance:               "generating",
		Description:                    "Demos and Tests",
		ResponsiblePrimaryContactEmail: "example@mail.com",
		CostObject: projects.CostObject{
			Inherited: true,
		},
		ProjectID:                   "e9141fb24eee4b3e9f25ae69cda31132",
		DomainID:                    "2bac466eed364d8a92e477459e908736",
		ProjectName:                 "project",
		NumberOfEndusers:            99,
		ResponsiblePrimaryContactID: "D123456",
		ParentID:                    "2bac466eed364d8a92e477459e908736",
		BusinessCriticality:         "dev",
	}

	s, err := projects.Update(fake.ServiceClient(), "e9141fb24eee4b3e9f25ae69cda31132", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *s, updateResponse)
}

func TestMarshalUnmarshal(t *testing.T) {
	// must be a pointer, becasue UnmarshalJSON uses a pointer receiver
	jj, err := json.Marshal(&projectsList[0])
	th.AssertNoErr(t, err)

	var unmarshalled projects.Project
	err = json.Unmarshal(jj, &unmarshalled)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, projectsList[0], unmarshalled)
}
