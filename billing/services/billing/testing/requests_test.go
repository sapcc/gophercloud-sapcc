// Copyright 2020 SAP SE
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

package testing

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fake "github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/billing/services/billing"
)

var billingList = []billing.Billing{
	{
		Region:         "region",
		ProjectID:      "1a894ddae4274a32a81eee43e4e5d67e",
		ProjectName:    "my-project",
		ObjectID:       "1a894ddae4274a32a81eee43e4e5d67e",
		MetricType:     "compute_ram_quota",
		Amount:         12.1688,
		Duration:       2.968,
		PriceLoc:       4.1,
		PriceSec:       0,
		CostObject:     "",
		CostObjectType: "",
		COInherited:    true,
		SendCC:         123456789,
	},
	{
		Region:         "region",
		ProjectID:      "1a894ddae4274a32a81eee43e4e5d67e",
		ProjectName:    "my-project",
		ObjectID:       "1a894ddae4274a32a81eee43e4e5d67e",
		MetricType:     "network_loadbalancers_quota",
		Amount:         3.2648,
		Duration:       2.968,
		PriceLoc:       1.1,
		PriceSec:       0,
		CostObject:     "123",
		CostObjectType: "CC",
		COInherited:    false,
		SendCC:         123456789,
	},
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/services/billing", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	allBillings, err := billing.List(fake.ServiceClient(), nil).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	actual, err := billing.ExtractBillings(allBillings)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, billingList, actual)
}

func TestListOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/services/billing", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.CheckEquals(t, r.URL.Query().Get("exclude_missing_co"), "true")
		th.CheckEquals(t, r.URL.Query().Get("year"), "2019")
		th.CheckEquals(t, r.URL.Query().Get("from"), "2018-08-20T14:39:39.786")
		th.CheckEquals(t, r.URL.Query().Get("to"), "2019-08-20T14:39:39.786")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	listOpts := billing.ListOpts{
		ExcludeMissingCO: true,
		Year:             2019,
		From:             time.Date(2018, time.August, 20, 14, 39, 39, 786000000, time.UTC),
		To:               time.Date(2019, time.August, 20, 14, 39, 39, 786000000, time.UTC),
	}

	allBillings, err := billing.List(fake.ServiceClient(), listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	actual, err := billing.ExtractBillings(allBillings)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, billingList, actual)
}
