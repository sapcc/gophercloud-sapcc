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

	"github.com/sapcc/gophercloud-sapcc/billing/services/costing"
)

var costingList = []costing.Costing{
	{
		Year:           2019,
		Month:          8,
		Region:         "region",
		ProjectID:      "1a894ddae4274a32a81eee43e4e5d67e",
		ObjectID:       "1a894ddae4274a32a81eee43e4e5d67e",
		CostObject:     "1234567",
		CostObjectType: "IO",
		COInherited:    false,
		AllocationType: "usable",
		Service:        "blockStorage",
		Measure:        "capacity",
		Amount:         671930,
		AmountUnit:     "GiBh",
		Duration:       671.93,
		DurationUnit:   "h",
		PriceLoc:       67.193,
		PriceSec:       0,
		Currency:       "EUR",
	},
	{
		Year:           2019,
		Month:          8,
		Region:         "region",
		ProjectID:      "1a894ddae4274a32a81eee43e4e5d67e",
		ObjectID:       "29940f04-961a-4903-a4c5-d91e750acc7f",
		CostObject:     "1234567",
		CostObjectType: "IO",
		COInherited:    false,
		AllocationType: "provisioned",
		Service:        "virtual",
		Measure:        "os_suse",
		Amount:         1.00299,
		AmountUnit:     "pieceh",
		Duration:       1.00299,
		DurationUnit:   "h",
		PriceLoc:       0.0001,
		PriceSec:       0,
		Currency:       "EUR",
	},
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/services/costing/objects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	allCostings, err := costing.ListObjects(fake.ServiceClient(), nil).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	actual, err := costing.ExtractCostings(allCostings)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, costingList, actual)
}

func TestListOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/services/costing/objects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.CheckEquals(t, r.URL.Query().Get("exclude_internal_co"), "true")
		th.CheckEquals(t, r.URL.Query().Get("last"), "12")
		th.CheckEquals(t, r.URL.Query().Get("start"), "2018-08-20T14:39:39.786")
		th.CheckEquals(t, r.URL.Query().Get("end"), "2019-08-20T14:39:39.786")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	listOpts := costing.ListOpts{
		ExcludeInternalCO: true,
		Last:              12,
		Start:             time.Date(2018, time.August, 20, 14, 39, 39, 786000000, time.UTC),
		End:               time.Date(2019, time.August, 20, 14, 39, 39, 786000000, time.UTC),
	}

	allCostings, err := costing.ListObjects(fake.ServiceClient(), listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	actual, err := costing.ExtractCostings(allCostings)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, costingList, actual)
}
