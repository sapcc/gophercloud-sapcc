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
	"github.com/majewsky/gg/option"
	"github.com/sapcc/go-api-declarations/castellum"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/resources"
)

var resourcesList = []resources.Resource{
	{
		ResourceType: "nfs-shares",
		Resource: castellum.Resource{
			AssetCount: 42,
			Checked:    option.Some(castellum.Checked{ErrorMessage: "cannot connect to OpenStack"}),
			LowThreshold: option.Some(castellum.Threshold{
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 20.0},
				DelaySeconds: 3600,
			}),
			HighThreshold: option.Some(castellum.Threshold{
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 80.0},
				DelaySeconds: 1800,
			}),
			CriticalThreshold: option.Some(castellum.Threshold{
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 95.0},
			}),
			SizeConstraints: option.Some(castellum.SizeConstraints{
				Minimum: option.Some(uint64(10)),
				Maximum: option.Some(uint64(2000)),
			}),
			SizeSteps: castellum.SizeSteps{Percent: 20.0},
		},
	},
	{
		ResourceType: "smb-shares",
		Resource: castellum.Resource{
			AssetCount: 42,
			Checked:    option.Some(castellum.Checked{ErrorMessage: "cannot connect to OpenStack"}),
			LowThreshold: option.Some(castellum.Threshold{
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 10.0},
				DelaySeconds: 3600,
			}),
			HighThreshold: option.Some(castellum.Threshold{
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 50.0},
				DelaySeconds: 1800,
			}),
			CriticalThreshold: option.Some(castellum.Threshold{
				UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 90.0},
			}),
			SizeConstraints: option.Some(castellum.SizeConstraints{
				Minimum: option.Some(uint64(20)),
				Maximum: option.Some(uint64(2000)),
			}),
			SizeSteps: castellum.SizeSteps{Percent: 10.0},
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
		LowThreshold: &castellum.Threshold{
			UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 20.0},
			DelaySeconds: 3600,
		},
		HighThreshold: &castellum.Threshold{
			UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 80.0},
			DelaySeconds: 1800,
		},
		CriticalThreshold: &castellum.Threshold{
			UsagePercent: castellum.UsageValues{castellum.SingularUsageMetric: 95.0},
		},
		SizeConstraints: &castellum.SizeConstraints{
			Minimum: option.Some(uint64(10)),
			Maximum: option.Some(uint64(2000)),
		},
		SizeSteps: &castellum.SizeSteps{Percent: 20.0},
	}

	err := resources.Create(t.Context(), client.ServiceClient(fakeServer), "88e5cad3-38e6-454f-b412-662cda03e7a1", "nfs-shares", options).ExtractErr()
	th.AssertNoErr(t, err)
}
