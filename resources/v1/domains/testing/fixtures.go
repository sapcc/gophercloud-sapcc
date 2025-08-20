// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

//nolint:dupl
package testing

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
)

// HandleListDomainsSuccessfully creates an HTTP handler at `/v1/domains` on the
// test handler mux that responds with a list of (two) domains.
func HandleListDomainsSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fixtureName := "list.json"
		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			fixtureName = "list-filtered.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}

// HandleGetDomainSuccessfully creates an HTTP handler at `/v1/domains/:domain_id` on the
// test handler mux that responds with a single domain.
func HandleGetDomainSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/domains/uuid-for-karachi", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fixtureName := "get.json"
		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			fixtureName = "get-filtered.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}
