// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/sapcc/go-api-declarations/castellum"
)

// Resource wraps castellum.Resource with the resource type name from the URL.
type Resource struct {
	castellum.Resource
	ResourceType string `json:"-"`
}

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

func (r ListResult) Extract() ([]Resource, error) {
	var s struct {
		Resources map[string]*castellum.Resource `json:"resources"`
	}

	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}

	extracted := make([]Resource, 0, len(s.Resources))
	for name, resource := range s.Resources {
		extracted = append(extracted, Resource{Resource: *resource, ResourceType: name})
	}

	return extracted, nil
}

func (r GetResult) Extract() (*Resource, error) {
	var s castellum.Resource
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return &Resource{Resource: s, ResourceType: r.resourceType}, err
}

func (r GetResult) ExtractInto(v any) error {
	return r.ExtractIntoStructPtr(v, "")
}
