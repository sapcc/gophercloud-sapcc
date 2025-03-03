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
	"encoding/json"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fake "github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesresources "github.com/sapcc/go-api-declarations/limes/resources"

	"github.com/sapcc/gophercloud-sapcc/v2/resources/v1/clusters"
)

func TestGetCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(t.Context(), fake.ServiceClient(), clusters.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	var capacity uint64 = 10
	var scrap int64 = 22
	expected := &limesresources.ClusterReport{
		ClusterInfo: limes.ClusterInfo{
			ID: "pakistan",
		},
		Services: limesresources.ClusterServiceReports{
			"shared": &limesresources.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limesresources.ClusterResourceReports{
					"stuff": &limesresources.ClusterResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:     &capacity,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
					"things": &limesresources.ClusterResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						Capacity:     &capacity,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
				},
				MaxScrapedAt: p2time(33),
				MinScrapedAt: p2time(33),
			},
			"unshared": &limesresources.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "contradiction",
				},
				Resources: limesresources.ClusterResourceReports{
					"stuff": &limesresources.ClusterResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:     &capacity,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
					"things": &limesresources.ClusterResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						Capacity:     &capacity,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
				},
				MaxScrapedAt: p2time(33),
				MinScrapedAt: p2time(33),
			},
		},
		MaxScrapedAt: p2time(scrap),
		MinScrapedAt: p2time(scrap),
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetFilteredCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(t.Context(), fake.ServiceClient(), clusters.GetOpts{
		Detail:    true,
		Services:  []limes.ServiceType{"unshared"},
		Resources: []limesresources.ResourceName{"stuff"},
	}).Extract()
	th.AssertNoErr(t, err)

	var capacity uint64 = 10
	var scrap int64 = 22
	expected := &limesresources.ClusterReport{
		ClusterInfo: limes.ClusterInfo{
			ID: "pakistan",
		},
		Services: limesresources.ClusterServiceReports{
			"unshared": &limesresources.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "contradiction",
				},
				Resources: limesresources.ClusterResourceReports{
					"stuff": &limesresources.ClusterResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:      &capacity,
						DomainsQuota:  p2ui64(4),
						Usage:         2,
						Subcapacities: json.RawMessage(`[{"cores":200,"hypervisor":"cluster-1"},{"cores":800,"hypervisor":"cluster-2"}]`),
					},
				},
				MaxScrapedAt: p2time(33),
				MinScrapedAt: p2time(33),
			},
		},
		MaxScrapedAt: p2time(scrap),
		MinScrapedAt: p2time(scrap),
	}
	th.CheckDeepEquals(t, expected, actual)
}

func p2time(timestamp int64) *limes.UnixEncodedTime {
	t := limes.UnixEncodedTime{Time: time.Unix(timestamp, 0).UTC()}
	return &t
}

func p2ui64(x uint64) *uint64 {
	return &x
}
