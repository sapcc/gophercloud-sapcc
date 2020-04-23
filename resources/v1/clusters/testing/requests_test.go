package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-sapcc/resources/v1/clusters"
	"github.com/sapcc/limes"
)

func TestListCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListClustersSuccessfully(t)

	actual, err := clusters.List(fake.ServiceClient(), clusters.ListOpts{}).ExtractClusters()
	th.AssertNoErr(t, err)

	var cap uint64 = 10
	var scrap int64 = 22
	expected := []limes.ClusterReport{
		{
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
							DomainsQuota: 5,
							Usage:        2,
						},
						"things": &limes.ClusterResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: p2i64(33),
					MinScrapedAt: p2i64(33),
				},
				"unshared": &limes.ClusterServiceReport{
					Shared: true,
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
							Comment:      "tasty tests are so tasty",
							DomainsQuota: 5,
							Usage:        2,
						},
						"things": &limes.ClusterResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: p2i64(33),
					MinScrapedAt: p2i64(33),
				},
			},
			MaxScrapedAt: &scrap,
			MinScrapedAt: &scrap,
		},
		{
			ID: "germany",
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
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: p2i64(33),
					MinScrapedAt: p2i64(33),
				},
			},
			MaxScrapedAt: &scrap,
			MinScrapedAt: &scrap,
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListFilteredCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListClustersSuccessfully(t)

	actual, err := clusters.List(fake.ServiceClient(), clusters.ListOpts{
		Service:  "shared",
		Resource: "stuff",
	}).ExtractClusters()
	th.AssertNoErr(t, err)

	var cap uint64 = 10
	var scrap int64 = 22
	expected := []limes.ClusterReport{
		{
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
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: p2i64(33),
					MinScrapedAt: p2i64(33),
				},
			},
			MaxScrapedAt: &scrap,
			MinScrapedAt: &scrap,
		},
		{
			ID: "germany",
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
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: p2i64(33),
					MinScrapedAt: p2i64(33),
				},
			},
			MaxScrapedAt: &scrap,
			MinScrapedAt: &scrap,
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetClusterSuccessfully(t)

	actual, err := clusters.Get(fake.ServiceClient(), "pakistan", clusters.GetOpts{}).Extract()
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
						DomainsQuota: 5,
						Usage:        2,
					},
					"things": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Capacity:     &cap,
						DomainsQuota: 5,
						Usage:        2,
					},
				},
				MaxScrapedAt: p2i64(33),
				MinScrapedAt: p2i64(33),
			},
			"unshared": &limes.ClusterServiceReport{
				Shared: true,
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
						Comment:      "tasty tests are so tasty",
						DomainsQuota: 5,
						Usage:        2,
					},
					"things": &limes.ClusterResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Capacity:     &cap,
						DomainsQuota: 5,
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

	actual, err := clusters.Get(fake.ServiceClient(), "pakistan", clusters.GetOpts{
		Detail:   true,
		Service:  "unshared",
		Resource: "stuff",
	}).Extract()
	th.AssertNoErr(t, err)

	var cap uint64 = 10
	var scrap int64 = 22
	expected := &limes.ClusterReport{
		ID: "pakistan",
		Services: limes.ClusterServiceReports{
			"unshared": &limes.ClusterServiceReport{
				Shared: true,
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
						Comment:       "tasty tests are so tasty",
						DomainsQuota:  4,
						Usage:         2,
						Subcapacities: `[{"cores":200,"hypervisor":"cluster-1"},{"cores":800,"hypervisor":"cluster-2"}]`,
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

func TestUpdateCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateClusterSuccessfully(t)

	unit := limes.UnitBytes
	opts := clusters.UpdateOpts{
		Services: []limes.ServiceCapacityRequest{
			{Type: "shared", Resources: []limes.ResourceCapacityRequest{
				{
					Name:     "stuff",
					Capacity: 99,
					Unit:     &unit,
					Comment:  "I got 99 problems, but a cluster ain't one.",
				},
			}},
		},
	}

	// if update succeeds then a 202 (no error) is returned.
	err := clusters.Update(fake.ServiceClient(), "germany", opts)
	th.AssertNoErr(t, err)
}

func p2i64(x int64) *int64 {
	return &x
}
