package testing

import (
	"fmt"
	"net/http"
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

		if r.URL.Query().Get("detail") != "" {
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
											"usage": 2,
											"subresources": [
												{
													"id": "thirdthing",
													"value": 5
												},
												{
													"id": "fourththing",
													"value": 123
												}
											]
										}
									],
									"scraped_at": 22
								}
							]
						}
					]
				}`)
			return
		}

		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")
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
											"name": "things",
											"quota": 10,
											"usage": 2
										}
									],
									"scraped_at": 22
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
											"name": "things",
											"quota": 10,
											"usage": 2
										}
									],
									"scraped_at": 44
								}
							]
						}
					]
				}`)
			return
		}

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

// HandleGetProjectSuccessfully creates an HTTP handler at `/domains/:domain_id/projects/:project_id` on the
// test handler mux that responds with a single project.
func HandleGetProjectSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/uuid-for-germany/projects/uuid-for-berlin", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if r.URL.Query().Get("detail") != "" {
			fmt.Fprintf(w, `
				{
					"project": {
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
										"usage": 2,
										"subresources": [
											{
												"id": "thirdthing",
												"value": 5
											},
											{
												"id": "fourththing",
												"value": 123
											}
										]
									}
								],
								"scraped_at": 22
							}
						]
					}
				}
			`)
			return
		}

		if (r.URL.Query().Get("service") == "shared" || r.URL.Query().Get("area") == "shared") &&
			r.URL.Query().Get("resource") == "things" {
			th.TestHeader(t, r, "X-Limes-Cluster-Id", "fakecluster")
			fmt.Fprintf(w, `
				{
					"project": {
						"id": "uuid-for-berlin",
						"name": "berlin",
						"parent_id": "uuid-for-germany",
						"services": [
							{
								"type": "shared",
								"area": "shared",
								"resources": [
									{
										"name": "things",
										"quota": 10,
										"usage": 2
									}
								],
								"scraped_at": 22
							}
						]
					}
				}
			`)
			return
		}

		fmt.Fprintf(w, `
			{
				"project": {
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
						}
					]
				}
			}
		`)
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
