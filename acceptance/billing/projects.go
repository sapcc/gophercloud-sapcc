package billing

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/sapcc/gophercloud-billing/billing/projects"
)

func ProjectToUpdateOpts(project *projects.Project) projects.UpdateOpts {
	return projects.UpdateOpts{
		ProjectID:                      &project.ProjectID,
		ProjectName:                    &project.ProjectName,
		DomainID:                       &project.DomainID,
		DomainName:                     &project.DomainName,
		Description:                    &project.Description,
		ParentID:                       &project.ParentID,
		ProjectType:                    &project.ProjectType,
		ResponsiblePrimaryContactID:    &project.ResponsiblePrimaryContactID,
		ResponsiblePrimaryContactEmail: &project.ResponsiblePrimaryContactEmail,
		ResponsibleOperatorID:          &project.ResponsibleOperatorID,
		ResponsibleOperatorEmail:       &project.ResponsibleOperatorEmail,
		ResponsibleSecurityExpertID:    &project.ResponsibleSecurityExpertID,
		ResponsibleSecurityExpertEmail: &project.ResponsibleSecurityExpertEmail,
		ResponsibleProductOwnerID:      &project.ResponsibleProductOwnerID,
		ResponsibleProductOwnerEmail:   &project.ResponsibleProductOwnerEmail,
		ResponsibleControllerID:        &project.ResponsibleControllerID,
		ResponsibleControllerEmail:     &project.ResponsibleControllerEmail,
		RevenueRelevance:               &project.RevenueRelevance,
		BusinessCriticality:            &project.BusinessCriticality,
		NumberOfEndusers:               &project.NumberOfEndusers,
		AdditionalInformation:          &project.AdditionalInformation,
		CostObject:                     &project.CostObject,
		CreatedAt:                      &project.CreatedAt,
		ChangedAt:                      &project.ChangedAt,
		ChangedBy:                      &project.ChangedBy,
		IsComplete:                     &project.IsComplete,
		MissingAttributes:              &project.MissingAttributes,
		Collector:                      &project.Collector,
		Region:                         &project.Region,
	}
}

func UpdateProject(t *testing.T, client *gophercloud.ServiceClient, id string, opts projects.UpdateOpts) {
	res := projects.Update(client, id, opts)
	th.AssertNoErr(t, res.Err)
}
