package billing

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/tools"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/sapcc/gophercloud-billing/billing/projects"
)

var projectID = "e9141fb24eee4b3e9f25ae69cda31132"

func TestProjectReadUpdate(t *testing.T) {
	client, err := NewBillingClient()
	th.AssertNoErr(t, err)

	// Get project
	project, err := projects.Get(client, projectID).Extract()
	th.AssertNoErr(t, err)

	// restore initial project data
	defer UpdateProject(t, client, projectID, ProjectToUpdateOpts(project))
}

func TestProjectList(t *testing.T) {
	client, err := NewBillingClient()
	th.AssertNoErr(t, err)

	// Get project
	project, err := projects.Get(client, projectID).Extract()
	th.AssertNoErr(t, err)

	// Get projects
	allPages, err := projects.List(client).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, allProjects)

	// compare project and projects list
	th.AssertDeepEquals(t, allProjects[0], project)
}
