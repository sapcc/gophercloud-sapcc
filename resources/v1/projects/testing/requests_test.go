package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/pagination"
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

	pager := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{})

	count := 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		actual, err := projects.ExtractProjects(page)
		if err != nil {
			return false, err
		}
		count += len(actual)
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
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, count)
}

func TestListProjectsDetailed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	pager := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Detail: true})

	count := 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		actual, err := projects.ExtractProjects(page)
		if err != nil {
			return false, err
		}
		count += len(actual)
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
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, count)
}

func TestListProjectsFiltered(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListProjectsSuccessfully(t)

	pager := projects.List(fakeclient.ServiceClient(), "uuid-for-germany", projects.ListOpts{Service: "shared", Resource: "things"})

	count := 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		actual, err := projects.ExtractProjects(page)
		if err != nil {
			return false, err
		}
		count += len(actual)
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
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, count)
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

func TestUpdateProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePutProjectSuccessfully(t)

	opts := projects.UpdateOpts{
		Services: api.ServiceQuotas{
			"compute": api.ResourceQuotas{
				"cores": limes.ValueWithUnit{42, limes.UnitNone},
			},
		},
	}

	actual, err := projects.Update(fakeclient.ServiceClient(), "uuid-for-germany", "uuid-for-berlin", opts).Extract()
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
						Quota: 42,
						Usage: 23,
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
