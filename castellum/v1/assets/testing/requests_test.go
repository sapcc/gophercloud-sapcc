// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/assets"
)

func intPtr(v int) *int { return &v }

const (
	projectID = "88e5cad3-38e6-454f-b412-662cda03e7a1"
	assetType = "nfs-shares"
	assetID   = "05620cba-c0c1-4e75-a5e9-b5decf643dc7"
)

var assetsList = []assets.Asset{
	{
		ID:           "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
		Size:         100,
		UsagePercent: 75.5,
		MinSize:      intPtr(10),
		MaxSize:      intPtr(1000),
		Checked:      &assets.AssetChecked{Error: ""},
		Stale:        false,
	},
	{
		ID:           "5d7f5c1c-3f2e-4b0a-9e6d-8a1b2c3d4e5f",
		Size:         200,
		UsagePercent: 20.0,
		Stale:        false,
	},
}

var singleAsset = assets.Asset{
	ID:           "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
	Size:         100,
	UsagePercent: 75.5,
	MinSize:      intPtr(10),
	MaxSize:      intPtr(1000),
	Checked:      &assets.AssetChecked{Error: ""},
	Stale:        false,
	PendingOperation: &assets.Operation{
		State:     "confirmed",
		Reason:    "high",
		OldSize:   100,
		NewSize:   120,
		Created:   assets.OperationEvent{At: 1700000000, UsagePercent: 82.0},
		Confirmed: &assets.OperationEvent{At: 1700003600},
		Greenlit:  &assets.GreenlitEvent{At: 1700007200},
	},
}

var singleAssetWithHistory = assets.Asset{
	ID:           "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
	Size:         100,
	UsagePercent: 75.5,
	Stale:        false,
	FinishedOperations: []assets.Operation{
		{
			State:     "succeeded",
			Reason:    "high",
			OldSize:   80,
			NewSize:   100,
			Created:   assets.OperationEvent{At: 1699000000, UsagePercent: 85.0},
			Confirmed: &assets.OperationEvent{At: 1699003600},
			Greenlit:  &assets.GreenlitEvent{At: 1699007200},
			Finished:  &assets.FinishedEvent{At: 1699010800},
		},
	},
}

func TestList(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/assets/%s", projectID, assetType), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	result, err := assets.List(t.Context(), client.ServiceClient(fakeServer), projectID, assetType).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, assetsList)
}

func TestGet(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/assets/%s/%s", projectID, assetType, assetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	result, err := assets.Get(t.Context(), client.ServiceClient(fakeServer), projectID, assetType, assetID, false).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, *result, singleAsset)
}

func TestGetWithHistory(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/assets/%s/%s", projectID, assetType, assetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		if r.URL.Query().Get("history") == "" && r.URL.RawQuery != "history" {
			t.Errorf("Expected ?history query parameter")
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetHistoryResponse)
	})

	result, err := assets.Get(t.Context(), client.ServiceClient(fakeServer), projectID, assetType, assetID, true).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, *result, singleAssetWithHistory)
}

func TestResolveError(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/assets/%s/%s/error-resolved", projectID, assetType, assetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := assets.ResolveError(t.Context(), client.ServiceClient(fakeServer), projectID, assetType, assetID).ExtractErr()
	th.AssertNoErr(t, err)
}
