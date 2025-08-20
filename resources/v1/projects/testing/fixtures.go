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

// HandleListProjectsSuccessfully creates an HTTP handler at `/domains/:domain_id/projects` on the
// test handler mux that responds with a list of (two) projects.
func HandleListProjectsSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/domains/uuid-for-germany/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var fixtureName string
		switch {
		case r.URL.Query().Has("detail"):
			fixtureName = "list-details.json"
		case (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things":
			fixtureName = "list-filtered.json"
		default:
			fixtureName = "list.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}

// HandleGetProjectSuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id` on the
// test handler mux that responds with a single project.
func HandleGetProjectSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-berlin", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var fixtureName string
		switch {
		case r.URL.Query().Has("detail"):
			fixtureName = "get-details.json"
		case (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things":
			fixtureName = "get-filtered.json"
		default:
			fixtureName = "get.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}

// HandleSyncProjectSuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id/sync` on the
// test handler mux that syncs a project.
func HandleSyncProjectSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-dresden/sync", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusAccepted)
	})
}
