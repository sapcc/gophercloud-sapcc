// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"strconv"

	"github.com/gophercloud/gophercloud/v2/pagination"
)

// GetCurrentAndTotalPages is a helper function that is used within the
// pagination.LastMarker() implementations of the pagination.MarkerPage interface.
func GetCurrentAndTotalPages(r pagination.MarkerPageBase) (currentPage, totalPages int, err error) {
	currentPage = 1
	if page := r.URL.Query().Get("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			return -1, -1, err
		}
		if p > 1 {
			currentPage = p
		}
	}

	totalPages = -1
	if pages := r.Header.Get("Pagination-Pages"); pages != "" {
		totalPages, err = strconv.Atoi(pages)
		if err != nil {
			return -1, -1, err
		}
	}

	return currentPage, totalPages, nil
}
