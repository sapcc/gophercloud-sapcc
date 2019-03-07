package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/kayrus/gophercloud-lyra/automation/v1/runs"
)

var RunObject = runs.Run{
	ID:                 "1",
	Log:                "Selecting nodes using filter @identity='88e5cad3-38e6-454f-b412-662cda03e7a1':\n88e5cad3-38e6-454f-b412-662cda03e7a1 automation-node\nUsing exiting artifact for revision a7af74be592a4637ae5de390b8e8888022130e63\nScheduled 1 job:\n61915ce7-f719-4b23-a163-cd1132668110\nScheduled 1 job:\n61915ce7-f719-4b23-a163-cd1132668110\n",
	CreatedAt:          time.Date(2019, time.March, 5, 19, 45, 40, 57000000, time.UTC),
	UpdatedAt:          time.Date(2019, time.March, 5, 19, 45, 57, 41000000, time.UTC),
	RepositoryRevision: "a7af74be592a4637ae5de390b8e8888022130e63",
	State:              "completed",
	Jobs:               []string{"61915ce7-f719-4b23-a163-cd1132668110"},
	Owner: runs.Owner{
		ID:         "b81eec56-5db9-49ae-8775-880b75d38a1a",
		Name:       "user",
		DomainID:   "6c2feb1a-1d38-4541-aba4-93ed61f2ccca",
		DomainName: "project",
	},
	AutomationID:   "2",
	AutomationName: "chef",
	Selector:       "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
	AutomationAttributes: map[string]interface{}{
		"name":                "chef",
		"debug":               true,
		"timeout":             float64(3600),
		"run_list":            []string{"recipe[application::app]"},
		"repository":          "https://github.com/org/chef.git",
		"chef_version":        "12.22.5",
		"repository_revision": "master",
	},
}

var CreatedObject = runs.Run{
	ID:        "2",
	CreatedAt: time.Date(2019, time.March, 5, 20, 3, 16, 954000000, time.UTC),
	UpdatedAt: time.Date(2019, time.March, 5, 20, 3, 16, 954000000, time.UTC),
	State:     "preparing",
	Owner: runs.Owner{
		ID:         "b81eec56-5db9-49ae-8775-880b75d38a1a",
		Name:       "user",
		DomainID:   "6c2feb1a-1d38-4541-aba4-93ed61f2ccca",
		DomainName: "project",
	},
	AutomationID:   "1",
	AutomationName: "chef",
	Selector:       "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/runs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	count := 0

	runs.List(fake.ServiceClient(), runs.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := runs.ExtractRuns(page)
		if err != nil {
			t.Errorf("Failed to extract runs: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, RunObject, actual[0])

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/runs/1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := runs.Get(fake.ServiceClient(), "1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, RunObject, *n)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/runs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := runs.CreateOpts{
		AutomationID: "2",
		Selector:     "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
	}
	n, err := runs.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, CreatedObject)
}

func TestRequiredCreateOpts(t *testing.T) {
	res := runs.Create(fake.ServiceClient(), runs.CreateOpts{})
	if res.Err == nil || !strings.Contains(fmt.Sprintf("%s", res.Err), "Missing input for argument") {
		t.Fatalf("Expected error, got none")
	}
}
