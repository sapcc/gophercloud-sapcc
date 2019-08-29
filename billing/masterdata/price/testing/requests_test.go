package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/sapcc/gophercloud-billing/billing/masterdata/price"
)

var priceList = []price.Price{
	{
		SendCC:              123456789,
		CostElement:         123456,
		PriceLoc:            0.123456,
		PriceSec:            0,
		ValidFrom:           time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC),
		ValidTo:             time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC),
		MetricType:          "foo",
		Region:              "region",
		ValidForProjectType: "quotaUsage",
		ObjectType:          "object",
	},
	{
		SendCC:              123456789,
		CostElement:         123457,
		PriceLoc:            0.023456,
		PriceSec:            0,
		ValidFrom:           time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC),
		ValidTo:             time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC),
		MetricType:          "bar",
		Region:              "region",
		ValidForProjectType: "quotaUsage",
		ObjectType:          "object",
	},
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/pricelist", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	allPrices, err := price.List(fake.ServiceClient(), price.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	actual, err := price.ExtractPrices(allPrices)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, priceList, actual)
}

func TestListOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/pricelist", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.CheckEquals(t, r.URL.Query().Get("onlyActive"), "true")
		th.CheckEquals(t, r.URL.Query().Get("METRIC_TYPE"), "foo")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	listOpts := price.ListOpts{
		OnlyActive: true,
		MetricType: "foo",
	}

	allPrices, err := price.List(fake.ServiceClient(), listOpts).AllPages()
	th.AssertNoErr(t, err)

	actual, err := price.ExtractPrices(allPrices)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, priceList, actual)
}

func TestDateListOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/price/region/foo/2018-08-20T14:39:39/2019-08-20T14:39:39", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	listOpts := price.ListOpts{
		Region:     "region",
		MetricType: "foo",
		From:       time.Date(2018, time.August, 20, 14, 39, 39, 786000000, time.UTC),
		To:         time.Date(2019, time.August, 20, 14, 39, 39, 786000000, time.UTC),
	}

	allPrices, err := price.List(fake.ServiceClient(), listOpts).AllPages()
	th.AssertNoErr(t, err)

	actual, err := price.ExtractPrices(allPrices)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, priceList, actual)
}

func TestRegionListOptsNoDate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/masterdata/price/region/foo/0001-01-01T00:00:00/9999-12-31T00:00:00", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	listOpts := price.ListOpts{
		Region:     "region",
		MetricType: "foo",
	}

	allPrices, err := price.List(fake.ServiceClient(), listOpts).AllPages()
	th.AssertNoErr(t, err)

	actual, err := price.ExtractPrices(allPrices)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, priceList, actual)
}
