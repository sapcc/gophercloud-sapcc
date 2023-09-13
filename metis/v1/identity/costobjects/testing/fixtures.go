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
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

// HandleGetCostObjectSuccessfully creates an HTTP handler at `/identity/costobject/:costobject_id` on the
// test handler mux that responds with a single costobject.
func HandleGetCostObjectSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/identity/costobject/costobject-1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", "get.json"))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}

// HandleListCostObjectsSuccessfully creates an HTTP handler at `/identity/costobject` on the
// test handler mux that responds with a list of costobjects.
func HandleListCostObjectsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/identity/costobject", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var jsonBytes []byte
		var err error

		// ensure that the query filters are present
		if !r.URL.Query().Has("domain") && !r.URL.Query().Has("project") {
			t.Fatal("HandleListCostObjectsSuccessfully failed missing domain and project query parameters")
		}

		switch r.URL.Query().Has("limit") && !r.URL.Query().Has("cursor") {
		case true:
			jsonBytes, err = os.ReadFile(filepath.Join("fixtures", "list_with_next.json"))
			th.AssertNoErr(t, err)

			var resp struct {
				APIVersion string         `json:"apiVersion"`
				Data       map[string]any `json:"data"`
			}
			err = json.Unmarshal(jsonBytes, &resp)
			th.AssertNoErr(t, err)
			// adding a nextLink to the pagination response
			resp.Data["nextLink"] = th.Endpoint() + "/identity/costobject?domain=foo&project=bar&cursor=dummycursor&limit=1"
			jsonBytes, err = json.Marshal(resp)
		default:
			jsonBytes, err = os.ReadFile(filepath.Join("fixtures", "list.json"))
		}
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}
