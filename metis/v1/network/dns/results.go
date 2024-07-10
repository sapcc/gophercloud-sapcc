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

package dns

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"

	v1 "github.com/sapcc/gophercloud-sapcc/v2/metis/v1"
)

type Zone struct {
	UUID               string            `json:"uuid,omitempty"`
	Name               string            `json:"name,omitempty"`
	Description        string            `json:"description,omitempty"`
	Email              string            `json:"email,omitempty"`
	Serial             int               `json:"serial,omitempty"`
	ParentZoneID       string            `json:"parentZoneId,omitempty"`
	ParentZoneName     string            `json:"parentZoneName,omitempty"`
	Pool               string            `json:"pool,omitempty"`
	PoolDescription    string            `json:"poolDescription,omitempty"`
	TTL                int               `json:"ttl,omitempty"`
	Status             string            `json:"status,omitempty"`
	Action             string            `json:"action,omitempty"`
	Type               string            `json:"type,omitempty"`
	Attributes         map[string]string `json:"attributes,omitempty"`
	SharedWithProjects []string          `json:"sharedWithProjects,omitempty"`
	ProjectID          string            `json:"projectId,omitempty"`
	ProjectName        string            `json:"projectName,omitempty"`
	DomainID           string            `json:"domainId,omitempty"`
	DomainName         string            `json:"domainName,omitempty"`
	CreatedAt          string            `json:"createdAt,omitempty"`
	UpdatedAt          string            `json:"updatedAt,omitempty"`
}

// Extract accepts a Page struct, specifically an v1.CommonPage struct,
// and extracts the elements into a slice of Zone structs.
func Extract(r pagination.Page) ([]Zone, error) {
	var s struct {
		Zones []Zone `json:"items"`
	}
	if err := (r.(v1.CommonPage)).ExtractInto(&s); err != nil {
		return nil, err
	}
	return s.Zones, nil
}

// GetResult represents the result of a get operation. Call its Extract method
// to interpret it as a Zone.
type GetResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a Zone
// resource.
func (r GetResult) Extract() (*Zone, error) {
	var s struct {
		Data struct {
			Item Zone `json:"item"`
		} `json:"data"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data.Item, nil
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}
