// +build acceptance networking

package automation

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/kayrus/gophercloud-lyra/automation/v1/automations"
	"github.com/kayrus/gophercloud-lyra/automation/v1/runs"
)

func TestRunCR(t *testing.T) {
	client, err := NewAutomationV1Client()
	th.AssertNoErr(t, err)

	// Create Chef and Script automations
	runList := []string{"foo"}
	chefAutomation, err := CreateChefAutomation(t, client, runList)
	th.AssertNoErr(t, err)
	defer automations.Delete(client, chefAutomation.ID)

	tools.PrintResource(t, chefAutomation)

	run, err := CreateRun(t, client, chefAutomation.ID)
	th.AssertNoErr(t, err)

	// second run for the list test below
	_, err = CreateRun(t, client, chefAutomation.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, run)

	// Read the run
	newRun, err := runs.Get(client, run.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newRun)
	th.AssertEquals(t, run.Selector, newRun.Selector)
	th.AssertEquals(t, run.AutomationID, newRun.AutomationID)
}

func TestRunList(t *testing.T) {
	client, err := NewAutomationV1Client()
	th.AssertNoErr(t, err)

	count := 0
	var allRuns []runs.Run

	runs.List(client, runs.ListOpts{PerPage: 1}).EachPage(func(page pagination.Page) (bool, error) {
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
