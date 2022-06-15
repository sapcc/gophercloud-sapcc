package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/billing/masterdata/domains"
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
		ResponsibleControllerID:    "D123456",
		ResponsibleControllerEmail: "example@mail.com",
		ChangedBy:                  "c48b0ce218848fd0bc78c8367ae9c40512024e2fc39451f47d9a62ad3ff41c26",
		ChangedAt:                  time.Date(2019, time.January, 29, 9, 37, 58, 792000000, time.UTC),
		Collector:                  "billing.region.local",
		Region:                     "region",
		IsComplete:                 false,
		MissingAttributes:          "Primary contact not specified",
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
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	allDomains, err := domains.List(fake.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := domains.ExtractDomains(allDomains)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, domainsList, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/domains/707c94677ac741ecb1f2cabc804c1285", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := domains.Get(fake.ServiceClient(), "707c94677ac741ecb1f2cabc804c1285").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *n, domainsList[0])
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/domains/707c94677ac741ecb1f2cabc804c1285", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
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

	s, err := domains.Update(fake.ServiceClient(), "707c94677ac741ecb1f2cabc804c1285", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *s, updateResponse)
}
