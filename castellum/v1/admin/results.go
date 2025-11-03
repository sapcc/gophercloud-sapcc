// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package admin

import (
	"github.com/gophercloud/gophercloud/v2"
)

// ResourceChecked holds the error from the last resource scrape attempt.
type ResourceChecked struct {
	Error string `json:"error"`
}

// FinishedEvent describes the outcome of a completed resize operation.
type FinishedEvent struct {
	At    int64  `json:"at"`
	Error string `json:"error,omitempty"`
}

// ResourceScrapeError represents a resource that failed its last scrape.
type ResourceScrapeError struct {
	AssetType string          `json:"asset_type"`
	Checked   ResourceChecked `json:"checked"`
	DomainID  string          `json:"domain_id"`
	ProjectID string          `json:"project_id,omitempty"`
}

// AssetScrapeError represents an asset that failed its last scrape.
type AssetScrapeError struct {
	AssetID   string          `json:"asset_id"`
	AssetType string          `json:"asset_type"`
	Checked   ResourceChecked `json:"checked"`
	DomainID  string          `json:"domain_id"`
	ProjectID string          `json:"project_id,omitempty"`
}

// AssetResizeError represents an asset whose last resize operation failed.
type AssetResizeError struct {
	AssetID   string        `json:"asset_id"`
	AssetType string        `json:"asset_type"`
	DomainID  string        `json:"domain_id"`
	ProjectID string        `json:"project_id,omitempty"`
	OldSize   int           `json:"old_size"`
	NewSize   int           `json:"new_size"`
	Finished  FinishedEvent `json:"finished"`
}

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
func (r ResourceScrapeErrorsResult) Extract() ([]ResourceScrapeError, error) {
	var s struct {
		Errors []ResourceScrapeError `json:"resource_scrape_errors"`
	}
	err := r.ExtractInto(&s)
	return s.Errors, err
}

// Extract returns a slice of AssetScrapeError entries.
func (r AssetScrapeErrorsResult) Extract() ([]AssetScrapeError, error) {
	var s struct {
		Errors []AssetScrapeError `json:"asset_scrape_errors"`
	}
	err := r.ExtractInto(&s)
	return s.Errors, err
}

// Extract returns a slice of AssetResizeError entries.
func (r AssetResizeErrorsResult) Extract() ([]AssetResizeError, error) {
	var s struct {
		Errors []AssetResizeError `json:"asset_resize_errors"`
	}
	err := r.ExtractInto(&s)
	return s.Errors, err
}
