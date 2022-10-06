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
	"encoding/json"
	"testing"

	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/go-api-declarations/limes"

	"github.com/sapcc/gophercloud-sapcc/resources/v1/projects"
)

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{}).ExtractProjects()
	th.AssertNoErr(t, err)

	backendQuota := int64(100)
	expected := []limes.ProjectReport{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"capacity": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2i64(22),
				},
				"unshared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"capacity": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2i64(11),
				},
			},
		},
		{
			UUID:       "uuid-for-dresden",
			Name:       "dresden",
			ParentUUID: "uuid-for-berlin",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"capacity": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota:        p2ui64(10),
							Usage:        2,
							BackendQuota: &backendQuota,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2i64(44),
				},
				"unshared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"capacity": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2i64(33),
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

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Detail: true}).ExtractProjects()
	th.AssertNoErr(t, err)

	expected := []limes.ProjectReport{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"capacity": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota:        p2ui64(10),
							Usage:        2,
							Subresources: json.RawMessage(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
						},
					},
					ScrapedAt: p2i64(22),
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

	actual, err := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{
		Services:  []string{"shared"},
		Resources: []string{"things"},
	}).ExtractProjects()
	th.AssertNoErr(t, err)

	expected := []limes.ProjectReport{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2i64(22),
				},
			},
		},
		{
			UUID:       "uuid-for-dresden",
			Name:       "dresden",
			ParentUUID: "uuid-for-berlin",
			Services: limes.ProjectServiceReports{
				"shared": &limes.ProjectServiceReport{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: limes.ProjectResourceReports{
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: p2ui64(10),
							Usage: 2,
						},
					},
					ScrapedAt: p2i64(44),
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

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	expected := &limes.ProjectReport{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: limes.ProjectServiceReports{
			"shared": &limes.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limes.ProjectResourceReports{
					"capacity": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
					"things": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
				},
				ScrapedAt: p2i64(22),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectDetailed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{Detail: true}).Extract()
	th.AssertNoErr(t, err)

	expected := &limes.ProjectReport{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: limes.ProjectServiceReports{
			"shared": &limes.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limes.ProjectResourceReports{
					"capacity": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
					"things": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota:        p2ui64(10),
						Usage:        2,
						Subresources: json.RawMessage(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
					},
				},
				ScrapedAt: p2i64(22),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetProjectFiltered(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetProjectSuccessfully(t)

	actual, err := projects.Get(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.GetOpts{
		Services:  []string{"shared"},
		Resources: []string{"things"},
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &limes.ProjectReport{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: limes.ProjectServiceReports{
			"shared": &limes.ProjectServiceReport{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: limes.ProjectResourceReports{
					"things": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota: p2ui64(10),
						Usage: 2,
					},
				},
				ScrapedAt: p2i64(22),
			},
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestUpdateProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateProjectSuccessfully(t)

	opts := projects.UpdateOpts{
		Services: limes.QuotaRequest{
			"compute": limes.ServiceQuotaRequest{
				Resources: limes.ResourceQuotaRequest{
					"cores": limes.ValueWithUnit{Value: 42, Unit: limes.UnitNone},
				},
			},
		},
	}

	// if update succeeds then a 202 (no error) is returned.
	warn, err := projects.Update(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []byte{}, warn)
}

func TestUpdateProjectWarning(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateProjectPartly(t)

	// expecting to get 202 response code with a warning
	body, err := projects.Update(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.UpdateOpts{}).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, body, []byte("it is currently not allowed to set bursting.enabled and quotas in the same request"))
}

func TestUpdateProjectError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateProjectUnsuccessfully(t)

	// expecting to get 400 response code
	warn, err := projects.Update(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", projects.UpdateOpts{}).Extract()
	th.AssertErr(t, err)
	th.AssertDeepEquals(t, []byte(nil), warn)
	if err, ok := err.(gophercloud.ErrDefault400); ok {
		th.AssertDeepEquals(t, err.Body, []byte("it is currently not allowed to set bursting.enabled and quotas in the same request"))
	} else {
		t.Fatalf("Unexpected error response")
	}
}

func TestSyncProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSyncProjectSuccessfully(t)

	// if sync succeeds then a 202 (no error) is returned.
	err := projects.Sync(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-dresden").ExtractErr()
	th.AssertNoErr(t, err)
}

func p2i64(x int64) *int64 {
	return &x
}

func p2ui64(x uint64) *uint64 {
	return &x
}
