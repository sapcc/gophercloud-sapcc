package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-limes/resources/v1/clusters"
	"github.com/sapcc/limes/pkg/api"
	"github.com/sapcc/limes/pkg/limes"
	"github.com/sapcc/limes/pkg/reports"
)

func TestListCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListClustersSuccessfully(t)

	actual, err := clusters.List(fake.ServiceClient(), clusters.ListOpts{}).ExtractClusters()
	th.AssertNoErr(t, err)

	var cap uint64 = 10
	var scrap int64 = 22
	expected := []reports.Cluster{
		{
			ID: "pakistan",
			Services: reports.ClusterServices{
				"shared": &reports.ClusterService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.ClusterResources{
						"stuff": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "stuff",
								Unit: limes.UnitBytes,
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
						"things": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: 33,
					MinScrapedAt: 33,
				},
				"unshared": &reports.ClusterService{
					Shared: true,
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "contradiction",
					},
					Resources: reports.ClusterResources{
						"stuff": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "stuff",
								Unit: limes.UnitBytes,
							},
							Capacity:     &cap,
							Comment:      "tasty tests are so tasty",
							DomainsQuota: 5,
							Usage:        2,
						},
						"things": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: 33,
					MinScrapedAt: 33,
				},
			},
			MaxScrapedAt: &scrap,
			MinScrapedAt: &scrap,
		},
		{
			ID: "germany",
			Services: reports.ClusterServices{
				"shared": &reports.ClusterService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.ClusterResources{
						"stuff": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "stuff",
								Unit: limes.UnitBytes,
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: 33,
					MinScrapedAt: 33,
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
	expected := []reports.Cluster{
		{
			ID: "pakistan",
			Services: reports.ClusterServices{
				"shared": &reports.ClusterService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.ClusterResources{
						"stuff": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "stuff",
								Unit: limes.UnitBytes,
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: 33,
					MinScrapedAt: 33,
				},
			},
			MaxScrapedAt: &scrap,
			MinScrapedAt: &scrap,
		},
		{
			ID: "germany",
			Services: reports.ClusterServices{
				"shared": &reports.ClusterService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.ClusterResources{
						"stuff": &reports.ClusterResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "stuff",
								Unit: limes.UnitBytes,
							},
							Capacity:     &cap,
							DomainsQuota: 5,
							Usage:        2,
						},
					},
					MaxScrapedAt: 33,
					MinScrapedAt: 33,
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
	expected := &reports.Cluster{
		ID: "pakistan",
		Services: reports.ClusterServices{
			"shared": &reports.ClusterService{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: reports.ClusterResources{
					"stuff": &reports.ClusterResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:     &cap,
						DomainsQuota: 5,
						Usage:        2,
					},
					"things": &reports.ClusterResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Capacity:     &cap,
						DomainsQuota: 5,
						Usage:        2,
					},
				},
				MaxScrapedAt: 33,
				MinScrapedAt: 33,
			},
			"unshared": &reports.ClusterService{
				Shared: true,
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "contradiction",
				},
				Resources: reports.ClusterResources{
					"stuff": &reports.ClusterResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "stuff",
							Unit: limes.UnitBytes,
						},
						Capacity:     &cap,
						Comment:      "tasty tests are so tasty",
						DomainsQuota: 5,
						Usage:        2,
					},
					"things": &reports.ClusterResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Capacity:     &cap,
						DomainsQuota: 5,
						Usage:        2,
					},
				},
				MaxScrapedAt: 33,
				MinScrapedAt: 33,
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
	expected := &reports.Cluster{
		ID: "pakistan",
		Services: reports.ClusterServices{
			"unshared": &reports.ClusterService{
				Shared: true,
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "contradiction",
				},
				Resources: reports.ClusterResources{
					"stuff": &reports.ClusterResource{
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
				MaxScrapedAt: 33,
				MinScrapedAt: 33,
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
		Services: []api.ServiceCapacities{
			{Type: "shared", Resources: []api.ResourceCapacity{
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
