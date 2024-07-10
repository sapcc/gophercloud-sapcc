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

	"github.com/sapcc/gophercloud-sapcc/v2/metis/v1/identity/costobjects"
)

func TestGetProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetCostObjectSuccessfully(t)

	actual, err := costobjects.Get(context.TODO(), fakeclient.ServiceClient(), "costobject-1").Extract()
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

	p, err := costobjects.List(fakeclient.ServiceClient(), opts).AllPages(context.TODO())
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
