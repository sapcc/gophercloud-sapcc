package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/sapcc/gophercloud-sapcc/arc/v1/agents"
)

// CreateAgent will bootstrap an arc agent. An error will be returned if the
// arc bootstrap agent could not be created.
func InitAgent(t *testing.T, client *gophercloud.ServiceClient, accept string) (*string, error) {
	t.Logf("Attempting to bootstrap an arc agent: %s", accept)

	createOpts := agents.InitOpts{
		Accept: accept,
	}

	response := agents.Init(client, createOpts)
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
