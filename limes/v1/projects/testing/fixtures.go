package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

// HandleListProjectsSuccessfully creates an HTTP handler at `/v1/domains/:domain_id/projects` on the test handler mux
// that responds with a `List` response.
func HandleListProjectsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/abc123/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
			{
				"projects": [
					{
						"id": "uuid-for-berlin",
						"name": "berlin",
						"parent_id": "uuid-for-germany",
						"services": [
							{
								"type": "shared",
								"area": "shared",
								"resources": [
									{
										"name": "capacity",
										"unit": "B",
										"quota": 10,
										"usage": 2
									},
									{
										"name": "things",
										"quota": 10,
										"usage": 2
									}
								],
								"scraped_at": 22
							},
							{
								"type": "unshared",
								"area": "unshared",
								"resources": [
									{
										"name": "capacity",
										"unit": "B",
										"quota": 10,
										"usage": 2
									},
									{
										"name": "things",
										"quota": 10,
										"usage": 2
									}
								],
								"scraped_at": 11
							}
						]
					},
					{
						"id": "uuid-for-dresden",
						"name": "dresden",
						"parent_id": "uuid-for-berlin",
						"services": [
							{
								"type": "shared",
								"area": "shared",
								"resources": [
									{
										"name": "capacity",
										"unit": "B",
										"quota": 10,
										"usage": 2,
										"backend_quota": 100
									},
									{
										"name": "things",
										"quota": 10,
										"usage": 2
									}
								],
								"scraped_at": 44
							},
							{
								"type": "unshared",
								"area": "unshared",
								"resources": [
									{
										"name": "capacity",
										"unit": "B",
										"quota": 10,
										"usage": 2
									},
									{
										"name": "things",
										"quota": 10,
										"usage": 2
									}
								],
								"scraped_at": 33
							}
						]
					}
				]
			}
    `)
	})
}
