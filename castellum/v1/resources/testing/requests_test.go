// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"sort"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/resources"
)

func intPtr(v int) *int { return &v }

var resourcesList = []resources.Resource{
	{
		ResourceType: "nfs-shares",
		AssetCount:   42,
		Checked:      &resources.ResourceChecked{Error: "cannot connect to OpenStack"},
		LowThreshold: &resources.Threshold{
			UsagePercent: 20.0,
			DelaySeconds: 3600,
		},
		HighThreshold: &resources.Threshold{
			UsagePercent: 80.0,
			DelaySeconds: 1800,
		},
		CriticalThreshold: &resources.Threshold{
			UsagePercent: 95.0,
		},
		SizeConstraints: &resources.SizeConstraints{
			Minimum: intPtr(10),
			Maximum: intPtr(2000),
		},
		SizeSteps: &resources.SizeSteps{
			Percent: 20.0,
		},
	},
	{
		ResourceType: "smb-shares",
		AssetCount:   42,
		Checked:      &resources.ResourceChecked{Error: "cannot connect to OpenStack"},
		LowThreshold: &resources.Threshold{
			UsagePercent: 10.0,
			DelaySeconds: 3600,
		},
		HighThreshold: &resources.Threshold{
			UsagePercent: 50.0,
			DelaySeconds: 1800,
		},
		CriticalThreshold: &resources.Threshold{
			UsagePercent: 90.0,
		},
		SizeConstraints: &resources.SizeConstraints{
			Minimum: intPtr(20),
			Maximum: intPtr(2000),
		},
		SizeSteps: &resources.SizeSteps{
			Percent: 10.0,
		},
	},
}

func TestList(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/88e5cad3-38e6-454f-b412-662cda03e7a1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	n, err := resources.List(t.Context(), client.ServiceClient(fakeServer), "88e5cad3-38e6-454f-b412-662cda03e7a1", resources.ListOpts{}).Extract()
	th.AssertNoErr(t, err)

	sort.Slice(n, func(i, j int) bool {
		return n[i].ResourceType < n[j].ResourceType
	})

	sort.Slice(resourcesList, func(i, j int) bool {
		return resourcesList[i].ResourceType < resourcesList[j].ResourceType
	})

	th.AssertDeepEquals(t, n, resourcesList)
}

func TestGet(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/88e5cad3-38e6-454f-b412-662cda03e7a1/resources/nfs-shares", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	n, err := resources.Get(t.Context(), client.ServiceClient(fakeServer), "88e5cad3-38e6-454f-b412-662cda03e7a1", "nfs-shares").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, resourcesList[0])
}

func TestDelete(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/88e5cad3-38e6-454f-b412-662cda03e7a1/resources/nfs-shares", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := resources.Delete(t.Context(), client.ServiceClient(fakeServer), "88e5cad3-38e6-454f-b412-662cda03e7a1", "nfs-shares").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestCreate(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/88e5cad3-38e6-454f-b412-662cda03e7a1/resources/nfs-shares", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprint(w, CreateResponse)
	})

	options := resources.CreateOpts{
		LowThreshold: &resources.Threshold{
			UsagePercent: 20.0,
			DelaySeconds: 3600,
		},
		HighThreshold: &resources.Threshold{
			UsagePercent: 80.0,
			DelaySeconds: 1800,
		},
		CriticalThreshold: &resources.Threshold{
			UsagePercent: 95.0,
		},
		SizeConstraints: &resources.SizeConstraints{
			Minimum: intPtr(10),
			Maximum: intPtr(2000),
		},
		SizeSteps: &resources.SizeSteps{
			Percent: 20.0,
		},
	}

	err := resources.Create(t.Context(), client.ServiceClient(fakeServer), "88e5cad3-38e6-454f-b412-662cda03e7a1", "nfs-shares", options).ExtractErr()
	th.AssertNoErr(t, err)
}
