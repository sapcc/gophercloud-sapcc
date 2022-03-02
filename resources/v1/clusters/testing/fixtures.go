package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

var clusterJSON = `
	{
		"cluster": {
			"id": "pakistan",
			"services": [
				{
					"type": "shared",
					"area": "shared",
					"resources": [
						{
							"name": "stuff",
							"capacity": 10,
							"domains_quota": 5,
							"unit": "B",
							"usage": 2
						},
						{
							"name": "things",
							"capacity": 10,
							"domains_quota": 5,
							"usage": 2
						}
					],
					"max_scraped_at": 33,
					"min_scraped_at": 33
				},
				{
					"shared": true,
					"type": "unshared",
					"area": "contradiction",
					"resources": [
						{
							"name": "stuff",
							"capacity": 10,
							"comment": "tasty tests are so tasty",
							"domains_quota": 5,
							"unit": "B",
							"usage": 2
						},
						{
							"name": "things",
							"capacity": 10,
							"domains_quota": 5,
							"usage": 2
						}
					],
					"max_scraped_at": 33,
					"min_scraped_at": 33
				}
			],
			"max_scraped_at": 22,
			"min_scraped_at": 22
		}
	}
`

var clusterFilteredJSON = `
	{
		"cluster": {
			"id": "pakistan",
			"services": [
				{
					"shared": true,
					"type": "unshared",
					"area": "contradiction",
					"resources": [
						{
							"name": "stuff",
							"capacity": 10,
							"comment": "tasty tests are so tasty",
							"domains_quota": 4,
							"unit": "B",
							"usage": 2,
							"subcapacities": [
								{ "hypervisor": "cluster-1", "cores": 200 },
								{ "hypervisor": "cluster-2", "cores": 800 }
							]
						}
					],
					"max_scraped_at": 33,
					"min_scraped_at": 33
				}
			],
			"max_scraped_at": 22,
			"min_scraped_at": 22
		}
	}
`

// HandleGetClusterSuccessfully creates an HTTP handler at `/v1/clusters/:cluster_id` on the
// test handler mux that responds with a single cluster.
func HandleGetClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/clusters/current", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if (r.URL.Query().Get("service") == "unshared" || r.URL.Query().Get("area") == "contradiction") &&
			r.URL.Query().Get("resource") == "stuff" && r.URL.Query().Get("detail") != "" {
			fmt.Fprintf(w, clusterFilteredJSON)
		}

		fmt.Fprintf(w, clusterJSON)
	})
}
