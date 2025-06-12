// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/metis/v1/identity/domains"
)

func TestGetDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetDomainSuccessfully(t)

	actual, err := domains.Get(t.Context(), fakeclient.ServiceClient(), "domain-1").Extract()
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

	p, err := domains.List(fakeclient.ServiceClient(), opts).AllPages(t.Context())
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
