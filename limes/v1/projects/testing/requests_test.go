package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/sapcc/gophercloud-limes/limes/v1/projects"
)

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListProjectsSuccessfully(t)

	t.Logf("Id\tName\tOwner\tChecksum\tSizeBytes")

	pager := projects.List(fakeclient.ServiceClient(), "abc123", projects.ListOpts{})

	count, pages := 0, 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		pages++
		projects, err := projects.ExtractProjects(page)
		if err != nil {
			return false, err
		}

		for _, i := range projects {
			t.Logf("%+v", i)
			count++
		}

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, pages)
	th.AssertEquals(t, 2, count)
}
