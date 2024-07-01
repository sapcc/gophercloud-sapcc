// Copyright 2024 SAP SE
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

package interconnections

import "github.com/gophercloud/gophercloud/v2"

const urlBase = "interconnection/interconnections"

// return /v2.0/interconnection/interconnections/{interconnection-id}
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(urlBase, id)
}

// return /v2.0/interconnection/interconnections
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(urlBase)
}

// return /v2.0/interconnection/interconnections/{interconnection-id}
func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/interconnection/interconnections
func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/interconnection/interconnections
func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/interconnection/interconnections/{interconnection-id}
func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/interconnection/interconnections/{interconnection-id}
func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
