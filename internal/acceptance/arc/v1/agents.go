// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package v1

import (
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

	response := agents.Init(t.Context(), client, createOpts)
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
