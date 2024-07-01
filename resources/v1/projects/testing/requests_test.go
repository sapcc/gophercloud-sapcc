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
	"context"
	"encoding/json"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"
	limesresources "github.com/sapcc/go-api-declarations/limes/resources"

	"github.com/sapcc/gophercloud-sapcc/resources/v1/projects"
)

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{}).ExtractProjects()
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

	actual, err := projects.List(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Detail: true}).ExtractProjects()
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

	actual, err := projects.List(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{
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

	actual, err := projects.Get(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{}).Extract()
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

	actual, err := projects.Get(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{Detail: true}).Extract()
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

	actual, err := projects.Get(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{
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
	err := projects.Sync(context.TODO(), fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-dresden").ExtractErr()
	th.AssertNoErr(t, err)
}

func p2time(timestamp int64) *limes.UnixEncodedTime {
	t := limes.UnixEncodedTime{Time: time.Unix(timestamp, 0).UTC()}
	return &t
}

func p2ui64(x uint64) *uint64 {
	return &x
}
