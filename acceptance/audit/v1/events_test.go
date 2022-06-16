// Copyright 2019 SAP SE
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

	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/sapcc/gophercloud-sapcc/audit/v1/events"
)

func TestEventList(t *testing.T) {
	client, err := NewHermesV1Client()
	th.AssertNoErr(t, err)

	var count int
	var allEvents []events.Event

	//nolint:errcheck
	events.List(client, events.ListOpts{Limit: 5000}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		tmp, err := events.ExtractEvents(page)
		if err != nil {
			t.Errorf("Failed to extract events: %v", err)
			return false, nil
		}

		allEvents = append(allEvents, tmp...)

		return true, nil
	})

	tools.PrintResource(t, allEvents)

	expectedPages := 2
	if count < expectedPages {
		t.Errorf("Expected %d page, got %d", expectedPages, count)
	}

	expectedEvents := 2
	if len(allEvents) < expectedEvents {
		t.Errorf("Expected %d events, got %d", expectedEvents, len(allEvents))
	}
}
