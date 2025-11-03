// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"github.com/gophercloud/gophercloud/v2"
)

// Threshold represents a usage threshold that triggers a resize operation.
type Threshold struct {
	UsagePercent float64 `json:"usage_percent"`
	DelaySeconds int     `json:"delay_seconds,omitempty"`
}

// SizeConstraints defines the boundaries for resize operations.
type SizeConstraints struct {
	Minimum               *int `json:"minimum,omitempty"`
	Maximum               *int `json:"maximum,omitempty"`
	MinimumFree           *int `json:"minimum_free,omitempty"`
	MinimumFreeIsCritical bool `json:"minimum_free_is_critical,omitempty"`
}

// SizeSteps defines how resize increments are calculated.
// Exactly one of Percent or Single should be set.
type SizeSteps struct {
	Percent float64 `json:"percent,omitempty"`
	Single  bool    `json:"single,omitempty"`
}

// ResourceChecked holds the result of the last resource scrape attempt.
type ResourceChecked struct {
	Error string `json:"error,omitempty"`
}

// Resource represents a Castellum resource with autoscaling configuration.
type Resource struct {
	ResourceType      string           `json:"-"`
	AssetCount        int              `json:"asset_count"`
	Checked           *ResourceChecked `json:"checked,omitempty"`
	LowThreshold      *Threshold       `json:"low_threshold,omitempty"`
	HighThreshold     *Threshold       `json:"high_threshold,omitempty"`
	CriticalThreshold *Threshold       `json:"critical_threshold,omitempty"`
	SizeConstraints   *SizeConstraints `json:"size_constraints,omitempty"`
	SizeSteps         *SizeSteps       `json:"size_steps,omitempty"`
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
		Resources map[string]*Resource `json:"resources"`
	}

	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}

	extracted := make([]Resource, 0, len(s.Resources))
	for name, resource := range s.Resources {
		resource.ResourceType = name
		extracted = append(extracted, *resource)
	}

	return extracted, nil
}

func (r GetResult) Extract() (*Resource, error) {
	var s Resource
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	s.ResourceType = r.resourceType
	return &s, err
}

func (r GetResult) ExtractInto(v any) error {
	return r.ExtractIntoStructPtr(v, "")
}
