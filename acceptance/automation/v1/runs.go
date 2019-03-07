package automation

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/kayrus/gophercloud-lyra/automation/v1/runs"
)

// CreateRun will create an automation run. An error will be
// returned if the run could not be created.
func CreateRun(t *testing.T, client *gophercloud.ServiceClient, automationID string) (*runs.Run, error) {
	t.Logf("Attempting to create run for the %s automation ID", automationID)

	createOpts := runs.CreateOpts{
		AutomationID: automationID,
		Selector:     "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
	}

	s, err := runs.Create(client, createOpts).Extract()
	if err != nil {
		return s, err
	}

	t.Logf("Successfully created run for the %s automation ID", automationID)

	th.AssertEquals(t, s.AutomationID, automationID)
	th.AssertEquals(t, s.Selector, "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'")
	th.AssertEquals(t, s.State, "preparing")

	return s, nil
}
