// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package price

import (
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Price represents a Billing Price.
type Price struct {
	SendCC              int       `json:"SEND_CC"`
	CostElement         int       `json:"COST_ELEMENT"`
	PriceLoc            float64   `json:"PRICE_LOC,string"`
	PriceSec            float64   `json:"PRICE_SEC,string"`
	ValidForProjectType string    `json:"VALID_FOR_PROJECT_TYPE"`
	ObjectType          string    `json:"OBJECT_TYPE"`
	MetricType          string    `json:"METRIC_TYPE"`
	Region              string    `json:"REGION"`
	ValidFrom           time.Time `json:"-"`
	ValidTo             time.Time `json:"-"`
}

func (r *Price) UnmarshalJSON(b []byte) error {
	type tmp Price
	var s struct {
		tmp
		ValidFrom gophercloud.JSONRFC3339NoZ `json:"VALID_FROM"`
		ValidTo   gophercloud.JSONRFC3339NoZ `json:"VALID_TO"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Price(s.tmp)

	r.ValidFrom = time.Time(s.ValidFrom)
	r.ValidTo = time.Time(s.ValidTo)

	return nil
}

// PricePage is the page returned by a pager when traversing over a collection
// of price.
type PricePage struct {
	pagination.SinglePageBase
}

// ExtractPrices accepts a Page struct, specifically a PricePage
// struct, and extracts the elements into a slice of Price structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractPrices(r pagination.Page) ([]Price, error) {
	var s []Price
	err := ExtractPricesInto(r, &s)
	return s, err
}

func ExtractPricesInto(r pagination.Page, v any) error {
	return r.(PricePage).ExtractIntoSlicePtr(v, "")
}
