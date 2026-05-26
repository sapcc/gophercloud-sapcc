// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/sapcc/go-api-declarations/castellum"
)

type commonResult struct {
	gophercloud.Result
}

// ListResult is the result of a List operation.
type ListResult struct {
	commonResult
}

// GetResult is the result of a Get operation.
type GetResult struct {
	commonResult
}

// ResolveErrorResult is the result of a ResolveError operation.
type ResolveErrorResult struct {
	gophercloud.ErrResult
}

// Extract returns a slice of Assets from a ListResult.
func (r ListResult) Extract() ([]castellum.Asset, error) {
	var s struct {
		Assets []castellum.Asset `json:"assets"`
	}
	err := r.ExtractInto(&s)
	return s.Assets, err
}

// Extract returns a single Asset from a GetResult.
func (r GetResult) Extract() (castellum.Asset, error) {
	var s castellum.Asset
	err := r.ExtractIntoStructPtr(&s, "")
	return s, err
}
