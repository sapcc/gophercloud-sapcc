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

package clusters

import (
	"github.com/gophercloud/gophercloud"
	"github.com/sapcc/go-api-declarations/limes"
)

// CommonResult is the result of a Get/List operation. Call its appropriate
// Extract method to interpret it as a Cluster or a slice of Clusters.
type CommonResult struct {
	gophercloud.Result
}

// ExtractClusters interprets a CommonResult as a slice of Clusters.
func (r CommonResult) ExtractClusters() ([]limes.ClusterReport, error) {
	var s struct {
		Clusters []limes.ClusterReport `json:"clusters"`
	}

	err := r.ExtractInto(&s)
	return s.Clusters, err
}

// Extract interprets a CommonResult as a Cluster.
func (r CommonResult) Extract() (*limes.ClusterReport, error) {
	var s struct {
		Cluster *limes.ClusterReport `json:"cluster"`
	}
	err := r.ExtractInto(&s)
	return s.Cluster, err
}
