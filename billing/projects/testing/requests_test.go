package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-billing/billing/projects"
)

var iTrue = true

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
			Inherited: new(bool),
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
		Name:      "1234567890",
		Type:      "IO",
		Inherited: &iTrue,
	},
	ProjectType:                    "quota",
	ResponsiblePrimaryContactID:    "D123456",
	ResponsiblePrimaryContactEmail: "example@mail.com",
	RevenueRelevance:               "generating",
	BusinessCriticality:            "dev",
	NumberOfEndusers:               99,
	AdditionalInformation:          "",
	ChangedBy:                      "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
	ChangedAt:                      time.Date(2019, time.August, 26, 9, 9, 5, 457000000, time.UTC),
	Collector:                      "billing.region.local",
	Region:                         "region",
	IsComplete:                     true,
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
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

	th.Mux.HandleFunc("/projects/e9141fb24eee4b3e9f25ae69cda31132", func(w http.ResponseWriter, r *http.Request) {
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

	th.Mux.HandleFunc("/projects/e9141fb24eee4b3e9f25ae69cda31132", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	revenueRelevance := "generating"
	description := "Demos and Tests"
	responsiblePrimaryContactEmail := "example@mail.com"
	additionalInformation := ""
	projectID := "e9141fb24eee4b3e9f25ae69cda31132"
	domainID := "2bac466eed364d8a92e477459e908736"
	projectName := "project"
	numberOfEndusers := 99
	responsiblePrimaryContactID := "D123456"
	parentID := "2bac466eed364d8a92e477459e908736"
	businessCriticality := "dev"

	options := projects.UpdateOpts{
		RevenueRelevance:               &revenueRelevance,
		Description:                    &description,
		ResponsiblePrimaryContactEmail: &responsiblePrimaryContactEmail,
		CostObject: &projects.CostObject{
			Type:      "IO",
			Name:      "1234567890",
			Inherited: &iTrue,
		},
		AdditionalInformation:       &additionalInformation,
		ProjectID:                   &projectID,
		DomainID:                    &domainID,
		ProjectName:                 &projectName,
		NumberOfEndusers:            &numberOfEndusers,
		ResponsiblePrimaryContactID: &responsiblePrimaryContactID,
		ParentID:                    &parentID,
		BusinessCriticality:         &businessCriticality,
	}

	s, err := projects.Update(fake.ServiceClient(), "e9141fb24eee4b3e9f25ae69cda31132", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *s, updateResponse)
}

func TestUpdateNoCost(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/projects/e9141fb24eee4b3e9f25ae69cda31132", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequestNoCost)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	revenueRelevance := "generating"
	description := "Demos and Tests"
	responsiblePrimaryContactEmail := "example@mail.com"
	additionalInformation := ""
	projectID := "e9141fb24eee4b3e9f25ae69cda31132"
	domainID := "2bac466eed364d8a92e477459e908736"
	projectName := "project"
	numberOfEndusers := 99
	responsiblePrimaryContactID := "D123456"
	parentID := "2bac466eed364d8a92e477459e908736"
	businessCriticality := "dev"

	options := projects.UpdateOpts{
		RevenueRelevance:               &revenueRelevance,
		Description:                    &description,
		ResponsiblePrimaryContactEmail: &responsiblePrimaryContactEmail,
		AdditionalInformation:          &additionalInformation,
		ProjectID:                      &projectID,
		DomainID:                       &domainID,
		ProjectName:                    &projectName,
		NumberOfEndusers:               &numberOfEndusers,
		ResponsiblePrimaryContactID:    &responsiblePrimaryContactID,
		ParentID:                       &parentID,
		BusinessCriticality:            &businessCriticality,
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
