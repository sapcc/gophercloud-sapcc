// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/sapcc/go-api-declarations/castellum"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/admin"
)

const (
	domainID  = "d7a35a2e-3b6a-4b3c-8d5e-9f0a1b2c3d4e"
	projectID = "88e5cad3-38e6-454f-b412-662cda03e7a1"
	assetID   = "05620cba-c0c1-4e75-a5e9-b5decf643dc7"
	assetType = "nfs-shares"
)

func TestGetResourceScrapeErrors(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/admin/resource-scrape-errors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ResourceScrapeErrorsResponse)
	})

	expected := []castellum.ResourceScrapeError{
		{
			AssetType:   assetType,
			Checked:     castellum.Checked{ErrorMessage: "cannot connect to backend"},
			DomainUUID:  domainID,
			ProjectUUID: projectID,
		},
	}

	result, err := admin.GetResourceScrapeErrors(t.Context(), client.ServiceClient(fakeServer)).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, expected)
}

func TestGetAssetScrapeErrors(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/admin/asset-scrape-errors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, AssetScrapeErrorsResponse)
	})

	expected := []castellum.AssetScrapeError{
		{
			AssetUUID:   assetID,
			AssetType:   assetType,
			Checked:     castellum.Checked{ErrorMessage: "share not found"},
			DomainUUID:  domainID,
			ProjectUUID: projectID,
		},
	}

	result, err := admin.GetAssetScrapeErrors(t.Context(), client.ServiceClient(fakeServer)).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, expected)
}

func TestGetAssetResizeErrors(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/admin/asset-resize-errors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, AssetResizeErrorsResponse)
	})

	expected := []castellum.AssetResizeError{
		{
			AssetUUID:   assetID,
			AssetType:   assetType,
			DomainUUID:  domainID,
			ProjectUUID: projectID,
			OldSize:     100,
			NewSize:     120,
			Finished:    castellum.OperationFinish{AtUnix: 1700010800, ErrorMessage: "quota exceeded"},
		},
	}

	result, err := admin.GetAssetResizeErrors(t.Context(), client.ServiceClient(fakeServer)).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, result, expected)
}
