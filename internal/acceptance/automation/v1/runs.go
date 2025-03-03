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
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/automation/v1/runs"
)

// CreateRun will create an automation run. An error will be
// returned if the run could not be created.
func CreateRun(t *testing.T, client *gophercloud.ServiceClient, automationID string) (*runs.Run, error) {
	t.Logf("Attempting to create run for the %s automation ID", automationID)

	createOpts := runs.CreateOpts{
		AutomationID: automationID,
		Selector:     "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
	}

	s, err := runs.Create(t.Context(), client, createOpts).Extract()
	if err != nil {
		return s, err
	}

	t.Logf("Successfully created run for the %s automation ID", automationID)

	th.AssertEquals(t, s.AutomationID, automationID)
	th.AssertEquals(t, s.Selector, "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'")
	th.AssertEquals(t, s.State, "preparing")

	return s, nil
}
