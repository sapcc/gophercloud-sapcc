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

//nolint:dupl
package testing

import (
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesresources "github.com/sapcc/go-api-declarations/limes/resources"

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

	actual, err := domains.List(fake.ServiceClient(), domains.ListOpts{
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

	actual, err := domains.Get(fake.ServiceClient(), "uuid-for-karachi", domains.GetOpts{}).Extract()
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

	actual, err := domains.Get(fake.ServiceClient(), "uuid-for-karachi", domains.GetOpts{
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

func TestUpdateDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateDomainSuccessfully(t)

	opts := domains.UpdateOpts{
		Services: limesresources.QuotaRequest{
			"shared": limesresources.ServiceQuotaRequest{
				"things": limesresources.ResourceQuotaRequest{
					Value: 99,
					Unit:  limes.UnitNone,
				},
			},
		},
	}

	// if update succeeds then a 202 (no error) is returned.
	err := domains.Update(fake.ServiceClient(), "uuid-for-karachi", opts).ExtractErr()
	th.AssertNoErr(t, err)
}

func p2time(timestamp int64) *limes.UnixEncodedTime {
	t := limes.UnixEncodedTime{Time: time.Unix(timestamp, 0).UTC()}
	return &t
}

func p2ui64(x uint64) *uint64 {
	return &x
}
