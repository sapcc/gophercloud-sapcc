// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesrates "github.com/sapcc/go-api-declarations/limes/rates"

	"github.com/sapcc/gophercloud-sapcc/v2/rates/v1/projects"
)

func TestListProjectsRates(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleListProjectsSuccessfully(t, fakeServer)

	actual, err := projects.List(t.Context(), client.ServiceClient(fakeServer), "uuid-for-germany", projects.ReadOpts{}).ExtractProjects()
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
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleListProjectsSuccessfully(t, fakeServer)

	actual, err := projects.List(t.Context(), client.ServiceClient(fakeServer), "uuid-for-germany", projects.ReadOpts{
		Services: []limes.ServiceType{"shared"},
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
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleGetProjectSuccessfully(t, fakeServer)

	actual, err := projects.Get(t.Context(), client.ServiceClient(fakeServer), "uuid-for-germany", "uuid-for-berlin", projects.ReadOpts{}).Extract()
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
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleGetProjectSuccessfully(t, fakeServer)

	actual, err := projects.Get(t.Context(), client.ServiceClient(fakeServer), "uuid-for-germany", "uuid-for-berlin", projects.ReadOpts{
		Services: []limes.ServiceType{"shared"},
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
