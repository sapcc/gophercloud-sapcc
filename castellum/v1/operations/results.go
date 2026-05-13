// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/sapcc/go-api-declarations/castellum"
)

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
func (r ListPendingResult) Extract() ([]castellum.StandaloneOperation, error) {
	var s struct {
		Operations []castellum.StandaloneOperation `json:"pending_operations"`
	}
	err := r.ExtractInto(&s)
	return s.Operations, err
}

// Extract returns a slice of recently-failed Operations.
func (r ListRecentlyFailedResult) Extract() ([]castellum.StandaloneOperation, error) {
	var s struct {
		Operations []castellum.StandaloneOperation `json:"recently_failed_operations"`
	}
	err := r.ExtractInto(&s)
	return s.Operations, err
}

// Extract returns a slice of recently-succeeded Operations.
func (r ListRecentlySucceededResult) Extract() ([]castellum.StandaloneOperation, error) {
	var s struct {
		Operations []castellum.StandaloneOperation `json:"recently_succeeded_operations"`
	}
	err := r.ExtractInto(&s)
	return s.Operations, err
}
