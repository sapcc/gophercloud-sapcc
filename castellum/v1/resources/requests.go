// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
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
	ToResourceCreateMap() (map[string]any, error)
}

// CreateOpts specifies the autoscaling configuration for a resource.
type CreateOpts struct {
	LowThreshold      *Threshold       `json:"low_threshold,omitempty"`
	HighThreshold     *Threshold       `json:"high_threshold,omitempty"`
	CriticalThreshold *Threshold       `json:"critical_threshold,omitempty"`
	SizeConstraints   *SizeConstraints `json:"size_constraints,omitempty"`
	SizeSteps         *SizeSteps       `json:"size_steps,omitempty"`
}

// ToResourceCreateMap assembles a request body based on the contents of CreateOpts.
func (opts CreateOpts) ToResourceCreateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "")
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
	b, err := opts.ToResourceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(ctx, createURL(c, projectID, resourceType), b, nil, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusAccepted, http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
