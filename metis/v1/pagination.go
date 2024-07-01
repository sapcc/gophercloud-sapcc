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

package v1

import "github.com/gophercloud/gophercloud/v2/pagination"

// CommonPage is the common base type for all pagination pages used by Metis V1 endpoints.
type CommonPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of projects contains any results.
func (p CommonPage) IsEmpty() (bool, error) {
	items, err := extractItems(p)
	if err != nil {
		return true, err
	}
	return len(items) == 0, nil
}

// NextPageURL extracts the "next" link from the response.
func (p CommonPage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}

	err := p.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, err
}

func extractItems(p pagination.Page) ([]any, error) {
	var s struct {
		Items []any `json:"items"`
	}
	if err := (p.(CommonPage)).ExtractInto(&s); err != nil {
		return []any{}, err
	}
	return s.Items, nil
}

// CreatePage is used to use pagination.NewPager with any of the Metis V1 endpoints.
func CreatePage() func(pagination.PageResult) pagination.Page {
	return func(r pagination.PageResult) pagination.Page {
		// The structure of the Metis API response has the list of projects not on the top level, but nested in the "data" field.
		// pagination.Pager does not work with this structure, so we extract the list of projects and the next page URL.
		// The body is then replaced with a map containing the list of projects and the next page URL, which is then
		// compatible with the pagination.
		var source struct {
			Data struct {
				Items []any  `json:"items"`
				Next  string `json:"nextLink"`
			} `json:"data"`
		}
		if err := r.ExtractInto(&source); err != nil {
			return nil
		}
		newBody := make(map[string]any, 2)
		newBody["items"] = source.Data.Items
		newBody["next"] = source.Data.Next
		r.Body = newBody
		return CommonPage{pagination.LinkedPageBase{PageResult: r}}
	}
}
