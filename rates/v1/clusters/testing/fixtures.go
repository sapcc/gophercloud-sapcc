// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
)

// HandleGetClusterSuccessfully creates an HTTP handler at `/v1/clusters/:cluster_id` on the
// test handler mux that responds with a single cluster.
func HandleGetClusterSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/clusters/current", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fixtureName := "get.json"
		if r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared" {
			fixtureName = "get-filtered.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}
