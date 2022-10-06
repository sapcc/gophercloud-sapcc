// Copyright 2022 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:dupl
package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"

	"github.com/sapcc/gophercloud-sapcc/rates/v1/projects"
)

func TestListProjectsRatesOnly(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ReadOpts{}).ExtractProjects()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limes.WindowMinutes
	expected := []limes.ProjectReport{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{},
					Rates: limes.ProjectRateLimitReports{
						"some_action": &limes.ProjectRateLimitReport{
							RateInfo: limes.RateInfo{
								Name: "some_action",
								Unit: limes.UnitBytes,
							},
							Limit:         5,
							Window:        &rateWindow,
							UsageAsBigint: "1069298",
						},
					},
					RatesScrapedAt: p2i64(24),
				},
				"unshared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{},
					Rates: limes.ProjectRateLimitReports{
						"service/something/action:update/removeFloatingIp": &limes.ProjectRateLimitReport{
							RateInfo: limes.RateInfo{
								Name: "service/something/action:update/removeFloatingIp",
							},
							Limit:  2,
							Window: &rateWindow,
						},
					},
					RatesScrapedAt: p2i64(24),
				},
			},
		},
		{
			UUID:       "uuid-for-dresden",
			Name:       "dresden",
			ParentUUID: "uuid-for-berlin",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{},
					Rates:     limes.ProjectRateLimitReports{},
				},
				"unshared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{},
					Rates: limes.ProjectRateLimitReports{
						"service/something:create": &limes.ProjectRateLimitReport{
							RateInfo: limes.RateInfo{
								Name: "service/something:create",
							},
							Limit:         5,
							Window:        &rateWindow,
							UsageAsBigint: "1069298",
						},
						"service/something/action:update/addFloatingIp": &limes.ProjectRateLimitReport{
							RateInfo: limes.RateInfo{
								Name: "service/something/action:update/addFloatingIp",
							},
							Limit:  2,
							Window: &rateWindow,
						},
					},
					RatesScrapedAt: p2i64(24),
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListProjectsFilteredWithRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ReadOpts{
		Services: []string{"shared"},
	}).ExtractProjects()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limes.WindowMinutes
	expected := []limes.ProjectReport{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					Rates: limes.ProjectRateLimitReports{
						"some_action": &limes.ProjectRateLimitReport{
							RateInfo: limes.RateInfo{
								Name: "some_action",
								Unit: limes.UnitBytes,
							},
							Limit:         5,
							Window:        &rateWindow,
							UsageAsBigint: "1069298",
						},
					},
					ScrapedAt:      p2i64(22),
					RatesScrapedAt: p2i64(24),
				},
			},
		},
		{
			UUID:       "uuid-for-dresden",
			Name:       "dresden",
			ParentUUID: "uuid-for-berlin",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					Rates:     limes.ProjectRateLimitReports{},
					ScrapedAt: p2i64(44),
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectRatesOnly(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.ReadOpts{}).Extract()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limes.WindowMinutes
	expected := &limes.ProjectReport{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: limes.ProjectServiceReports{
			"shared": &limes.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limes.ProjectResourceReports{},
				Rates: limes.ProjectRateLimitReports{
					"some_action": &limes.ProjectRateLimitReport{
						RateInfo: limes.RateInfo{
							Name: "some_action",
							Unit: limes.UnitBytes,
						},
						Limit:         5,
						Window:        &rateWindow,
						UsageAsBigint: "1069298",
					},
				},
				RatesScrapedAt: p2i64(24),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectFilteredWithRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.ReadOpts{
		Services: []string{"shared"},
	}).Extract()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limes.WindowMinutes
	expected := &limes.ProjectReport{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: limes.ProjectServiceReports{
			"shared": &limes.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limes.ProjectResourceReports{
					"things": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
				},
				Rates: limes.ProjectRateLimitReports{
					"some_action": &limes.ProjectRateLimitReport{
						RateInfo: limes.RateInfo{
							Name: "some_action",
							Unit: limes.UnitBytes,
						},
						Limit:         5,
						Window:        &rateWindow,
						UsageAsBigint: "1069298",
					},
				},
				ScrapedAt:      p2i64(22),
				RatesScrapedAt: p2i64(24),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func p2i64(x int64) *int64 {
	return &x
}

func p2ui64(x uint64) *uint64 {
	return &x
}
