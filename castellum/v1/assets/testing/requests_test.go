// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/majewsky/gg/option"
	"github.com/sapcc/go-api-declarations/castellum"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/assets"
)

const (
	projectID = "88e5cad3-38e6-454f-b412-662cda03e7a1"
	assetType = "nfs-shares"
	assetID   = "05620cba-c0c1-4e75-a5e9-b5decf643dc7"
)

var assetsList = []castellum.Asset{
	{
		UUID:         "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
		Size:         100,
		UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 75.5},
		MinimumSize:  option.Some(uint64(10)),
		MaximumSize:  option.Some(uint64(1000)),
		Checked:      option.Some(castellum.Checked{}),
		Stale:        false,
	},
	{
		UUID:         "5d7f5c1c-3f2e-4b0a-9e6d-8a1b2c3d4e5f",
		Size:         200,
		UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 20.0},
		Stale:        false,
	},
}

var singleAsset = castellum.Asset{
	UUID:         "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
	Size:         100,
	UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 75.5},
	MinimumSize:  option.Some(uint64(10)),
	MaximumSize:  option.Some(uint64(1000)),
	Checked:      option.Some(castellum.Checked{}),
	Stale:        false,
	PendingOperation: option.Some(castellum.StandaloneOperation{
		Operation: castellum.Operation{
			State:   castellum.OperationStateConfirmed,
			Reason:  castellum.OperationReasonHigh,
			OldSize: 100,
			NewSize: 120,
			Created: castellum.OperationCreation{
				AtUnix:       1700000000,
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 82.0},
			},
			Confirmed: option.Some(castellum.OperationConfirmation{AtUnix: 1700003600}),
			Greenlit:  option.Some(castellum.OperationGreenlight{AtUnix: 1700007200}),
		},
	}),
}

var singleAssetWithHistory = castellum.Asset{
	UUID:         "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
	Size:         100,
	UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 75.5},
	Stale:        false,
	FinishedOperations: []castellum.StandaloneOperation{
		{
			Operation: castellum.Operation{
				State:   castellum.OperationStateSucceeded,
				Reason:  castellum.OperationReasonHigh,
				OldSize: 80,
				NewSize: 100,
				Created: castellum.OperationCreation{
					AtUnix:       1699000000,
					UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 85.0},
				},
				Confirmed: option.Some(castellum.OperationConfirmation{AtUnix: 1699003600}),
				Greenlit:  option.Some(castellum.OperationGreenlight{AtUnix: 1699007200}),
				Finished:  option.Some(castellum.OperationFinish{AtUnix: 1699010800}),
			},
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
