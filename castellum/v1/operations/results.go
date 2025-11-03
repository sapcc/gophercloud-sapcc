// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"github.com/gophercloud/gophercloud/v2"
)

// OperationEvent describes when a threshold crossing was observed.
type OperationEvent struct {
	At           int64   `json:"at"`
	UsagePercent float64 `json:"usage_percent,omitempty"`
}

// GreenlitEvent describes when an operation was approved to proceed.
type GreenlitEvent struct {
	At     int64  `json:"at,omitempty"`
	ByUser string `json:"by_user,omitempty"`
}

// FinishedEvent describes the outcome of a completed operation.
type FinishedEvent struct {
	At    int64  `json:"at"`
	Error string `json:"error,omitempty"`
}

// Operation represents a resize operation visible in the global operations endpoints.
type Operation struct {
	ProjectID string          `json:"project_id"`
	AssetType string          `json:"asset_type"`
	AssetID   string          `json:"asset_id"`
	State     string          `json:"state"`
	Reason    string          `json:"reason"`
	OldSize   int             `json:"old_size"`
	NewSize   int             `json:"new_size"`
	Created   OperationEvent  `json:"created"`
	Confirmed *OperationEvent `json:"confirmed,omitempty"`
	Greenlit  *GreenlitEvent  `json:"greenlit,omitempty"`
	Finished  *FinishedEvent  `json:"finished,omitempty"`
}

type commonResult struct {
	gophercloud.Result
}

// ListPendingResult is the result of a ListPending operation.
type ListPendingResult struct {
	commonResult
}

// ListRecentlyFailedResult is the result of a ListRecentlyFailed operation.
type ListRecentlyFailedResult struct {
	commonResult
}

// ListRecentlySucceededResult is the result of a ListRecentlySucceeded operation.
type ListRecentlySucceededResult struct {
	commonResult
}

// Extract returns a slice of pending Operations.
func (r ListPendingResult) Extract() ([]Operation, error) {
	var s struct {
		Operations []Operation `json:"pending_operations"`
	}
	err := r.ExtractInto(&s)
	return s.Operations, err
}

// Extract returns a slice of recently-failed Operations.
func (r ListRecentlyFailedResult) Extract() ([]Operation, error) {
	var s struct {
		Operations []Operation `json:"recently_failed_operations"`
	}
	err := r.ExtractInto(&s)
	return s.Operations, err
}

// Extract returns a slice of recently-succeeded Operations.
func (r ListRecentlySucceededResult) Extract() ([]Operation, error) {
	var s struct {
		Operations []Operation `json:"recently_succeeded_operations"`
	}
	err := r.ExtractInto(&s)
	return s.Operations, err
}
