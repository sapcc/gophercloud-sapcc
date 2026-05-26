// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/sapcc/go-api-declarations/castellum"
	. "go.xyrillian.de/gg/option"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToResourceListQuery() (string, error)
}

// ListOpts filters the resources returned by the List function.
type ListOpts struct {
	ResourceType string `q:"resource_type"`
}

// ToResourceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourceListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToResourceCreateBody() ([]byte, error)
}

// CreateOpts specifies the autoscaling configuration for a resource.
type CreateOpts struct {
	LowThreshold      Option[castellum.Threshold]       `json:"low_threshold,omitzero"`
	HighThreshold     Option[castellum.Threshold]       `json:"high_threshold,omitzero"`
	CriticalThreshold Option[castellum.Threshold]       `json:"critical_threshold,omitzero"`
	SizeConstraints   Option[castellum.SizeConstraints] `json:"size_constraints,omitzero"`
	SizeSteps         Option[castellum.SizeSteps]       `json:"size_steps,omitzero"`
}

// ToResourceCreateBody marshals the CreateOpts into a JSON request body.
func (opts CreateOpts) ToResourceCreateBody() ([]byte, error) {
	return json.Marshal(opts)
}

// List returns all resources configured for a project.
func List(ctx context.Context, c *gophercloud.ServiceClient, projectID string, opts ListOptsBuilder) (r ListResult) {
	url := listURL(c, projectID)
	if opts != nil {
		query, err := opts.ToResourceListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves the autoscaling configuration for a specific resource type.
func Get(ctx context.Context, c *gophercloud.ServiceClient, projectID, resourceType string) (r GetResult) {
	resp, err := c.Get(ctx, getURL(c, projectID, resourceType), &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	r.resourceType = resourceType
	return
}

// Delete disables autoscaling for a resource type and removes all operation logs.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, projectID, resourceType string) (r DeleteResult) {
	resp, err := c.Delete(ctx, deleteURL(c, projectID, resourceType), &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Create enables autoscaling for a resource type with the given configuration.
func Create(ctx context.Context, c *gophercloud.ServiceClient, projectID, resourceType string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToResourceCreateBody()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(ctx, createURL(c, projectID, resourceType), json.RawMessage(b), nil, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusAccepted, http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
