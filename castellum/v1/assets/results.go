// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"github.com/gophercloud/gophercloud/v2"
)

// AssetChecked holds the result of the last asset scrape attempt.
type AssetChecked struct {
	Error string `json:"error,omitempty"`
}

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

// Operation represents a resize operation on an asset.
type Operation struct {
	State     string          `json:"state"`
	Reason    string          `json:"reason"`
	OldSize   int             `json:"old_size"`
	NewSize   int             `json:"new_size"`
	Created   OperationEvent  `json:"created"`
	Confirmed *OperationEvent `json:"confirmed,omitempty"`
	Greenlit  *GreenlitEvent  `json:"greenlit,omitempty"`
	Finished  *FinishedEvent  `json:"finished,omitempty"`
}

// Asset represents a resizable object tracked by Castellum.
type Asset struct {
	ID                 string        `json:"id"`
	Size               int           `json:"size"`
	UsagePercent       float64       `json:"usage_percent"`
	MinSize            *int          `json:"min_size,omitempty"`
	MaxSize            *int          `json:"max_size,omitempty"`
	Checked            *AssetChecked `json:"checked,omitempty"`
	Stale              bool          `json:"stale"`
	PendingOperation   *Operation    `json:"pending_operation,omitempty"`
	FinishedOperations []Operation   `json:"finished_operations,omitempty"`
}

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
func (r ListResult) Extract() ([]Asset, error) {
	var s struct {
		Assets []Asset `json:"assets"`
	}
	err := r.ExtractInto(&s)
	return s.Assets, err
}

// Extract returns a single Asset from a GetResult.
func (r GetResult) Extract() (*Asset, error) {
	var s Asset
	err := r.ExtractIntoStructPtr(&s, "")
	return &s, err
}
