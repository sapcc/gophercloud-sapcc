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

package projects

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"

	v1 "github.com/sapcc/gophercloud-sapcc/metis/v1"
)

// ProjectPage
type ProjectPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of projects contains any results.
func (r ProjectPage) IsEmpty() (bool, error) {
	projects, err := ExtractProjects(r)
	if err != nil {
		return true, err
	}
	return len(projects) == 0, nil
}

// NextPageURL extracts the "next" link from the response.
func (r ProjectPage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}

	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, err
}

// ExtractProjects accepts a Page struct, specifically an ProjectsPage struct,
// and extracts the elements into a slice of Projects structs.
func ExtractProjects(r pagination.Page) ([]Project, error) {
	var s struct {
		Projects []Project `json:"items"`
	}
	if err := (r.(v1.CommonPage)).ExtractInto(&s); err != nil {
		return []Project{}, err
	}
	return s.Projects, nil
}

// GetResult represents the result of a get operation. Call its Extract method
// to interpret it as a Project.
type GetResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a project
// resource.
func (r GetResult) Extract() (*Project, error) {
	var s struct {
		Data struct {
			Items []Project `json:"items"`
		} `json:"data"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	if s.Data.Items == nil {
		return nil, nil
	}
	return &s.Data.Items[0], nil
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}
