package testing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

// HandleListProjectsSuccessfully creates an HTTP handler at `/domains/:domain_id/projects` on the
// test handler mux that responds with a list of (two) projects.
func HandleListProjectsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var fixtureName string
		switch {
		case r.URL.Query().Get("detail") != "":
			fixtureName = "list-details.json"
		case (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things":

			th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")
			fixtureName = "list-filtered.json"
		default:
			fixtureName = "list.json"
		}

		jsonBytes, err := ioutil.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		fmt.Fprint(w, string(jsonBytes))
	})
}

// HandleGetProjectSuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id` on the
// test handler mux that responds with a single project.
func HandleGetProjectSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-berlin", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var fixtureName string
		switch {
		case r.URL.Query().Get("detail") != "":
			fixtureName = "get-details.json"
		case (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things":

			th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")
			fixtureName = "get-filtered.json"
		default:
			fixtureName = "get.json"
		}

		jsonBytes, err := ioutil.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		fmt.Fprint(w, string(jsonBytes))
	})
}

// HandleUpdateProjectSuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id` on the
// test handler mux that tests project updates.
func HandleUpdateProjectSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-berlin", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")

		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleSyncProjectSuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id/sync` on the
// test handler mux that syncs a project.
func HandleSyncProjectSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-dresden/sync", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")

		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleUpdateProjectPartly creates an HTTP handler at `/domains/:domain_id/projects/:project_id` on the
// test handler mux that tests partly project updates.
func HandleUpdateProjectPartly(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-berlin", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `it is currently not allowed to set bursting.enabled and quotas in the same request`)
	})
}

// HandleUpdateProjectUnsuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id` on the
// test handler mux that tests an unsuccessful project updates.
func HandleUpdateProjectUnsuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-berlin", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `it is currently not allowed to set bursting.enabled and quotas in the same request`)
	})
}
