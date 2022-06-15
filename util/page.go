// Copyright 2022 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"strconv"

	"github.com/gophercloud/gophercloud/pagination"
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
