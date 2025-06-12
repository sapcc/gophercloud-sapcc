// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fake "github.com/gophercloud/gophercloud/v2/testhelper/client"
)

// HandleGetIPAddressSuccessfully creates an HTTP handler at `/network/ip/:ipaddress` on the
// test handler mux that responds with a single ipaddress.
func HandleGetIPAddressSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/network/ip/10.216.24.194", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", "get.json"))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}

// HandleListIPAddressesSuccessfully creates an HTTP handler at `/network/ip` on the
// test handler mux that responds with a list of ipaddresses.
func HandleListIPAddressesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/network/ip", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var jsonBytes []byte
		var err error

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
			resp.Data["nextLink"] = th.Endpoint() + "/network/ip?cursor=dummycursor&limit=1"
			jsonBytes, err = json.Marshal(resp)
		default:
			jsonBytes, err = os.ReadFile(filepath.Join("fixtures", "list.json"))
		}
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}
