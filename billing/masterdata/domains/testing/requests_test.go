// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/billing/masterdata/domains"
)

var domainsList = []domains.Domain{
	{
		IID:         123,
		DomainID:    "707c94677ac741ecb1f2cabc804c1285",
		DomainName:  "master",
		Description: "example domain",
		CostObject: domains.CostObject{
			Name:               "1234567",
			Type:               "IO",
			ProjectsCanInherit: false,
		},
		ChangedBy:         "c48b0ce218848fd0bc78c8367ae9c40512024e2fc39451f47d9a62ad3ff41c26",
		ChangedAt:         time.Date(2019, time.January, 29, 9, 37, 58, 792000000, time.UTC),
		Collector:         "billing.region.local",
		Region:            "region",
		IsComplete:        false,
		MissingAttributes: "Primary contact not specified",
	},
}

var updateResponse = domains.Domain{
	IID:         123,
	DomainID:    "707c94677ac741ecb1f2cabc804c1285",
	DomainName:  "master",
	Description: "new example domain",
	CostObject: domains.CostObject{
		ProjectsCanInherit: true,
	},
	ResponsiblePrimaryContactID: "D123456",
	Collector:                   "billing.region.local",
	Region:                      "region",
	ChangedBy:                   "c48b0ce218848fd0bc78c8367ae9c40512024e2fc39451f47d9a62ad3ff41c26",
	ChangedAt:                   time.Date(2019, time.January, 29, 9, 37, 58, 792000000, time.UTC),
	IsComplete:                  true,
}

func TestList(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/masterdata/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	allDomains, err := domains.List(client.ServiceClient(fakeServer), domains.ListOpts{}).AllPages(t.Context())
	th.AssertNoErr(t, err)

	actual, err := domains.ExtractDomains(allDomains)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, domainsList, actual)
}

func TestListOpts(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/masterdata/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		th.CheckEquals(t, "true", r.URL.Query().Get("checkCOValidity"))
		th.CheckEquals(t, "true", r.URL.Query().Get("excludeDeleted"))
		th.CheckEquals(t, "2023-04-26T12:31:42.1337", r.URL.Query().Get("from"))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	listOpts := domains.ListOpts{
		CheckCOValidity: true,
		ExcludeDeleted:  true,
		From:            time.Date(2023, 04, 26, 12, 31, 42, 133700000, time.UTC),
	}

	allDomains, err := domains.List(client.ServiceClient(fakeServer), listOpts).AllPages(t.Context())
	th.AssertNoErr(t, err)

	actual, err := domains.ExtractDomains(allDomains)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, domainsList, actual)
}

func TestGet(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/masterdata/domains/707c94677ac741ecb1f2cabc804c1285", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	n, err := domains.Get(t.Context(), client.ServiceClient(fakeServer), "707c94677ac741ecb1f2cabc804c1285").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, domainsList[0])
}

func TestUpdate(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/masterdata/domains/707c94677ac741ecb1f2cabc804c1285", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})

	options := domains.UpdateOpts{
		Description:                    "new example domain",
		ResponsiblePrimaryContactEmail: "example@mail.com",
		CostObject: domains.CostObject{
			ProjectsCanInherit: true,
		},
		DomainID:                    "707c94677ac741ecb1f2cabc804c1285",
		DomainName:                  "master",
		ResponsiblePrimaryContactID: "D123456",
		Collector:                   "billing.region.local",
		Region:                      "region",
	}

	s, err := domains.Update(t.Context(), client.ServiceClient(fakeServer), "707c94677ac741ecb1f2cabc804c1285", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *s, updateResponse)
}
