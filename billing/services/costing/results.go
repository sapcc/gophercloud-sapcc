// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package costing

import (
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Costing represents a Costing Costing.
type Costing struct {
	Year           int     `json:"year"`
	Month          int     `json:"month"`
	Region         string  `json:"region"`
	ProjectID      string  `json:"project_id"`
	ObjectID       string  `json:"object_id"`
	CostObject     string  `json:"cost_object"`
	CostObjectType string  `json:"cost_object_type"`
	COInherited    bool    `json:"co_inherited"`
	AllocationType string  `json:"allocation_type"`
	Service        string  `json:"service"`
	Measure        string  `json:"measure"`
	Amount         float64 `json:"amount"`
	AmountUnit     string  `json:"amount_unit"`
	Duration       float64 `json:"duration"`
	DurationUnit   string  `json:"duration_unit"`
	PriceLoc       float64 `json:"price_loc"`
	PriceSec       float64 `json:"price_sec"`
	Currency       string  `json:"currency"`
}

// CostingPage is the page returned by a pager when traversing over a collection
// of costing.
type CostingPage struct {
	pagination.SinglePageBase
}

// ExtractCostings accepts a Page struct, specifically a CostingPage
// struct, and extracts the elements into a slice of Costing structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractCostings(r pagination.Page) ([]Costing, error) {
	var s []Costing
	err := ExtractCostingsInto(r, &s)
	return s, err
}

func ExtractCostingsInto(r pagination.Page, v interface{}) error {
	return r.(CostingPage).ExtractIntoSlicePtr(v, "")
}
