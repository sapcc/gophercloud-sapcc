package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"

	"github.com/sapcc/gophercloud-sapcc/resources/v1/domains"
)

func TestListDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDomainsSuccessfully(t)

	actual, err := domains.List(fake.ServiceClient(), domains.ListOpts{}).ExtractDomains()
	th.AssertNoErr(t, err)

	var backendQ uint64
	infiniteBackendQ := true
	expected := []limes.DomainReport{
		{
			UUID: "uuid-for-karachi",
			Name: "karachi",
			Services: limes.DomainServiceReports{
				"shared": &limes.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limes.DomainResourceReports{
						"capacity": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
						"things": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2i64(22),
					MinScrapedAt: p2i64(22),
				},
				"unshared": &limes.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "unshared",
					},
					Resources: limes.DomainResourceReports{
						"capacity": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   p2ui64(55),
							ProjectsQuota: p2ui64(25),
							Usage:         10,
						},
						"things": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(55),
							ProjectsQuota: p2ui64(25),
							Usage:         10,
						},
					},
					MaxScrapedAt: p2i64(11),
					MinScrapedAt: p2i64(11),
				},
			},
		},
		{
			UUID: "uuid-for-lahore",
			Name: "lahore",
			Services: limes.DomainServiceReports{
				"shared": &limes.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limes.DomainResourceReports{
						"capacity": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
						"things": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2i64(22),
					MinScrapedAt: p2i64(22),
				},
				"unshared": &limes.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "unshared",
					},
					Resources: limes.DomainResourceReports{
						"capacity": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:          p2ui64(55),
							ProjectsQuota:        p2ui64(25),
							Usage:                10,
							BackendQuota:         &backendQ,
							InfiniteBackendQuota: &infiniteBackendQ,
						},
						"things": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name:              "things",
								ExternallyManaged: true,
							},
							DomainQuota:   p2ui64(55),
							ProjectsQuota: p2ui64(25),
							Usage:         10,
						},
					},
					MaxScrapedAt: p2i64(11),
					MinScrapedAt: p2i64(11),
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
		Services:  []string{"shared"},
		Resources: []string{"things"},
	}).ExtractDomains()
	th.AssertNoErr(t, err)

	expected := []limes.DomainReport{
		{
			UUID: "uuid-for-karachi",
			Name: "karachi",
			Services: limes.DomainServiceReports{
				"shared": &limes.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limes.DomainResourceReports{
						"things": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2i64(22),
					MinScrapedAt: p2i64(22),
				},
			},
		},
		{
			UUID: "uuid-for-lahore",
			Name: "lahore",
			Services: limes.DomainServiceReports{
				"shared": &limes.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limes.DomainResourceReports{
						"things": &limes.DomainResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2i64(22),
					MinScrapedAt: p2i64(22),
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

	expected := &limes.DomainReport{
		UUID: "uuid-for-karachi",
		Name: "karachi",
		Services: limes.DomainServiceReports{
			"shared": &limes.DomainServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limes.DomainResourceReports{
					"capacity": &limes.DomainResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   p2ui64(10),
						ProjectsQuota: p2ui64(5),
						Usage:         2,
					},
					"things": &limes.DomainResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   p2ui64(10),
						ProjectsQuota: p2ui64(5),
						Usage:         2,
					},
				},
				MaxScrapedAt: p2i64(22),
				MinScrapedAt: p2i64(22),
			},
			"unshared": &limes.DomainServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "unshared",
				},
				Resources: limes.DomainResourceReports{
					"capacity": &limes.DomainResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   p2ui64(55),
						ProjectsQuota: p2ui64(25),
						Usage:         10,
					},
					"things": &limes.DomainResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   p2ui64(55),
						ProjectsQuota: p2ui64(25),
						Usage:         10,
					},
				},
				MaxScrapedAt: p2i64(11),
				MinScrapedAt: p2i64(11),
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
		Services:  []string{"shared"},
		Resources: []string{"things"},
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &limes.DomainReport{
		UUID: "uuid-for-karachi",
		Name: "karachi",
		Services: limes.DomainServiceReports{
			"shared": &limes.DomainServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limes.DomainResourceReports{
					"things": &limes.DomainResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   p2ui64(10),
						ProjectsQuota: p2ui64(5),
						Usage:         2,
					},
				},
				MaxScrapedAt: p2i64(22),
				MinScrapedAt: p2i64(22),
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
		Services: limes.QuotaRequest{
			"shared": limes.ServiceQuotaRequest{
				Resources: limes.ResourceQuotaRequest{
					"things": limes.ValueWithUnit{Value: 99, Unit: limes.UnitNone},
				},
			},
		},
	}

	// if update succeeds then a 202 (no error) is returned.
	err := domains.Update(fake.ServiceClient(), "uuid-for-karachi", opts).ExtractErr()
	th.AssertNoErr(t, err)
}

func p2i64(x int64) *int64 {
	return &x
}

func p2ui64(x uint64) *uint64 {
	return &x
}
