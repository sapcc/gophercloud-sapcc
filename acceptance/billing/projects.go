package billing

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/sapcc/gophercloud-billing/billing/projects"
)

func getIntField(v interface{}, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	if f.Kind() == reflect.Ptr {
		return int(f.Elem().Int())
	}

	return int(f.Int())
}

func getStrField(v interface{}, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	if f.Kind() == reflect.Ptr {
		return f.Elem().String()
	}

	return f.String()
}

func UpdateProjectField(t *testing.T, client *gophercloud.ServiceClient, project *projects.Project, field string) {
	opts := projects.UpdateOpts{
		ResponsiblePrimaryContactID:    project.ResponsiblePrimaryContactID,
		ResponsiblePrimaryContactEmail: project.ResponsiblePrimaryContactEmail,
		CostObject:                     project.CostObject,
	}

	switch field {
	case "Description":
		opts.Description = getStrField(project, field) + " updated"
		data := update(t, client, project.ProjectID, opts)
		th.AssertDeepEquals(t, getStrField(opts, field), getStrField(data, field))
	case "RevenueRelevance":
		opts.RevenueRelevance = "generating"
		data := update(t, client, project.ProjectID, opts)
		th.AssertDeepEquals(t, getStrField(opts, field), getStrField(data, field))
	case "BusinessCriticality":
		opts.BusinessCriticality = "test"
		data := update(t, client, project.ProjectID, opts)
		th.AssertDeepEquals(t, getStrField(opts, field), getStrField(data, field))
	case "AdditionalInformation":
		opts.AdditionalInformation = "extra info"
		data := update(t, client, project.ProjectID, opts)
		th.AssertDeepEquals(t, getStrField(opts, field), getStrField(data, field))
	case "NumberOfEndusers":
		opts.NumberOfEndusers = 100
		data := update(t, client, project.ProjectID, opts)
		th.AssertDeepEquals(t, getIntField(opts, field), getIntField(data, field))
	case "CostObject":
		opts.CostObject = projects.CostObject{
			Inherited: true,
		}
		data := update(t, client, project.ProjectID, opts)
		th.AssertDeepEquals(t, opts.CostObject, data.CostObject)
	default:
		th.AssertNoErr(t, fmt.Errorf("Unknown field %s", field))
	}
}

func UpdateProject(t *testing.T, client *gophercloud.ServiceClient, id string, opts projects.UpdateOpts) {
	res := projects.Update(client, id, opts)
	th.AssertNoErr(t, res.Err)
}

func update(t *testing.T, client *gophercloud.ServiceClient, id string, opts projects.UpdateOpts) *projects.Project {
	data, err := projects.Update(client, id, opts).Extract()
	th.AssertNoErr(t, err)
	return data
}
