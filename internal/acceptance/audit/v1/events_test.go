// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"testing"

	"github.com/sapcc/gophercloud-sapcc/v2/internal/acceptance/tools"

	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/audit/v1/events"
)

func TestEventList(t *testing.T) {
	client, err := NewHermesV1Client(t.Context())
	th.AssertNoErr(t, err)

	var count int
	var allEvents []events.Event

	//nolint:errcheck
	events.List(client, events.ListOpts{Limit: 5000}).EachPage(t.Context(), func(ctx context.Context, page pagination.Page) (bool, error) {
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
