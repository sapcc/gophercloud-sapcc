package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-limes/resources/v1/projects"
	"github.com/sapcc/limes/pkg/api"
	"github.com/sapcc/limes/pkg/limes"
	"github.com/sapcc/limes/pkg/reports"
	"github.com/sapcc/limes/pkg/util"
)

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	result := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{})
	actual, err := result.ExtractProjects()
	th.AssertNoErr(t, err)

	backendQuota := int64(100)
	expected := []reports.Project{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: reports.ProjectServices{
				"shared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"capacity": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: 10,
							Usage: 2,
						},
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
							Usage: 2,
						},
					},
					ScrapedAt: 22,
				},
				"unshared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"capacity": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: 10,
							Usage: 2,
						},
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
							Usage: 2,
						},
					},
					ScrapedAt: 11,
				},
			},
		},
		{
			UUID:       "uuid-for-dresden",
			Name:       "dresden",
			ParentUUID: "uuid-for-berlin",
			Services: reports.ProjectServices{
				"shared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"capacity": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota:        10,
							Usage:        2,
							BackendQuota: &backendQuota,
						},
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
							Usage: 2,
						},
					},
					ScrapedAt: 44,
				},
				"unshared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "unshared",
						Area:        "unshared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"capacity": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: 10,
							Usage: 2,
						},
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
							Usage: 2,
						},
					},
					ScrapedAt: 33,
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

	result := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Detail: true})
	actual, err := result.ExtractProjects()
	th.AssertNoErr(t, err)

	expected := []reports.Project{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: reports.ProjectServices{
				"shared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"capacity": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "capacity",
								Unit: limes.UnitBytes,
							},
							Quota: 10,
							Usage: 2,
						},
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota:        10,
							Usage:        2,
							Subresources: util.JSONString(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
						},
					},
					ScrapedAt: 22,
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

	result := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Service: "shared", Resource: "things"})
	actual, err := result.ExtractProjects()
	th.AssertNoErr(t, err)

	expected := []reports.Project{
		{
			UUID:       "uuid-for-berlin",
			Name:       "berlin",
			ParentUUID: "uuid-for-germany",
			Services: reports.ProjectServices{
				"shared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
							Usage: 2,
						},
					},
					ScrapedAt: 22,
				},
			},
		},
		{
			UUID:       "uuid-for-dresden",
			Name:       "dresden",
			ParentUUID: "uuid-for-berlin",
			Services: reports.ProjectServices{
				"shared": &reports.ProjectService{
					ServiceInfo: limes.ServiceInfo{
						Type:        "shared",
						Area:        "shared",
						ProductName: "",
					},
					Resources: reports.ProjectResources{
						"things": &reports.ProjectResource{
							ResourceInfo: limes.ResourceInfo{
								Name: "things",
							},
							Quota: 10,
							Usage: 2,
						},
					},
					ScrapedAt: 44,
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

	expected := &reports.Project{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: reports.ProjectServices{
			"shared": &reports.ProjectService{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: reports.ProjectResources{
					"capacity": &reports.ProjectResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						Quota: 10,
						Usage: 2,
					},
					"things": &reports.ProjectResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota: 10,
						Usage: 2,
					},
				},
				ScrapedAt: 22,
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

	expected := &reports.Project{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: reports.ProjectServices{
			"shared": &reports.ProjectService{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: reports.ProjectResources{
					"capacity": &reports.ProjectResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "capacity",
							Unit: limes.UnitBytes,
						},
						Quota: 10,
						Usage: 2,
					},
					"things": &reports.ProjectResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota:        10,
						Usage:        2,
						Subresources: util.JSONString(`[{"id":"thirdthing","value":5},{"id":"fourththing","value":123}]`),
					},
				},
				ScrapedAt: 22,
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
		Service:  "shared",
		Resource: "things",
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &reports.Project{
		UUID:       "uuid-for-berlin",
		Name:       "berlin",
		ParentUUID: "uuid-for-germany",
		Services: reports.ProjectServices{
			"shared": &reports.ProjectService{
				ServiceInfo: limes.ServiceInfo{
					Type:        "shared",
					Area:        "shared",
					ProductName: "",
				},
				Resources: reports.ProjectResources{
					"things": &reports.ProjectResource{
						ResourceInfo: limes.ResourceInfo{
							Name: "things",
						},
						Quota: 10,
						Usage: 2,
					},
				},
				ScrapedAt: 22,
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
		Services: api.ServiceQuotas{
			"compute": api.ResourceQuotas{
				"cores": limes.ValueWithUnit{42, limes.UnitNone},
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
	err := projects.Sync(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-dresden")
	th.AssertNoErr(t, err)
}
