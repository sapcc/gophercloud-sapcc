// Copyright 2020 SAP SE
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
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesrates "github.com/sapcc/go-api-declarations/limes/rates"

	"github.com/sapcc/gophercloud-sapcc/rates/v1/clusters"
)

func TestGetClusterRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(fake.ServiceClient(), clusters.GetOpts{}).Extract()
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

	actual, err := clusters.Get(fake.ServiceClient(), clusters.GetOpts{
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
