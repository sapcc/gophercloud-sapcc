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

package agents

import "github.com/gophercloud/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("agents")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func initURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("agents", "init")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func tagsURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id, "tags")
}

func deleteTagURL(c *gophercloud.ServiceClient, id, key string) string {
	return c.ServiceURL("agents", id, "tags", key)
}

func factsURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("agents", id, "facts")
}
