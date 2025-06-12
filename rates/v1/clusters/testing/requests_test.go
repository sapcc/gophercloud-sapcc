// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fake "github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesrates "github.com/sapcc/go-api-declarations/limes/rates"

	"github.com/sapcc/gophercloud-sapcc/v2/rates/v1/clusters"
)

func TestGetClusterRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(t.Context(), fake.ServiceClient(), clusters.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesrates.ClusterReport{
		ClusterInfo: limes.ClusterInfo{
			ID: "current",
		},
		Services: limesrates.ClusterServiceReports{
			"shared": &limesrates.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Rates: limesrates.ClusterRateReports{
					"service/shared/objects:create": &limesrates.ClusterRateReport{
						RateInfo: limesrates.RateInfo{
							Name: "service/shared/objects:create",
						},
						Limit:  5000,
						Window: limesrates.MustParseWindow("1s"),
					},
				},
				MaxScrapedAt: p2time(45),
				MinScrapedAt: p2time(23),
			},
			"unshared": &limesrates.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "unshared",
				},
				MaxScrapedAt: p2time(34),
				MinScrapedAt: p2time(12),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetFilteredClusterRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(t.Context(), fake.ServiceClient(), clusters.GetOpts{
		Services: []limes.ServiceType{"shared"},
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesrates.ClusterReport{
		ClusterInfo: limes.ClusterInfo{
			ID: "current",
		},
		Services: limesrates.ClusterServiceReports{
			"shared": &limesrates.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Rates: limesrates.ClusterRateReports{
					"service/shared/objects:create": &limesrates.ClusterRateReport{
						RateInfo: limesrates.RateInfo{
							Name: "service/shared/objects:create",
						},
						Limit:  5000,
						Window: limesrates.MustParseWindow("1s"),
					},
				},
				MaxScrapedAt: p2time(45),
				MinScrapedAt: p2time(23),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func p2time(timestamp int64) *limes.UnixEncodedTime {
	t := limes.UnixEncodedTime{Time: time.Unix(timestamp, 0).UTC()}
	return &t
}
