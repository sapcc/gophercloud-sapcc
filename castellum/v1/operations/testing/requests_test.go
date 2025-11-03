// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/operations"
)

const (
	projectID = "88e5cad3-38e6-454f-b412-662cda03e7a1"
	assetType = "nfs-shares"
	assetID   = "05620cba-c0c1-4e75-a5e9-b5decf643dc7"
)

var pendingOp = operations.Operation{
	ProjectID: projectID,
	AssetType: assetType,
	AssetID:   assetID,
	State:     "confirmed",
	Reason:    "high",
	OldSize:   100,
	NewSize:   120,
	Created:   operations.OperationEvent{At: 1700000000, UsagePercent: 82.0},
	Confirmed: &operations.OperationEvent{At: 1700003600},
	Greenlit:  &operations.GreenlitEvent{At: 1700007200},
}

var failedOp = operations.Operation{
	ProjectID: projectID,
	AssetType: assetType,
	AssetID:   assetID,
	State:     "failed",
	Reason:    "high",
	OldSize:   100,
	NewSize:   120,
	Created:   operations.OperationEvent{At: 1700000000, UsagePercent: 82.0},
	Confirmed: &operations.OperationEvent{At: 1700003600},
	Greenlit:  &operations.GreenlitEvent{At: 1700007200},
	Finished:  &operations.FinishedEvent{At: 1700010800, Error: "quota exceeded"},
}

var succeededOp = operations.Operation{
	ProjectID: projectID,
	AssetType: assetType,
	AssetID:   assetID,
	State:     "succeeded",
	Reason:    "high",
	OldSize:   80,
	NewSize:   100,
	Created:   operations.OperationEvent{At: 1699000000, UsagePercent: 85.0},
	Confirmed: &operations.OperationEvent{At: 1699003600},
	Greenlit:  &operations.GreenlitEvent{At: 1699007200},
	Finished:  &operations.FinishedEvent{At: 1699010800},
}

func TestListPending(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/operations/pending", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListPendingResponse)
	})

	result, err := operations.ListPending(t.Context(), client.ServiceClient(fakeServer), nil).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{pendingOp})
}

func TestListPendingWithOpts(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/operations/pending", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestFormValues(t, r, map[string]string{"project": projectID})

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListPendingResponse)
	})

	result, err := operations.ListPending(t.Context(), client.ServiceClient(fakeServer), operations.ListOpts{Project: projectID}).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{pendingOp})
}

func TestListRecentlyFailed(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/operations/recently-failed", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListRecentlyFailedResponse)
	})

	result, err := operations.ListRecentlyFailed(t.Context(), client.ServiceClient(fakeServer), nil).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{failedOp})
}

func TestListRecentlySucceeded(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/operations/recently-succeeded", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListRecentlySucceededResponse)
	})

	result, err := operations.ListRecentlySucceeded(t.Context(), client.ServiceClient(fakeServer), nil).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{succeededOp})
}

func TestListProjectPending(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/resources/%s/operations/pending", projectID, assetType), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListPendingResponse)
	})

	result, err := operations.ListProjectPending(t.Context(), client.ServiceClient(fakeServer), projectID, assetType, nil).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{pendingOp})
}

func TestListProjectRecentlyFailed(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/resources/%s/operations/recently-failed", projectID, assetType), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListRecentlyFailedResponse)
	})

	result, err := operations.ListProjectRecentlyFailed(t.Context(), client.ServiceClient(fakeServer), projectID, assetType, nil).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{failedOp})
}

func TestListProjectRecentlySucceeded(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc(fmt.Sprintf("/projects/%s/resources/%s/operations/recently-succeeded", projectID, assetType), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListRecentlySucceededResponse)
	})

	result, err := operations.ListProjectRecentlySucceeded(t.Context(), client.ServiceClient(fakeServer), projectID, assetType, nil).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, []operations.Operation{succeededOp})
}
