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

package billing

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/tools"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/sapcc/gophercloud-sapcc/billing/masterdata/projects"
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
	allPages, err := projects.List(client, projects.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, allProjects)

	// compare project and projects list
	th.AssertDeepEquals(t, allProjects[0], *project)
}
