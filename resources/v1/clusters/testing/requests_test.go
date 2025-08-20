// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"encoding/json"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesresources "github.com/sapcc/go-api-declarations/limes/resources"

	"github.com/sapcc/gophercloud-sapcc/v2/resources/v1/clusters"
)

func TestGetCluster(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleGetClusterSuccessfully(t, fakeServer)

	actual, err := clusters.Get(t.Context(), client.ServiceClient(fakeServer), clusters.GetOpts{}).Extract()
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
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleGetClusterSuccessfully(t, fakeServer)

	actual, err := clusters.Get(t.Context(), client.ServiceClient(fakeServer), clusters.GetOpts{
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
