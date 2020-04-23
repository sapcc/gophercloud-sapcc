package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/kayrus/gophercloud-arc/arc/v1/agents"
)

var agentTags = map[string]string{
	"pool":      "green",
	"landscape": "staging",
}

var agentFacts = map[string]interface{}{
	"agents": map[string]interface{}{
		"chef":    "enabled",
		"execute": "enabled",
		"rpc":     "enabled",
	},
	"os":               "linux",
	"platform":         "ubuntu",
	"platform_family":  "debian",
	"platform_version": "16.04",
	"project":          "3946cfbc1fda4ce19561da1df5443c86",
}

var agentsList = []agents.Agent{
	{
		DisplayName:  "instance1",
		AgentID:      "88e5cad3-38e6-454f-b412-662cda03e7a1",
		Project:      "3946cfbc1fda4ce19561da1df5443c86",
		Organization: "41aac04ce58c428b9ed2262798d0d336",
		CreatedAt:    time.Date(2018, time.March, 13, 10, 44, 27, 432827000, time.UTC),
		UpdatedAt:    time.Date(2019, time.March, 6, 10, 2, 19, 62626000, time.UTC),
		UpdatedWith:  "a76ddf9f-748d-421e-bcd1-c1a6afc922e4",
		UpdatedBy:    "linux",
	},
	{
		DisplayName:  "instance2",
		AgentID:      "7bf82bb6-61a6-4d01-aa50-16e19d1dca99",
		Project:      "3946cfbc1fda4ce19561da1df5443c86",
		Organization: "41aac04ce58c428b9ed2262798d0d336",
		CreatedAt:    time.Date(2018, time.November, 12, 10, 17, 12, 455872000, time.UTC),
		UpdatedAt:    time.Date(2018, time.December, 3, 8, 48, 57, 400890000, time.UTC),
		UpdatedWith:  "40db2bb1-e6b9-4c64-8353-fae5553a0092",
		UpdatedBy:    "linux",
	},
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	count := 0

	agents.List(fake.ServiceClient(), agents.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := agents.ExtractAgents(page)
		if err != nil {
			t.Errorf("Failed to extract agents: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, agentsList, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/88e5cad3-38e6-454f-b412-662cda03e7a1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := agents.Get(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, agentsList[0])
}

func TestInit(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/init", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "text/x-powershellscript")

		w.Header().Add("Content-Type", "text/x-powershellscript")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, InitResponsePowerShell)
	})

	options := agents.InitOpts{
		Accept: "text/x-powershellscript",
	}
	response := agents.Init(fake.ServiceClient(), options)
	th.AssertNoErr(t, response.Err)

	expectedHeader := &agents.InitHeader{ContentType: "text/x-powershellscript"}

	headers, err := response.ExtractHeaders()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, *expectedHeader, *headers)

	n, err := response.ExtractContent()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, InitResponsePowerShell, string(n))
}

func TestInitJSON(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/init", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, InitResponseJson)
	})

	jsonResp := &agents.InitJSON{
		Token:        "4d523051-089f-41ce-aaf7-727fee19c28a",
		URL:          "https://arc.example.com/api/v1/agents/init/4d523051-089f-41ce-aaf7-727fee19c28a",
		EndpointURL:  "tls://arc-broker.example.com:8883",
		UpdateURL:    "https://stable.arc.example.com",
		RenewCertURL: "https://example.com/renew",
	}

	options := agents.InitOpts{
		Accept: "application/json",
	}
	response := agents.Init(fake.ServiceClient(), options)
	th.AssertNoErr(t, response.Err)

	expectedHeader := &agents.InitHeader{ContentType: "application/json"}

	headers, err := response.ExtractHeaders()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, *expectedHeader, *headers)

	n, err := response.ExtractContent()
	th.AssertNoErr(t, err)

	var initJson agents.InitJSON
	err = json.Unmarshal(n, &initJson)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *jsonResp, initJson)
}

// TODO required headers
/*
func TestRequiredInitOpts(t *testing.T) {
	res := agents.Init(fake.ServiceClient(), agents.InitOpts{})
	if res.Err == nil || !strings.Contains(fmt.Sprintf("%s", res.Err), "Missing input for argument") {
		t.Fatalf("Expected error, got none")
	}
}
*/

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/88e5cad3-38e6-454f-b412-662cda03e7a1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := agents.Delete(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/88e5cad3-38e6-454f-b412-662cda03e7a1/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetTagsResponse)
	})

	n, err := agents.GetTags(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, n, agentTags)
}

func TestCreateTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/88e5cad3-38e6-454f-b412-662cda03e7a1/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := agents.CreateTags(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1", agentTags).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestDeleteTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/88e5cad3-38e6-454f-b412-662cda03e7a1/tags/pool", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := agents.DeleteTag(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1", "pool").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetFacts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/agents/88e5cad3-38e6-454f-b412-662cda03e7a1/facts", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetFactsResponse)
	})

	n, err := agents.GetFacts(fake.ServiceClient(), "88e5cad3-38e6-454f-b412-662cda03e7a1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, n, agentFacts)
}
