// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fake "github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesresources "github.com/sapcc/go-api-declarations/limes/resources"

	"github.com/sapcc/gophercloud-sapcc/v2/resources/v1/domains"
)

func TestListDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDomainsSuccessfully(t)

	actual, err := domains.List(t.Context(), fake.ServiceClient(), domains.ListOpts{}).ExtractDomains()
	th.AssertNoErr(t, err)

	var backendQ uint64
	infiniteBackendQ := true
	expected := []limesresources.DomainReport{
		{
			DomainInfo: limes.DomainInfo{
				UUID: "uuid-for-karachi",
				Name: "karachi",
			},
			Services: limesresources.DomainServiceReports{
				"shared": &limesresources.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limesresources.DomainResourceReports{
						"capacity": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
						"things": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2time(22),
					MinScrapedAt: p2time(22),
				},
				"unshared": &limesresources.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "unshared",
					},
					Resources: limesresources.DomainResourceReports{
						"capacity": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   p2ui64(55),
							ProjectsQuota: p2ui64(25),
							Usage:         10,
						},
						"things": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(55),
							ProjectsQuota: p2ui64(25),
							Usage:         10,
						},
					},
					MaxScrapedAt: p2time(11),
					MinScrapedAt: p2time(11),
				},
			},
		},
		{
			DomainInfo: limes.DomainInfo{
				UUID: "uuid-for-lahore",
				Name: "lahore",
			},
			Services: limesresources.DomainServiceReports{
				"shared": &limesresources.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limesresources.DomainResourceReports{
						"capacity": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
						"things": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2time(22),
					MinScrapedAt: p2time(22),
				},
				"unshared": &limesresources.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "unshared",
						Area: "unshared",
					},
					Resources: limesresources.DomainResourceReports{
						"capacity": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							DomainQuota:          p2ui64(55),
							ProjectsQuota:        p2ui64(25),
							Usage:                10,
							BackendQuota:         &backendQ,
							InfiniteBackendQuota: &infiniteBackendQ,
						},
						"things": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(55),
							ProjectsQuota: p2ui64(25),
							Usage:         10,
						},
					},
					MaxScrapedAt: p2time(11),
					MinScrapedAt: p2time(11),
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

	actual, err := domains.List(t.Context(), fake.ServiceClient(), domains.ListOpts{
		Services:  []limes.ServiceType{"shared"},
		Resources: []limesresources.ResourceName{"things"},
	}).ExtractDomains()
	th.AssertNoErr(t, err)

	expected := []limesresources.DomainReport{
		{
			DomainInfo: limes.DomainInfo{
				UUID: "uuid-for-karachi",
				Name: "karachi",
			},
			Services: limesresources.DomainServiceReports{
				"shared": &limesresources.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limesresources.DomainResourceReports{
						"things": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2time(22),
					MinScrapedAt: p2time(22),
				},
			},
		},
		{
			DomainInfo: limes.DomainInfo{
				UUID: "uuid-for-lahore",
				Name: "lahore",
			},
			Services: limesresources.DomainServiceReports{
				"shared": &limesresources.DomainServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type: "shared",
						Area: "shared",
					},
					Resources: limesresources.DomainResourceReports{
						"things": &limesresources.DomainResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							DomainQuota:   p2ui64(10),
							ProjectsQuota: p2ui64(5),
							Usage:         2,
						},
					},
					MaxScrapedAt: p2time(22),
					MinScrapedAt: p2time(22),
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

	actual, err := domains.Get(t.Context(), fake.ServiceClient(), "uuid-for-karachi", domains.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesresources.DomainReport{
		DomainInfo: limes.DomainInfo{
			UUID: "uuid-for-karachi",
			Name: "karachi",
		},
		Services: limesresources.DomainServiceReports{
			"shared": &limesresources.DomainServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limesresources.DomainResourceReports{
					"capacity": &limesresources.DomainResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   p2ui64(10),
						ProjectsQuota: p2ui64(5),
						Usage:         2,
					},
					"things": &limesresources.DomainResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   p2ui64(10),
						ProjectsQuota: p2ui64(5),
						Usage:         2,
					},
				},
				MaxScrapedAt: p2time(22),
				MinScrapedAt: p2time(22),
			},
			"unshared": &limesresources.DomainServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "unshared",
					Area: "unshared",
				},
				Resources: limesresources.DomainResourceReports{
					"capacity": &limesresources.DomainResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						DomainQuota:   p2ui64(55),
						ProjectsQuota: p2ui64(25),
						Usage:         10,
					},
					"things": &limesresources.DomainResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   p2ui64(55),
						ProjectsQuota: p2ui64(25),
						Usage:         10,
					},
				},
				MaxScrapedAt: p2time(11),
				MinScrapedAt: p2time(11),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetDomainFiltered(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetDomainSuccessfully(t)

	actual, err := domains.Get(t.Context(), fake.ServiceClient(), "uuid-for-karachi", domains.GetOpts{
		Services:  []limes.ServiceType{"shared"},
		Resources: []limesresources.ResourceName{"things"},
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesresources.DomainReport{
		DomainInfo: limes.DomainInfo{
			UUID: "uuid-for-karachi",
			Name: "karachi",
		},
		Services: limesresources.DomainServiceReports{
			"shared": &limesresources.DomainServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type: "shared",
					Area: "shared",
				},
				Resources: limesresources.DomainResourceReports{
					"things": &limesresources.DomainResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						DomainQuota:   p2ui64(10),
						ProjectsQuota: p2ui64(5),
						Usage:         2,
					},
				},
				MaxScrapedAt: p2time(22),
				MinScrapedAt: p2time(22),
			},
		},
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
