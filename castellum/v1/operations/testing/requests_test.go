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

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/operations"
)

const (
	projectID = "88e5cad3-38e6-454f-b412-662cda03e7a1"
	assetType = "nfs-shares"
	assetID   = "05620cba-c0c1-4e75-a5e9-b5decf643dc7"
)

var pendingOp = castellum.StandaloneOperation{
	ProjectUUID: projectID,
	AssetType:   assetType,
	AssetID:     assetID,
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
}

var failedOp = castellum.StandaloneOperation{
	ProjectUUID: projectID,
	AssetType:   assetType,
	AssetID:     assetID,
	Operation: castellum.Operation{
		State:   castellum.OperationStateFailed,
		Reason:  castellum.OperationReasonHigh,
		OldSize: 100,
		NewSize: 120,
		Created: castellum.OperationCreation{
			AtUnix:       1700000000,
			UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 82.0},
		},
		Confirmed: option.Some(castellum.OperationConfirmation{AtUnix: 1700003600}),
		Greenlit:  option.Some(castellum.OperationGreenlight{AtUnix: 1700007200}),
		Finished:  option.Some(castellum.OperationFinish{AtUnix: 1700010800, ErrorMessage: "quota exceeded"}),
	},
}

var succeededOp = castellum.StandaloneOperation{
	ProjectUUID: projectID,
	AssetType:   assetType,
	AssetID:     assetID,
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{pendingOp})
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{pendingOp})
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{failedOp})
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{succeededOp})
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{pendingOp})
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{failedOp})
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
	th.AssertDeepEquals(t, result, []castellum.StandaloneOperation{succeededOp})
}
