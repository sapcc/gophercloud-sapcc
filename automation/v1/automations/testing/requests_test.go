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
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fake "github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/automation/v1/automations"
)

var automationsList = []automations.Automation{
	{
		ID:                 "1",
		Type:               "Script",
		Name:               "script",
		ProjectID:          "3946cfbc1fda4ce19561da1df5443c86",
		Repository:         "https://github.com/org/script.git",
		RepositoryRevision: "master",
		Timeout:            3600,
		CreatedAt:          time.Date(2018, time.April, 29, 11, 39, 13, 412000000, time.UTC),
		UpdatedAt:          time.Date(2018, time.April, 29, 11, 39, 13, 412000000, time.UTC),
		Path:               "/install_nginx.sh",
		Arguments:          []string{},
		Environment:        map[string]string{"X": "y"},
	},
	{
		ID:                 "2",
		Type:               "Chef",
		Name:               "chef",
		ProjectID:          "3946cfbc1fda4ce19561da1df5443c86",
		Repository:         "https://github.com/org/chef.git",
		RepositoryRevision: "master",
		Timeout:            3600,
		CreatedAt:          time.Date(2018, time.December, 27, 14, 20, 8, 521000000, time.UTC),
		UpdatedAt:          time.Date(2018, time.December, 28, 13, 5, 52, 241000000, time.UTC),
		RunList:            []string{"recipe[application::app]"},
		ChefAttributes:     map[string]interface{}{},
		Debug:              true,
		ChefVersion:        "12.22.5",
	},
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	count := 0

	//nolint:errcheck
	automations.List(fake.ServiceClient(), automations.ListOpts{}).EachPage(context.TODO(), func(ctx context.Context, page pagination.Page) (bool, error) {
		count++
		actual, err := automations.ExtractAutomations(page)
		if err != nil {
			t.Errorf("Failed to extract automations: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, automationsList, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	n, err := automations.Get(context.TODO(), fake.ServiceClient(), "2").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, automationsList[1])
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, CreateResponse)
	})

	options := automations.CreateOpts{
		Type:               "Chef",
		Name:               "chef",
		Repository:         "https://github.com/org/chef.git",
		RepositoryRevision: "master",
		Timeout:            3600,
		RunList: []string{
			"recipe[application::app]",
		},
		Debug:       true,
		ChefVersion: "12.22.5",
	}
	n, err := automations.Create(context.TODO(), fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, automationsList[1])
}

func TestRequiredCreateOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := automations.Create(context.TODO(), fake.ServiceClient(), automations.CreateOpts{})
	if res.Err == nil || !strings.Contains(res.Err.Error(), "Missing input for argument") {
		t.Fatalf("Expected error, got none")
	}
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})

	options := automations.UpdateOpts{Debug: new(bool)}

	s, err := automations.Update(context.TODO(), fake.ServiceClient(), "2", options).Extract()
	th.AssertNoErr(t, err)

	tmp := automationsList[1]
	tmp.Debug = false
	th.AssertDeepEquals(t, *s, tmp)
}

func TestUpdateCreds(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequestCreds)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponseCreds)
	})

	creds := "foobar"
	options := automations.UpdateOpts{RepositoryCredentials: &creds}

	s, err := automations.Update(context.TODO(), fake.ServiceClient(), "2", options).Extract()
	th.AssertNoErr(t, err)

	tmp := automationsList[1]
	tmp.RepositoryAuthenticationEnabled = true
	th.AssertDeepEquals(t, *s, tmp)
}

func TestRemoveRunList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, RemoveRunListRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, RemoveRunListResponse)
	})

	options := automations.UpdateOpts{
		Tags:    map[string]string{},
		RunList: []string{},
	}

	s, err := automations.Update(context.TODO(), fake.ServiceClient(), "2", options).Extract()
	th.AssertNoErr(t, err)

	tmp := automationsList[1]
	tmp.RunList = nil
	th.AssertDeepEquals(t, *s, tmp)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/automations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := automations.Delete(context.TODO(), fake.ServiceClient(), "2")
	th.AssertNoErr(t, res.Err)
}
