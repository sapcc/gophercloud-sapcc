// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package billing

import (
	"testing"

	"github.com/sapcc/gophercloud-sapcc/v2/internal/acceptance/tools"

	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/billing/masterdata/projects"
)

var projectID = "3e0fd3f8e9ec449686ef26a16a284265"

func TestProjectReadUpdate(t *testing.T) {
	client, err := NewBillingClient(t.Context())
	th.AssertNoErr(t, err)

	// Get project
	project, err := projects.Get(t.Context(), client, projectID).Extract()
	th.AssertNoErr(t, err)

	// restore initial project data
	defer UpdateProject(t, client, projectID, projects.ProjectToUpdateOpts(project))

	UpdateProjectField(t, client, project, "Description")
	UpdateProjectField(t, client, project, "RevenueRelevance")
	UpdateProjectField(t, client, project, "BusinessCriticality")
	UpdateProjectField(t, client, project, "AdditionalInformation")
	UpdateProjectField(t, client, project, "NumberOfEndusers")

	// valid project is required
	// UpdateProjectField(t, client, project, "CostObject")
}

func TestProjectList(t *testing.T) {
	client, err := NewBillingClient(t.Context())
	th.AssertNoErr(t, err)

	// Get project
	project, err := projects.Get(t.Context(), client, projectID).Extract()
	th.AssertNoErr(t, err)

	// Get projects
	allPages, err := projects.List(client, projects.ListOpts{}).AllPages(t.Context())
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, allProjects)

	// compare project and projects list
	th.AssertDeepEquals(t, allProjects[0], *project)
}
