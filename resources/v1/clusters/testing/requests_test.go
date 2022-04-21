package testing

import (
	"encoding/json"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"

	"github.com/sapcc/gophercloud-sapcc/resources/v1/clusters"
)

func TestGetCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(fake.ServiceClient(), clusters.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	var cap uint64 = 10
	var scrap int64 = 22
	expected := &limes.ClusterReport{
		ID: "pakistan",
		Services: limes.ClusterServiceReports{
			"shared": &limes.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limes.ClusterResourceReports{
					"stuff": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:     &cap,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
					"things": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Capacity:     &cap,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
				},
				MaxScrapedAt: p2i64(33),
				MinScrapedAt: p2i64(33),
			},
			"unshared": &limes.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "contradiction",
				},
				Resources: limes.ClusterResourceReports{
					"stuff": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:     &cap,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
					"things": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Capacity:     &cap,
						DomainsQuota: p2ui64(5),
						Usage:        2,
					},
				},
				MaxScrapedAt: p2i64(33),
				MinScrapedAt: p2i64(33),
			},
		},
		MaxScrapedAt: &scrap,
		MinScrapedAt: &scrap,
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetFilteredCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(fake.ServiceClient(), clusters.GetOpts{
		Detail:    true,
		Services:  []string{"unshared"},
		Resources: []string{"stuff"},
	}).Extract()
	th.AssertNoErr(t, err)

	var cap uint64 = 10
	var scrap int64 = 22
	expected := &limes.ClusterReport{
		ID: "pakistan",
		Services: limes.ClusterServiceReports{
			"unshared": &limes.ClusterServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "contradiction",
				},
				Resources: limes.ClusterResourceReports{
					"stuff": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:      &cap,
						DomainsQuota:  p2ui64(4),
						Usage:         2,
						Subcapacities: json.RawMessage(`[{"cores":200,"hypervisor":"cluster-1"},{"cores":800,"hypervisor":"cluster-2"}]`),
					},
				},
				MaxScrapedAt: p2i64(33),
				MinScrapedAt: p2i64(33),
			},
		},
		MaxScrapedAt: &scrap,
		MinScrapedAt: &scrap,
	}
	th.CheckDeepEquals(t, expected, actual)
}

func p2i64(x int64) *int64 {
	return &x
}

func p2ui64(x uint64) *uint64 {
	return &x
}
