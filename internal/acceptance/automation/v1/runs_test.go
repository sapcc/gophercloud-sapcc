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
	"github.com/sapcc/gophercloud-sapcc/v2/automation/v1/runs"
)

func TestRunCR(t *testing.T) {
	client, err := NewAutomationV1Client(t.Context())
	th.AssertNoErr(t, err)

	// Create Chef and Script automations
	runList := []string{"foo"}
	chefAutomation, err := CreateChefAutomation(t, client, runList)
	th.AssertNoErr(t, err)
	defer automations.Delete(t.Context(), client, chefAutomation.ID)

	tools.PrintResource(t, chefAutomation)

	run, err := CreateRun(t, client, chefAutomation.ID)
	th.AssertNoErr(t, err)

	// second run for the list test below
	_, err = CreateRun(t, client, chefAutomation.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, run)

	// Read the run
	newRun, err := runs.Get(t.Context(), client, run.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newRun)
	th.AssertEquals(t, run.Selector, newRun.Selector)
	th.AssertEquals(t, run.AutomationID, newRun.AutomationID)
}

func TestRunList(t *testing.T) {
	client, err := NewAutomationV1Client(t.Context())
	th.AssertNoErr(t, err)

	count := 0
	var allRuns []runs.Run

	//nolint:errcheck
	runs.List(client, runs.ListOpts{PerPage: 1}).EachPage(t.Context(), func(ctx context.Context, page pagination.Page) (bool, error) {
		count++
		tmp, err := runs.ExtractRuns(page)
		if err != nil {
			t.Errorf("Failed to extract runs: %v", err)
			return false, nil
		}

		allRuns = append(allRuns, tmp...)

		return true, nil
	})

	tools.PrintResource(t, allRuns)

	expectedPages := 2
	if count < expectedPages {
		t.Errorf("Expected more than %d page, got %d", expectedPages, count)
	}

	expectedRuns := 2
	if len(allRuns) < expectedRuns {
		t.Errorf("Expected more than %d runs, got %d", expectedRuns, len(allRuns))
	}
}
