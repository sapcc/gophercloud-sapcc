package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-sapcc/resources/v1/projects"
	"github.com/sapcc/limes"
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
							Quota: 10,
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
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
							Quota: 10,
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
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
							Quota:        10,
							Usage:        2,
							BackendQuota: &backendQuota,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
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
							Quota: 10,
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
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
							Quota: 10,
							Usage: 2,
						},
						"things": &limes.ProjectResourceReport{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota:        10,
							Usage:        2,
							Subresources: limes.JSONString(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
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
		Cluster:  "fakecluster",
		Service:  "shared",
		Resource: "things",
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
							Quota: 10,
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
							Quota: 10,
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
						Quota: 10,
						Usage: 2,
					},
					"things": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota: 10,
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
						Quota: 10,
						Usage: 2,
					},
					"things": &limes.ProjectResourceReport{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota:        10,
						Usage:        2,
						Subresources: limes.JSONString(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
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
		Cluster:  "fakecluster",
		Service:  "shared",
		Resource: "things",
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
						Quota: 10,
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
		Cluster: "fakecluster",
		Services: limes.QuotaRequest{
			"compute": limes.ServiceQuotaRequest{
				"cores": limes.ValueWithUnit{Value: 42, Unit: limes.UnitNone},
			},
		},
	}

	// if update succeeds then a 202 (no error) is returned.
	_, err := projects.Update(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", opts)
	th.AssertNoErr(t, err)
}

func TestSyncProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSyncProjectSuccessfully(t)

	// if sync succeeds then a 202 (no error) is returned.
	err := projects.Sync(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-dresden", projects.SyncOpts{
		Cluster: "fakecluster"})
	th.AssertNoErr(t, err)
}

func p2i64(x int64) *int64 {
	return &x
}
