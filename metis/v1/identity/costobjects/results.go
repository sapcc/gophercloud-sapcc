// Copyright 2023 SAP SE
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

package costobjects

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"

	v1 "github.com/sapcc/gophercloud-sapcc/v2/metis/v1"
)

type CostObject struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Extract accepts a Page struct, specifically an v1.CommonPage struct,
// and extracts the elements into a slice of CostObjects structs.
func Extract(r pagination.Page) ([]CostObject, error) {
	var s struct {
		CostObjects []CostObject `json:"items"`
	}
	if err := (r.(v1.CommonPage)).ExtractInto(&s); err != nil {
		return nil, err
	}
	return s.CostObjects, nil
}

// GetResult represents the result of a get operation. Call its Extract method
// to interpret it as a CostObject.
type GetResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a CostObject
// resource.
func (r GetResult) Extract() (*CostObject, error) {
	var s struct {
		Data struct {
			Item CostObject `json:"item"`
		} `json:"data"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data.Item, nil
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.ExtractIntoStructPtr(v, "")
}
