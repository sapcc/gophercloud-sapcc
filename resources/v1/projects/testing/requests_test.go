// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"encoding/json"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesresources "github.com/sapcc/go-api-declarations/limes/resources"

	"github.com/sapcc/gophercloud-sapcc/v2/resources/v1/projects"
)

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{}).ExtractProjects()
	th.AssertNoErr(t, err)

	backendQuota := int64(100)
	expected := []limesresources.ProjectReport{
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-berlin",
				Name:       "berlin",
				ParentUUID: "uuid-for-germany",
			},
			Services: limesresources.ProjectServiceReports{
				"shared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"capacity": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2time(22),
				},
				"unshared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"capacity": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2time(11),
				},
			},
		},
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-dresden",
				Name:       "dresden",
				ParentUUID: "uuid-for-berlin",
			},
			Services: limesresources.ProjectServiceReports{
				"shared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"capacity": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota:        p2ui64(10),
							Usage:        2,
							BackendQuota: &backendQuota,
						},
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2time(44),
				},
				"unshared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"capacity": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2time(33),
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListProjectsDetailed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Detail: true}).ExtractProjects()
	th.AssertNoErr(t, err)

	expected := []limesresources.ProjectReport{
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-berlin",
				Name:       "berlin",
				ParentUUID: "uuid-for-germany",
			},
			Services: limesresources.ProjectServiceReports{
				"shared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"capacity": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota:        p2ui64(10),
							Usage:        2,
							Subresources: json.RawMessage(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
						},
					},
					ScrapedAt: p2time(22),
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListProjectsFiltered(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{
		Services:  []limes.ServiceType{"shared"},
		Resources: []limesresources.ResourceName{"things"},
	}).ExtractProjects()
	th.AssertNoErr(t, err)

	expected := []limesresources.ProjectReport{
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-berlin",
				Name:       "berlin",
				ParentUUID: "uuid-for-germany",
			},
			Services: limesresources.ProjectServiceReports{
				"shared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2time(22),
				},
			},
		},
		{
			ProjectInfo: limes.ProjectInfo{
				UUID:       "uuid-for-dresden",
				Name:       "dresden",
				ParentUUID: "uuid-for-berlin",
			},
			Services: limesresources.ProjectServiceReports{
				"shared": &limesresources.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limesresources.ProjectResourceReports{
						"things": &limesresources.ProjectResourceReport{
							ResourceInfo: limesresources.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2time(44),
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesresources.ProjectReport{
		ProjectInfo: limes.ProjectInfo{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
		},
		Services: limesresources.ProjectServiceReports{
			"shared": &limesresources.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limesresources.ProjectResourceReports{
					"capacity": &limesresources.ProjectResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
					"things": &limesresources.ProjectResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
				},
				ScrapedAt: p2time(22),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectDetailed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{Detail: true}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesresources.ProjectReport{
		ProjectInfo: limes.ProjectInfo{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
		},
		Services: limesresources.ProjectServiceReports{
			"shared": &limesresources.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limesresources.ProjectResourceReports{
					"capacity": &limesresources.ProjectResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
					"things": &limesresources.ProjectResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						Quota:        p2ui64(10),
						Usage:        2,
						Subresources: json.RawMessage(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
					},
				},
				ScrapedAt: p2time(22),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectFiltered(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{
		Services:  []limes.ServiceType{"shared"},
		Resources: []limesresources.ResourceName{"things"},
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &limesresources.ProjectReport{
		ProjectInfo: limes.ProjectInfo{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
		},
		Services: limesresources.ProjectServiceReports{
			"shared": &limesresources.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limesresources.ProjectResourceReports{
					"things": &limesresources.ProjectResourceReport{
						ResourceInfo: limesresources.ResourceInfo{
							Name: "things",
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
				},
				ScrapedAt: p2time(22),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestSyncProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSyncProjectSuccessfully(t)

	// if sync succeeds then a 202 (no error) is returned.
	err := projects.Sync(t.Context(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-dresden").ExtractErr()
	th.AssertNoErr(t, err)
}

func p2time(timestamp int64) *limes.UnixEncodedTime {
	t := limes.UnixEncodedTime{Time: time.Unix(timestamp, 0).UTC()}
	return &t
}

func p2ui64(x uint64) *uint64 {
	return &x
}
