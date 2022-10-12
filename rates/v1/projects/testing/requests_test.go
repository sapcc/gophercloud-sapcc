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

package testing

import (
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesrates "github.com/sapcc/go-api-declarations/limes/rates"

	"github.com/sapcc/gophercloud-sapcc/rates/v1/projects"
)

func TestListProjectsRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ReadOpts{}).ExtractProjects()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limesrates.WindowMinutes
	expected := []limesrates.ProjectReport{
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-berlin",
				Name:       "berlin",
				ParentUUID: "uuid-for-germany",
			},
			Services: limesrates.ProjectServiceReports{
				"shared": &limesrates.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Rates: limesrates.ProjectRateReports{
						"some_action": &limesrates.ProjectRateReport{
							RateInfo: limesrates.RateInfo{
								Name: "some_action",
								Unit: limes.UnitBytes,
							},
							Limit:         5,
							Window:        &rateWindow,
							UsageAsBigint: "1069298",
						},
					},
					ScrapedAt: p2time(24),
				},
				"unshared": &limesrates.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Rates: limesrates.ProjectRateReports{
						"service/something/action:update/removeFloatingIp": &limesrates.ProjectRateReport{
							RateInfo: limesrates.RateInfo{
								Name: "service/something/action:update/removeFloatingIp",
							},
							Limit:  2,
							Window: &rateWindow,
						},
					},
					ScrapedAt: p2time(24),
				},
			},
		},
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-dresden",
				Name:       "dresden",
				ParentUUID: "uuid-for-berlin",
			},
			Services: limesrates.ProjectServiceReports{
				"shared": &limesrates.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Rates: limesrates.ProjectRateReports{},
				},
				"unshared": &limesrates.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Rates: limesrates.ProjectRateReports{
						"service/something:create": &limesrates.ProjectRateReport{
							RateInfo: limesrates.RateInfo{
								Name: "service/something:create",
							},
							Limit:         5,
							Window:        &rateWindow,
							UsageAsBigint: "1069298",
						},
						"service/something/action:update/addFloatingIp": &limesrates.ProjectRateReport{
							RateInfo: limesrates.RateInfo{
								Name: "service/something/action:update/addFloatingIp",
							},
							Limit:  2,
							Window: &rateWindow,
						},
					},
					ScrapedAt: p2time(35),
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListProjectsFilteredRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ReadOpts{
		Services: []string{"shared"},
	}).ExtractProjects()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limesrates.WindowMinutes
	expected := []limesrates.ProjectReport{
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-berlin",
				Name:       "berlin",
				ParentUUID: "uuid-for-germany",
			},
			Services: limesrates.ProjectServiceReports{
				"shared": &limesrates.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Rates: limesrates.ProjectRateReports{
						"some_action": &limesrates.ProjectRateReport{
							RateInfo: limesrates.RateInfo{
								Name: "some_action",
								Unit: limes.UnitBytes,
							},
							Limit:         5,
							Window:        &rateWindow,
							UsageAsBigint: "1069298",
						},
					},
					ScrapedAt: p2time(24),
				},
			},
		},
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-dresden",
				Name:       "dresden",
				ParentUUID: "uuid-for-berlin",
			},
			Services: limesrates.ProjectServiceReports{
				"shared": &limesrates.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Rates: limesrates.ProjectRateReports{},
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.ReadOpts{}).Extract()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limesrates.WindowMinutes
	expected := &limesrates.ProjectReport{
		ProjectInfo: limes.ProjectInfo{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
		},
		Services: limesrates.ProjectServiceReports{
			"shared": &limesrates.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Rates: limesrates.ProjectRateReports{
					"some_action": &limesrates.ProjectRateReport{
						RateInfo: limesrates.RateInfo{
							Name: "some_action",
							Unit: limes.UnitBytes,
						},
						Limit:         5,
						Window:        &rateWindow,
						UsageAsBigint: "1069298",
					},
				},
				ScrapedAt: p2time(24),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectFilteredRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.ReadOpts{
		Services: []string{"shared"},
	}).Extract()
	th.AssertNoErr(t, err)

	rateWindow := 2 * limesrates.WindowMinutes
	expected := &limesrates.ProjectReport{
		ProjectInfo: limes.ProjectInfo{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
		},
		Services: limesrates.ProjectServiceReports{
			"shared": &limesrates.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Rates: limesrates.ProjectRateReports{
					"some_action": &limesrates.ProjectRateReport{
						RateInfo: limesrates.RateInfo{
							Name: "some_action",
							Unit: limes.UnitBytes,
						},
						Limit:         5,
						Window:        &rateWindow,
						UsageAsBigint: "1069298",
					},
				},
				ScrapedAt: p2time(24),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func p2time(timestamp int64) *limes.UnixEncodedTime {
	t := limes.UnixEncodedTime{Time: time.Unix(timestamp, 0).UTC()}
	return &t
}
