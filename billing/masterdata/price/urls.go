// Copyright 2020 SAP SE
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

package price

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

func listURL(c *gophercloud.ServiceClient, opts ListOpts) string {
	if opts.Region == "" {
		return c.ServiceURL("masterdata", "pricelist")
	}

	if opts.To == (time.Time{}) {
		opts.To = time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)
	}

	return c.ServiceURL(
		"masterdata",
		"price",
		opts.Region,
		opts.MetricType,
		opts.From.Format(gophercloud.RFC3339NoZ),
		opts.To.Format(gophercloud.RFC3339NoZ),
	)
}
