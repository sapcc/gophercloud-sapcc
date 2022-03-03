package testing

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

// HandleListDomainsSuccessfully creates an HTTP handler at `/v1/domains` on the
// test handler mux that responds with a list of (two) domains.
func HandleListDomainsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fixtureName := "list.json"
		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			fixtureName = "list-filtered.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes)
	})
}

// HandleGetDomainSuccessfully creates an HTTP handler at `/v1/domains/:domain_id` on the
// test handler mux that responds with a single domain.
func HandleGetDomainSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-karachi", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fixtureName := "get.json"
		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			fixtureName = "get-filtered.json"
		}

		jsonBytes, err := os.ReadFile(filepath.Join("fixtures", fixtureName))
		th.AssertNoErr(t, err)
		w.Write(jsonBytes)
	})
}

// HandleUpdateDomainSuccessfully creates an HTTP handler at `/v1/domains/:domain_id` on the
// test handler mux that tests domain update.
func HandleUpdateDomainSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-karachi", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.WriteHeader(http.StatusAccepted)
	})
}
