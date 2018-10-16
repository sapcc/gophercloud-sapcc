package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

var domainListJSON = `
	{
		"domains": [
			{
				"id": "uuid-for-karachi",
				"name": "karachi",
				"services": [
					{
						"type": "shared",
						"area": "shared",
						"resources": [
							{
								"name": "capacity",
								"unit": "B",
								"quota": 10,
								"projects_quota": 5,
								"usage": 2
							},
							{
								"name": "things",
								"quota": 10,
								"projects_quota": 5,
								"usage": 2
							}
						],
						"max_scraped_at": 22,
						"min_scraped_at": 22
					},
					{
						"type": "unshared",
						"area": "unshared",
						"resources": [
							{
								"name": "capacity",
								"unit": "B",
								"quota": 55,
								"projects_quota": 25,
								"usage": 10
							},
							{
								"name": "things",
								"quota": 55,
								"projects_quota": 25,
								"usage": 10
							}
						],
						"max_scraped_at": 11,
						"min_scraped_at": 11
					}
				]
			},
			{
				"id": "uuid-for-lahore",
				"name": "lahore",
				"services": [
					{
						"type": "shared",
						"area": "shared",
						"resources": [
							{
								"name": "capacity",
								"unit": "B",
								"quota": 10,
								"projects_quota": 5,
								"usage": 2
							},
							{
								"name": "things",
								"quota": 10,
								"projects_quota": 5,
								"usage": 2
							}
						],
						"max_scraped_at": 22,
						"min_scraped_at": 22
					},
					{
						"type": "unshared",
						"area": "unshared",
						"resources": [
							{
								"name": "capacity",
								"unit": "B",
								"quota": 55,
								"projects_quota": 25,
								"usage": 10,
								"backend_quota": 0,
								"infinite_backend_quota": true
							},
							{
								"name": "things",
								"externally_managed": true,
								"quota": 55,
								"projects_quota": 25,
								"usage": 10
							}
						],
						"max_scraped_at": 11,
						"min_scraped_at": 11
					}
				]
			}
		]
	}
`

var domainFilteredListJSON = `
	{
		"domains": [
			{
				"id": "uuid-for-karachi",
				"name": "karachi",
				"services": [
					{
						"type": "shared",
						"area": "shared",
						"resources": [
							{
								"name": "things",
								"quota": 10,
								"projects_quota": 5,
								"usage": 2
							}
						],
						"max_scraped_at": 22,
						"min_scraped_at": 22
					}
				]
			},
			{
				"id": "uuid-for-lahore",
				"name": "lahore",
				"services": [
					{
						"type": "shared",
						"area": "shared",
						"resources": [
							{
								"name": "things",
								"quota": 10,
								"projects_quota": 5,
								"usage": 2
							}
						],
						"max_scraped_at": 22,
						"min_scraped_at": 22
					}
				]
			}
		]
	}
`

var domainJSON = `
	{
		"domain": {
			"id": "uuid-for-karachi",
			"name": "karachi",
			"services": [
				{
					"type": "shared",
					"area": "shared",
					"resources": [
						{
							"name": "capacity",
							"unit": "B",
							"quota": 10,
							"projects_quota": 5,
							"usage": 2
						},
						{
							"name": "things",
							"quota": 10,
							"projects_quota": 5,
							"usage": 2
						}
					],
					"max_scraped_at": 22,
					"min_scraped_at": 22
				},
				{
					"type": "unshared",
					"area": "unshared",
					"resources": [
						{
							"name": "capacity",
							"unit": "B",
							"quota": 55,
							"projects_quota": 25,
							"usage": 10
						},
						{
							"name": "things",
							"quota": 55,
							"projects_quota": 25,
							"usage": 10
						}
					],
					"max_scraped_at": 11,
					"min_scraped_at": 11
				}
			]
		}
	}
`

var domainFilteredJSON = `
	{
		"domain": {
			"id": "uuid-for-karachi",
			"name": "karachi",
			"services": [
				{
					"type": "shared",
					"area": "shared",
					"resources": [
						{
							"name": "things",
							"quota": 10,
							"projects_quota": 5,
							"usage": 2
						}
					],
					"max_scraped_at": 22,
					"min_scraped_at": 22
				}
			]
		}
	}
`

// HandleListDomainsSuccessfully creates an HTTP handler at `/v1/domains` on the
// test handler mux that responds with a list of (two) domains.
func HandleListDomainsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			fmt.Fprintf(w, domainFilteredListJSON)
		}

		fmt.Fprintf(w, domainListJSON)
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

		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			fmt.Fprintf(w, domainFilteredJSON)
		}

		fmt.Fprintf(w, domainJSON)
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
