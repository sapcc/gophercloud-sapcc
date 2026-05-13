// SPDX-FileCopyrightText: 2026 Dexter Le <dextersydney2001@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package admin

import (
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
)

// GetResourceScrapeErrors returns all resource scrape errors across all resources.
func GetResourceScrapeErrors(ctx context.Context, c *gophercloud.ServiceClient) (r ResourceScrapeErrorsResult) {
	resp, err := c.Get(ctx, resourceScrapeErrorsURL(c), &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetAssetScrapeErrors returns all asset scrape errors across all resources.
func GetAssetScrapeErrors(ctx context.Context, c *gophercloud.ServiceClient) (r AssetScrapeErrorsResult) {
	resp, err := c.Get(ctx, assetScrapeErrorsURL(c), &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetAssetResizeErrors returns all asset resize errors across all resources.
func GetAssetResizeErrors(ctx context.Context, c *gophercloud.ServiceClient) (r AssetResizeErrorsResult) {
	resp, err := c.Get(ctx, assetResizeErrorsURL(c), &r.Body, &gophercloud.RequestOpts{ //nolint:bodyclose // already handled by gophercloud
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
