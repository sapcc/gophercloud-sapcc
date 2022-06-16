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

package costing

import (
	"github.com/gophercloud/gophercloud/pagination"
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
	return r.(CostingPage).Result.ExtractIntoSlicePtr(v, "")
}
