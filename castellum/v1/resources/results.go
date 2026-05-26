// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/sapcc/go-api-declarations/castellum"
)

type GetResult struct {
	gophercloud.Result
	resourceType string
}

type ListResult struct {
	gophercloud.Result
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

type CreateResult struct {
	gophercloud.ErrResult
}

// Extract returns the full set of resources, keyed on asset type.
func (r ListResult) Extract() (map[string]castellum.Resource, error) {
	var s struct {
		Resources map[string]castellum.Resource `json:"resources"`
	}
	err := r.ExtractInto(&s)
	return s.Resources, err
}

func (r GetResult) Extract() (castellum.Resource, error) {
	var s castellum.Resource
	err := r.ExtractInto(&s)
	return s, err
}

func (r GetResult) ExtractInto(v any) error {
	return r.ExtractIntoStructPtr(v, "")
}
