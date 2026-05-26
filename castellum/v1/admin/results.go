// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package admin

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/sapcc/go-api-declarations/castellum"
)

type commonResult struct {
	gophercloud.Result
}

// ResourceScrapeErrorsResult is the result of a GetResourceScrapeErrors operation.
type ResourceScrapeErrorsResult struct {
	commonResult
}

// AssetScrapeErrorsResult is the result of a GetAssetScrapeErrors operation.
type AssetScrapeErrorsResult struct {
	commonResult
}

// AssetResizeErrorsResult is the result of a GetAssetResizeErrors operation.
type AssetResizeErrorsResult struct {
	commonResult
}

// Extract returns a slice of ResourceScrapeError entries.
func (r ResourceScrapeErrorsResult) Extract() ([]castellum.ResourceScrapeError, error) {
	var s struct {
		Errors []castellum.ResourceScrapeError `json:"resource_scrape_errors"`
	}
	err := r.ExtractInto(&s)
	return s.Errors, err
}

// Extract returns a slice of AssetScrapeError entries.
func (r AssetScrapeErrorsResult) Extract() ([]castellum.AssetScrapeError, error) {
	var s struct {
		Errors []castellum.AssetScrapeError `json:"asset_scrape_errors"`
	}
	err := r.ExtractInto(&s)
	return s.Errors, err
}

// Extract returns a slice of AssetResizeError entries.
func (r AssetResizeErrorsResult) Extract() ([]castellum.AssetResizeError, error) {
	var s struct {
		Errors []castellum.AssetResizeError `json:"asset_resize_errors"`
	}
	err := r.ExtractInto(&s)
	return s.Errors, err
}
