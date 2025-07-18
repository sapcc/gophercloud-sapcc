// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/internal/acceptance/tools"

	"github.com/sapcc/gophercloud-sapcc/v2/automation/v1/automations"
)

func TestAutomationCRUD(t *testing.T) {
	client, err := NewAutomationV1Client(t.Context())
	th.AssertNoErr(t, err)

	// Create Chef and Script automations
	runList := []string{"foo"}
	chefAutomation, err := CreateChefAutomation(t, client, runList)
	th.AssertNoErr(t, err)
	defer automations.Delete(t.Context(), client, chefAutomation.ID)

	tools.PrintResource(t, chefAutomation)

	path := "foo"
	scriptAutomation, err := CreateScriptAutomation(t, client, path)
	th.AssertNoErr(t, err)
	defer automations.Delete(t.Context(), client, scriptAutomation.ID)

	tools.PrintResource(t, scriptAutomation)

	// Update Chef automation
	newRunList := []string{"bar"}
	chefAttributes := map[string]any{"foo": "bar"}
	updateOpts := automations.UpdateOpts{
		RunList:        newRunList,
		ChefAttributes: chefAttributes,
	}

	newChefAutomation, err := automations.Update(t.Context(), client, chefAutomation.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newChefAutomation)

	// Read the updated automation
	newGetAutomation, err := automations.Get(t.Context(), client, chefAutomation.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newGetAutomation)
	th.AssertDeepEquals(t, newChefAutomation, newGetAutomation)
	th.AssertDeepEquals(t, newChefAutomation.RunList, newRunList)
	th.AssertDeepEquals(t, newChefAutomation.ChefAttributes, chefAttributes)

	// Unset attributes from the Chef automation
	newRevision := "dev"
	updateOpts = automations.UpdateOpts{
		ChefAttributes:     map[string]any{},
		RepositoryRevision: &newRevision,
	}

	newChefAutomation, err = automations.Update(t.Context(), client, chefAutomation.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newChefAutomation)

	// Read the updated automation
	newGetAutomation, err = automations.Get(t.Context(), client, chefAutomation.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newGetAutomation)
	th.AssertDeepEquals(t, newChefAutomation, newGetAutomation)
	th.AssertDeepEquals(t, newChefAutomation.ChefAttributes, map[string]any(nil))
	th.AssertDeepEquals(t, newChefAutomation.RepositoryRevision, newRevision)

	// Update Script automation
	newPath := "bar"
	arguments := []string{"foo", "bar"}
	environment := map[string]string{"FOO": "BAR"}
	updateOpts = automations.UpdateOpts{
		RepositoryRevision: &newRevision,
		Path:               &newPath,
		Arguments:          arguments,
		Environment:        environment,
		Timeout:            1,
	}

	newScriptAutomation, err := automations.Update(t.Context(), client, scriptAutomation.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newScriptAutomation)

	// Read the updated automation
	newGetAutomation, err = automations.Get(t.Context(), client, scriptAutomation.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newGetAutomation)
	th.AssertDeepEquals(t, newScriptAutomation, newGetAutomation)
	th.AssertDeepEquals(t, newScriptAutomation.Path, newPath)
	th.AssertDeepEquals(t, newScriptAutomation.Arguments, arguments)
	th.AssertDeepEquals(t, newScriptAutomation.Environment, environment)
	th.AssertDeepEquals(t, newScriptAutomation.Timeout, 1)
	th.AssertDeepEquals(t, newChefAutomation.RepositoryRevision, newRevision)

	// Unset attributes from the Script automation
	updateOpts = automations.UpdateOpts{
		RepositoryRevision: new(string),
		Arguments:          []string{},
		Environment:        map[string]string{},
		RunList:            []string{},
		Timeout:            3600,
	}

	newScriptAutomation, err = automations.Update(t.Context(), client, scriptAutomation.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newScriptAutomation)

	// Read the updated automation
	newGetAutomation, err = automations.Get(t.Context(), client, scriptAutomation.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newGetAutomation)
	th.AssertDeepEquals(t, newScriptAutomation, newGetAutomation)
	th.AssertDeepEquals(t, newScriptAutomation.Path, newPath)
	th.AssertDeepEquals(t, newScriptAutomation.Arguments, []string(nil))
	th.AssertDeepEquals(t, newScriptAutomation.Environment, map[string]string(nil))
	th.AssertDeepEquals(t, newScriptAutomation.Timeout, 3600)
	th.AssertDeepEquals(t, newScriptAutomation.RepositoryRevision, "")
}

func TestAutomationList(t *testing.T) {
	client, err := NewAutomationV1Client(t.Context())
	th.AssertNoErr(t, err)

	// Create automations
	runList := []string{"foo"}
	chefAutomation, err := CreateChefAutomation(t, client, runList)
	th.AssertNoErr(t, err)
	defer automations.Delete(t.Context(), client, chefAutomation.ID)

	path := "foo"
	scriptAutomation, err := CreateScriptAutomation(t, client, path)
	th.AssertNoErr(t, err)
	defer automations.Delete(t.Context(), client, scriptAutomation.ID)

	count := 0
	var allAutomations []automations.Automation

	//nolint:errcheck
	automations.List(client, automations.ListOpts{PerPage: 1}).EachPage(t.Context(), func(ctx context.Context, page pagination.Page) (bool, error) {
		count++
		tmp, err := automations.ExtractAutomations(page)
		if err != nil {
			t.Errorf("Failed to extract automations: %v", err)
			return false, nil
		}

		allAutomations = append(allAutomations, tmp...)

		return true, nil
	})

	tools.PrintResource(t, allAutomations)

	expectedPages := 2
	if count != expectedPages {
		t.Errorf("Expected %d page, got %d", expectedPages, count)
	}
	th.AssertDeepEquals(t, 2, len(allAutomations))
}
