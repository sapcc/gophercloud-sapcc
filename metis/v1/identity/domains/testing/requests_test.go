// Copyright 2023 SAP SE
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

package testing

import (
	"context"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/metis/v1/identity/domains"
)

func TestGetDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetDomainSuccessfully(t)

	actual, err := domains.Get(context.TODO(), fakeclient.ServiceClient(), "domain-1").Extract()
	th.AssertNoErr(t, err)

	expected := &domains.Domain{
		Name:        "domain1",
		ID:          "test-domain",
		Description: "domain1 descr",
		BillingMetadata: domains.BillingDomainMetadata{
			PrimaryContactUserID:  "i001337",
			PrimaryContactEmail:   "max.mustermann@sample.com",
			AdditionalInformation: "",
			CostObjectName:        "my-cost-center",
			CostObjectType:        "CC",
			ProjectsCanInherit:    true,
		},
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListDomains(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDomainsSuccessfully(t)
	opts := domains.ListOpts{
		Limit: 1,
	}

	p, err := domains.List(fakeclient.ServiceClient(), opts).AllPages(context.TODO())
	th.AssertNoErr(t, err)
	actual, err := domains.Extract(p)
	th.AssertNoErr(t, err)

	expected := []domains.Domain{
		{
			Name:        "domain1",
			ID:          "domain-1",
			Description: "domain1 descr",
			BillingMetadata: domains.BillingDomainMetadata{
				PrimaryContactUserID:  "i001337",
				PrimaryContactEmail:   "max.mustermann@sample.com",
				AdditionalInformation: "",
				CostObjectName:        "my-cost-center",
				CostObjectType:        "CC",
				ProjectsCanInherit:    true,
			},
		},
		{
			Name:        "domain2",
			ID:          "domain-2",
			Description: "domain2 descr",
			BillingMetadata: domains.BillingDomainMetadata{
				PrimaryContactUserID:  "i000042",
				PrimaryContactEmail:   "dies.das@sample.com",
				AdditionalInformation: "",
				CostObjectName:        "my-internal-order",
				CostObjectType:        "IO",
				ProjectsCanInherit:    false,
			},
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}
