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

package v1

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/arc/v1/agents"
)

// CreateAgent will bootstrap an arc agent. An error will be returned if the
// arc bootstrap agent could not be created.
func InitAgent(t *testing.T, client *gophercloud.ServiceClient, accept string) (*string, error) {
	t.Logf("Attempting to bootstrap an arc agent: %s", accept)

	createOpts := agents.InitOpts{
		Accept: accept,
	}

	response := agents.Init(context.TODO(), client, createOpts)
	if response.Err != nil {
		return nil, response.Err
	}

	headers, err := response.ExtractHeaders()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, headers.ContentType, accept)

	t.Logf("Successfully bootstrapped an arc agent: %s", accept)

	c, err := response.ExtractContent()
	th.AssertNoErr(t, err)
	res := string(c)

	return &res, nil
}
