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

package billing

import (
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Bool allows 0/1 to also become boolean.
type Bool bool

func (v *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*v = Bool(txt == "1" || txt == "true")
	return nil
}

// Billing represents a Billing Billing.
type Billing struct {
	Region         string  `json:"REGION"`
	ProjectID      string  `json:"PROJECT_ID"`
	ProjectName    string  `json:"PROJECT_NAME"`
	ObjectID       string  `json:"OBJECT_ID"`
	MetricType     string  `json:"METRIC_TYPE"`
	Amount         float64 `json:"AMOUNT,string"`
	Duration       float64 `json:"DURATION,string"`
	PriceLoc       float64 `json:"PRICE_LOC,string"`
	PriceSec       float64 `json:"PRICE_SEC,string"`
	CostObject     string  `json:"COST_OBJECT"`
	CostObjectType string  `json:"COST_OBJECT_TYPE"`
	COInherited    Bool    `json:"CO_INHERITED"`
	SendCC         int     `json:"SEND_CC"`
}

// BillingPage is the page returned by a pager when traversing over a collection
// of billing.
type BillingPage struct {
	pagination.SinglePageBase
}

// ExtractBillings accepts a Page struct, specifically a BillingPage
// struct, and extracts the elements into a slice of Billing structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractBillings(r pagination.Page) ([]Billing, error) {
	var s []Billing
	err := ExtractBillingsInto(r, &s)
	return s, err
}

func ExtractBillingsInto(r pagination.Page, v interface{}) error {
	return r.(BillingPage).ExtractIntoSlicePtr(v, "")
}
