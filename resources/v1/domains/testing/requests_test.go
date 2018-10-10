package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-limes/resources/v1/domains"
	"github.com/sapcc/limes/pkg/api"
	"github.com/sapcc/limes/pkg/limes"
	"github.com/sapcc/limes/pkg/reports"
)

func TestListDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDomainsSuccessfully(t)

	actual, err := domains.List(fake.ServiceClient(), domains.ListOpts{}).ExtractDomains()
	th.AssertNoErr(t, err)

	var backendQ uint64
	infiniteBackendQ := true
	expected := []reports.Domain{
		{
			UUID: "uuid-for-karachi",
			Name: "karachi",
			Services: reports.DomainServices{
				"shared": &reports.DomainService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.DomainResources{
						"capacity": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   10,
							ProjectsQuota: 5,
							Usage:         2,
						},
						"things": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   10,
							ProjectsQuota: 5,
							Usage:         2,
						},
					},
					MaxScrapedAt: 22,
					MinScrapedAt: 22,
				},
				"unshared": &reports.DomainService{
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "unshared",
					},
					Resources: reports.DomainResources{
						"capacity": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   55,
							ProjectsQuota: 25,
							Usage:         10,
						},
						"things": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   55,
							ProjectsQuota: 25,
							Usage:         10,
						},
					},
					MaxScrapedAt: 11,
					MinScrapedAt: 11,
				},
			},
		},
		{
			UUID: "uuid-for-lahore",
			Name: "lahore",
			Services: reports.DomainServices{
				"shared": &reports.DomainService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.DomainResources{
						"capacity": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   10,
							ProjectsQuota: 5,
							Usage:         2,
						},
						"things": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   10,
							ProjectsQuota: 5,
							Usage:         2,
						},
					},
					MaxScrapedAt: 22,
					MinScrapedAt: 22,
				},
				"unshared": &reports.DomainService{
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "unshared",
					},
					Resources: reports.DomainResources{
						"capacity": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:          55,
							ProjectsQuota:        25,
							Usage:                10,
							BackendQuota:         &backendQ,
							InfiniteBackendQuota: &infiniteBackendQ,
						},
						"things": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name:              "things",
								ExternallyManaged: true,
							},
							DomainQuota:   55,
							ProjectsQuota: 25,
							Usage:         10,
						},
					},
					MaxScrapedAt: 11,
					MinScrapedAt: 11,
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListFilteredDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDomainsSuccessfully(t)

	actual, err := domains.List(fake.ServiceClient(), domains.ListOpts{
		Service:  "shared",
		Resource: "things",
	}).ExtractDomains()
	th.AssertNoErr(t, err)

	expected := []reports.Domain{
		{
			UUID: "uuid-for-karachi",
			Name: "karachi",
			Services: reports.DomainServices{
				"shared": &reports.DomainService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.DomainResources{
						"things": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   10,
							ProjectsQuota: 5,
							Usage:         2,
						},
					},
					MaxScrapedAt: 22,
					MinScrapedAt: 22,
				},
			},
		},
		{
			UUID: "uuid-for-lahore",
			Name: "lahore",
			Services: reports.DomainServices{
				"shared": &reports.DomainService{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: reports.DomainResources{
						"things": &reports.DomainResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   10,
							ProjectsQuota: 5,
							Usage:         2,
						},
					},
					MaxScrapedAt: 22,
					MinScrapedAt: 22,
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetDomainSuccessfully(t)

	actual, err := domains.Get(fake.ServiceClient(), "uuid-for-karachi", domains.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	expected := &reports.Domain{
		UUID: "uuid-for-karachi",
		Name: "karachi",
		Services: reports.DomainServices{
			"shared": &reports.DomainService{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: reports.DomainResources{
					"capacity": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   10,
						ProjectsQuota: 5,
						Usage:         2,
					},
					"things": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   10,
						ProjectsQuota: 5,
						Usage:         2,
					},
				},
				MaxScrapedAt: 22,
				MinScrapedAt: 22,
			},
			"unshared": &reports.DomainService{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "unshared",
				},
				Resources: reports.DomainResources{
					"capacity": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   55,
						ProjectsQuota: 25,
						Usage:         10,
					},
					"things": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   55,
						ProjectsQuota: 25,
						Usage:         10,
					},
				},
				MaxScrapedAt: 11,
				MinScrapedAt: 11,
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetDomainFiltered(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetDomainSuccessfully(t)

	actual, err := domains.Get(fake.ServiceClient(), "uuid-for-karachi", domains.GetOpts{
		Service:  "shared",
		Resource: "things",
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &reports.Domain{
		UUID: "uuid-for-karachi",
		Name: "karachi",
		Services: reports.DomainServices{
			"shared": &reports.DomainService{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: reports.DomainResources{
					"things": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   10,
						ProjectsQuota: 5,
						Usage:         2,
					},
				},
				MaxScrapedAt: 22,
				MinScrapedAt: 22,
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestUpdateDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateDomainSuccessfully(t)

	opts := domains.UpdateOpts{
		Services: api.ServiceQuotas{
			"shared": api.ResourceQuotas{
				"things": limes.ValueWithUnit{99, limes.UnitNone},
			},
		},
	}

	actual, err := domains.Update(fake.ServiceClient(), "uuid-for-karachi", opts).Extract()
	th.AssertNoErr(t, err)

	expected := &reports.Domain{
		UUID: "uuid-for-karachi",
		Name: "karachi",
		Services: reports.DomainServices{
			"shared": &reports.DomainService{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: reports.DomainResources{
					"capacity": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   10,
						ProjectsQuota: 5,
						Usage:         2,
					},
					"things": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   99,
						ProjectsQuota: 5,
						Usage:         2,
					},
				},
				MaxScrapedAt: 22,
				MinScrapedAt: 22,
			},
			"unshared": &reports.DomainService{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "unshared",
				},
				Resources: reports.DomainResources{
					"capacity": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   55,
						ProjectsQuota: 25,
						Usage:         10,
					},
					"things": &reports.DomainResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   55,
						ProjectsQuota: 25,
						Usage:         10,
					},
				},
				MaxScrapedAt: 11,
				MinScrapedAt: 11,
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}
