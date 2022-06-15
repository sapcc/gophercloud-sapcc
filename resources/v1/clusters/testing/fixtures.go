package testing

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

// HandleGetClusterSuccessfully creates an HTTP handler at `/v1/clusters/:cluster_id` on the
// test handler mux that responds with a single cluster.
func HandleGetClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/clusters/current", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fixtureName := "get.json"
		if (r.URL.Query().Get("service") == "unshared" || r.URL.Query().Get("area") == "contradiction") &&
			r.URL.Query().Get("resource") == "stuff" && r.URL.Query().Get("detail") != "" {
			fixtureName = "get-filtered.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes) //nolint:errcheck
	})
}
