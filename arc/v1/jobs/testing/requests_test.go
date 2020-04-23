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
	"github.com/sapcc/gophercloud-sapcc/arc/v1/jobs"
)

var jobsList = []jobs.Job{
	{
		Version:   1,
		Sender:    "linux",
		RequestID: "afa7a1d0-12a0-4848-ae4d-7bb7b01f126d",
		To:        "fe3da83f-919e-4b2e-8200-7acb1816c8d0",
		Timeout:   3600,
		Agent:     "execute",
		Action:    "tarball",
		Payload:   "{\"path\":\"/\",\"environment\":{\"X\":\"y\"},\"url\":\"https://objectstore:443/v1/AUTH_3946cfbc1fda4ce19561da1df5443c86/path/to\"}",
		Status:    "failed",
		CreatedAt: time.Date(2018, time.April, 29, 11, 29, 51, 321560000, time.UTC),
		UpdatedAt: time.Date(2018, time.April, 29, 11, 29, 51, 385979000, time.UTC),
		Project:   "3946cfbc1fda4ce19561da1df5443c86",
		User: jobs.User{
			DomainID:   "123",
			DomainName: "domain",
			ID:         "123",
			Name:       "user",
			Roles:      []string{"automation_admin"},
		},
	},
	{
		Version:   1,
		Sender:    "linux",
		RequestID: "c6c5e3a4-9a6a-40c5-a46b-cc8f2482e3df",
		To:        "88e5cad3-38e6-454f-b412-662cda03e7a1",
		Timeout:   60,
		Agent:     "execute",
		Action:    "script",
		Payload:   "echo \"Scritp start\"\n\nfor i in {1..10}\ndo\n\techo $i\n  sleep 1s\ndone\n\necho \"Script done\"",
		Status:    "complete",
		CreatedAt: time.Date(2019, time.March, 6, 13, 17, 13, 823592000, time.UTC),
		UpdatedAt: time.Date(2019, time.March, 6, 13, 17, 23, 853638000, time.UTC),
		Project:   "3946cfbc1fda4ce19561da1df5443c86",
		User: jobs.User{
			DomainID:   "123",
			DomainName: "domain",
			ID:         "123",
			Name:       "user",
			Roles:      []string{"automation_admin"},
		},
	},
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	count := 0

	jobs.List(fake.ServiceClient(), jobs.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := jobs.ExtractJobs(page)
		if err != nil {
			t.Errorf("Failed to extract jobs: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, jobsList, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/jobs/88e5cad3-38e6-454f-b412-662cda03e7a1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := jobs.Get(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, jobsList[1])
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, CreateResponse)
	})

	options := jobs.CreateOpts{
		To:      "7ec336bd-fcd1-42af-a663-da578dd0b224",
		Timeout: 60,
		Agent:   "execute",
		Action:  "script",
		Payload: "echo \"Script start\"\n\nfor i in {1..10}\ndo\n\techo $i\n  sleep 1s\ndone\n\necho \"Script done\"",
	}
	_, err := jobs.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
}

func TestGetLog(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/jobs/7ec336bd-fcd1-42af-a663-da578dd0b224/log", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, LogResponse)
	})

	response := jobs.GetLog(fake.ServiceClient(), "7ec336bd-fcd1-42af-a663-da578dd0b224")
	th.AssertNoErr(t, response.Err)

	expectedHeader := &jobs.GetLogHeader{ContentType: "text/plain; charset=utf-8"}

	headers, err := response.ExtractHeaders()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, *expectedHeader, *headers)

	n, err := response.ExtractContent()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, LogResponse, string(n))
}

func TestRequiredCreateOpts(t *testing.T) {
	res := jobs.Create(fake.ServiceClient(), jobs.CreateOpts{})
	if res.Err == nil || !strings.Contains(fmt.Sprintf("%s", res.Err), "Missing input for argument") {
		t.Fatalf("Expected error, got none")
	}
}
