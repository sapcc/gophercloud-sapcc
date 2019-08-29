package billing

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/tools"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/sapcc/gophercloud-billing/billing/masterdata/projects"
)

var projectID = "3e0fd3f8e9ec449686ef26a16a284265"

func TestProjectReadUpdate(t *testing.T) {
	client, err := NewBillingClient()
	th.AssertNoErr(t, err)

	// Get project
	project, err := projects.Get(client, projectID).Extract()
	th.AssertNoErr(t, err)

	// restore initial project data
	defer UpdateProject(t, client, projectID, projects.ProjectToUpdateOpts(project))

	UpdateProjectField(t, client, project, "Description")
	UpdateProjectField(t, client, project, "RevenueRelevance")
	UpdateProjectField(t, client, project, "BusinessCriticality")
	UpdateProjectField(t, client, project, "AdditionalInformation")
	UpdateProjectField(t, client, project, "NumberOfEndusers")

	// valid project is required
	//UpdateProjectField(t, client, project, "CostObject")
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
	th.AssertDeepEquals(t, allProjects[0], *project)
}
