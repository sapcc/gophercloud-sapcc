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

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"

	"github.com/sapcc/gophercloud-sapcc/resources/v1/clusters"
)

func TestGetClusterRates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(fake.ServiceClient(), clusters.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	expected := &limes.ClusterReport{
		ID: "current",
		Services: limes.ClusterServiceReports{
			"shared": &limes.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limes.ClusterResourceReports{},
				Rates: limes.ClusterRateLimitReports{
					"service/shared/objects:create": &limes.ClusterRateLimitReport{
						RateInfo: limes.RateInfo{
							Name: "service/shared/objects:create",
						},
						Limit:  5000,
						Window: limes.MustParseWindow("1s"),
					},
				},
				MaxRatesScrapedAt: p2i64(45),
				MinRatesScrapedAt: p2i64(23),
			},
			"unshared": &limes.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "unshared",
				},
				Resources:         limes.ClusterResourceReports{},
				MaxRatesScrapedAt: p2i64(34),
				MinRatesScrapedAt: p2i64(12),
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
		Services: []string{"shared"},
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &limes.ClusterReport{
		ID: "current",
		Services: limes.ClusterServiceReports{
			"shared": &limes.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limes.ClusterResourceReports{},
				Rates: limes.ClusterRateLimitReports{
					"service/shared/objects:create": &limes.ClusterRateLimitReport{
						RateInfo: limes.RateInfo{
							Name: "service/shared/objects:create",
						},
						Limit:  5000,
						Window: limes.MustParseWindow("1s"),
					},
				},
				MaxRatesScrapedAt: p2i64(45),
				MinRatesScrapedAt: p2i64(23),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func p2i64(x int64) *int64 {
	return &x
}
