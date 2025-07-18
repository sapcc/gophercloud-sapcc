// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"testing"

	"github.com/sapcc/gophercloud-sapcc/v2/internal/acceptance/tools"

	"github.com/gophercloud/gophercloud/v2"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/automation/v1/automations"
)

// CreateChefAutomation will create a Chef automation. An error will be
// returned if the automation could not be created.
func CreateChefAutomation(t *testing.T, client *gophercloud.ServiceClient, runList []string) (*automations.Automation, error) {
	automationName := tools.RandomString("TESTACC-", 8)

	t.Logf("Attempting to create Chef automation: %s", automationName)

	createOpts := automations.CreateOpts{
		Name:       automationName,
		Repository: "https://example.com/chef/repo.git",
		Type:       "Chef",
		Debug:      true,
		RunList:    runList,
	}

	s, err := automations.Create(t.Context(), client, createOpts).Extract()
	if err != nil {
		return s, err
	}

	t.Logf("Successfully created Chef automation: %s", automationName)

	th.AssertEquals(t, s.Name, automationName)
	th.AssertEquals(t, s.Repository, "https://example.com/chef/repo.git")
	th.AssertEquals(t, s.Type, "Chef")
	th.AssertEquals(t, s.Debug, true)
	th.AssertDeepEquals(t, s.RunList, runList)

	return s, nil
}

// CreateScriptAutomation will create a Script automation. An error will be
// returned if the automation could not be created.
func CreateScriptAutomation(t *testing.T, client *gophercloud.ServiceClient, path string) (*automations.Automation, error) {
	automationName := tools.RandomString("TESTACC-", 8)

	t.Logf("Attempting to create Script automation: %s", automationName)

	createOpts := automations.CreateOpts{
		Name:       automationName,
		Repository: "https://example.com/script/repo.git",
		Type:       "Script",
		Path:       path,
	}

	s, err := automations.Create(t.Context(), client, createOpts).Extract()
	if err != nil {
		return s, err
	}

	t.Logf("Successfully created Script automation: %s", automationName)

	th.AssertEquals(t, s.Name, automationName)
	th.AssertEquals(t, s.Repository, "https://example.com/script/repo.git")
	th.AssertEquals(t, s.Type, "Script")
	th.AssertEquals(t, s.Path, path)

	return s, nil
}
