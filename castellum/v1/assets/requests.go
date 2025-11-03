// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
)

// List returns all assets for a given project resource type.
func List(ctx context.Context, c *gophercloud.ServiceClient, projectID, assetType string) (r ListResult) {
	resp, err := c.Get(ctx, assetsURL(c, projectID, assetType), &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a single asset. Set history to true to include finished operations.
func Get(ctx context.Context, c *gophercloud.ServiceClient, projectID, assetType, assetID string, history bool) (r GetResult) {
	url := assetURL(c, projectID, assetType, assetID)
	if history {
		url += "?history"
	}
	resp, err := c.Get(ctx, url, &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ResolveError marks the latest errored operation on an asset as resolved.
func ResolveError(ctx context.Context, c *gophercloud.ServiceClient, projectID, assetType, assetID string) (r ResolveErrorResult) {
	resp, err := c.Post(ctx, errorResolvedURL(c, projectID, assetType, assetID), nil, nil, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
