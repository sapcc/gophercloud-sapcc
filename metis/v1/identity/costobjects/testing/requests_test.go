// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/metis/v1/identity/costobjects"
)

func TestGetProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetCostObjectSuccessfully(t)

	actual, err := costobjects.Get(t.Context(), fakeclient.ServiceClient(), "costobject-1").Extract()
	th.AssertNoErr(t, err)

	expected := &costobjects.CostObject{
		Name: "costobject-1",
		Type: "CC",
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListCostObjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListCostObjectsSuccessfully(t)
	opts := costobjects.ListOpts{
		Limit: 1,
		// Set the domain and project options to verify the filtering works
		Domain:  "foo",
		Project: "bar",
	}

	p, err := costobjects.List(fakeclient.ServiceClient(), opts).AllPages(t.Context())
	th.AssertNoErr(t, err)
	actual, err := costobjects.Extract(p)
	th.AssertNoErr(t, err)

	expected := []costobjects.CostObject{
		{
			Name: "costobject-1",
			Type: "CC",
		},
		{
			Name: "costobject-2",
			Type: "IO",
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}
