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

package v1

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/internal/acceptance/tools"

	"github.com/sapcc/gophercloud-sapcc/automation/v1/automations"
	"github.com/sapcc/gophercloud-sapcc/automation/v1/runs"
)

func TestRunCR(t *testing.T) {
	client, err := NewAutomationV1Client(context.TODO())
	th.AssertNoErr(t, err)

	// Create Chef and Script automations
	runList := []string{"foo"}
	chefAutomation, err := CreateChefAutomation(t, client, runList)
	th.AssertNoErr(t, err)
	defer automations.Delete(context.TODO(), client, chefAutomation.ID)

	tools.PrintResource(t, chefAutomation)

	run, err := CreateRun(t, client, chefAutomation.ID)
	th.AssertNoErr(t, err)

	// second run for the list test below
	_, err = CreateRun(t, client, chefAutomation.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, run)

	// Read the run
	newRun, err := runs.Get(context.TODO(), client, run.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newRun)
	th.AssertEquals(t, run.Selector, newRun.Selector)
	th.AssertEquals(t, run.AutomationID, newRun.AutomationID)
}

func TestRunList(t *testing.T) {
	client, err := NewAutomationV1Client(context.TODO())
	th.AssertNoErr(t, err)

	count := 0
	var allRuns []runs.Run

	//nolint:errcheck
	runs.List(client, runs.ListOpts{PerPage: 1}).EachPage(context.TODO(), func(ctx context.Context, page pagination.Page) (bool, error) {
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
