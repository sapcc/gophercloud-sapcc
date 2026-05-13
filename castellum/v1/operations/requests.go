// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
)

// ListOptsBuilder allows extensions to add additional parameters to list requests.
type ListOptsBuilder interface {
	ToOperationListQuery() (string, error)
}

// ListOpts filters the operations returned by list functions.
type ListOpts struct {
	Project   string `q:"project"`
	Domain    string `q:"domain"`
	AssetType string `q:"asset-type"`
	// MaxAge filters recently-succeeded operations by age (e.g. "1d", "2h", "30m").
	// Only applicable to ListRecentlySucceeded and its project-scoped variant.
	MaxAge string `q:"max-age"`
}

// ToOperationListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToOperationListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func buildURL(url string, opts ListOptsBuilder) (string, error) {
	if opts == nil {
		return url, nil
	}
	query, err := opts.ToOperationListQuery()
	if err != nil {
		return "", err
	}
	return url + query, nil
}

// ListPending returns all pending operations accessible to the authenticated user.
func ListPending(ctx context.Context, c *gophercloud.ServiceClient, opts ListOptsBuilder) (r ListPendingResult) {
	url, err := buildURL(pendingURL(c), opts)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListRecentlyFailed returns operations that recently failed and still require intervention.
func ListRecentlyFailed(ctx context.Context, c *gophercloud.ServiceClient, opts ListOptsBuilder) (r ListRecentlyFailedResult) {
	url, err := buildURL(recentlyFailedURL(c), opts)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListRecentlySucceeded returns operations that recently succeeded.
func ListRecentlySucceeded(ctx context.Context, c *gophercloud.ServiceClient, opts ListOptsBuilder) (r ListRecentlySucceededResult) {
	url, err := buildURL(recentlySucceededURL(c), opts)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListProjectPending returns pending operations for a specific project resource type.
func ListProjectPending(ctx context.Context, c *gophercloud.ServiceClient, projectID, assetType string, opts ListOptsBuilder) (r ListPendingResult) {
	url, err := buildURL(projectPendingURL(c, projectID, assetType), opts)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListProjectRecentlyFailed returns recently-failed operations for a specific project resource type.
func ListProjectRecentlyFailed(ctx context.Context, c *gophercloud.ServiceClient, projectID, assetType string, opts ListOptsBuilder) (r ListRecentlyFailedResult) {
	url, err := buildURL(projectRecentlyFailedURL(c, projectID, assetType), opts)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListProjectRecentlySucceeded returns recently-succeeded operations for a specific project resource type.
func ListProjectRecentlySucceeded(ctx context.Context, c *gophercloud.ServiceClient, projectID, assetType string, opts ListOptsBuilder) (r ListRecentlySucceededResult) {
	url, err := buildURL(projectRecentlySucceededURL(c, projectID, assetType), opts)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
