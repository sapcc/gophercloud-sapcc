package v1

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/sapcc/gophercloud-sapcc/arc/v1/agents"
)

func TestAgentInit(t *testing.T) {
	client, err := NewArcV1Client()
	th.AssertNoErr(t, err)

	cloudConfig, err := InitAgent(t, client, "text/cloud-config")
	th.AssertNoErr(t, err)

	tools.PrintResource(t, *cloudConfig)
	if !strings.Contains(*cloudConfig, "#cloud-config") {
		t.Fatalf("Result doesn't contain '#cloud-config'")
	}

	jsonConfig, err := InitAgent(t, client, "application/json")
	th.AssertNoErr(t, err)

	var initJson agents.InitJSON
	err = json.Unmarshal([]byte(*jsonConfig), &initJson)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, initJson)
}

func TestAgentList(t *testing.T) {
	client, err := NewArcV1Client()
	th.AssertNoErr(t, err)

	var count int
	var allAgents []agents.Agent

	//nolint:errcheck
	agents.List(client, agents.ListOpts{PerPage: 1}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		tmp, err := agents.ExtractAgents(page)
		if err != nil {
			t.Errorf("Failed to extract agents: %v", err)
			return false, nil
		}

		allAgents = append(allAgents, tmp...)

		return true, nil
	})

	tools.PrintResource(t, allAgents)

	expectedPages := 2
	if count < expectedPages {
		t.Errorf("Expected %d page, got %d", expectedPages, count)
	}

	expectedAgents := 2
	if len(allAgents) < expectedAgents {
		t.Errorf("Expected %d agents, got %d", expectedAgents, len(allAgents))
	}
}

// TODO test tags CRUD
// TODO test facts Read
