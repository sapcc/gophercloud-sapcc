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

	"github.com/sapcc/gophercloud-sapcc/v2/metis/v1/network/ip"
)

func TestGetProject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetIPAddressSuccessfully(t)

	actual, err := ip.Get(context.TODO(), fakeclient.ServiceClient(), "10.216.24.194").Extract()
	th.AssertNoErr(t, err)

	expected := &ip.IPAddress{
		IP:          "10.216.24.194",
		PortUUID:    "9cf53dfa-8a72-1337-bf69-523d11ffccb9",
		Description: "",
		Status:      "ACTIVE",
		DeviceID:    "dhcp5d784bae-4201-530e-90df-393914f8601b-1563904c-ac3d-4281-994a-676d9a1716c6",
		DeviceOwner: "network:dhcp",
		NetworkID:   "1563904c-ac3d-1337-994a-676d9a1716c6",
		NetworkName: "network-name",
		SubnetID:    "e6f6ff0c-42fa-1337-9e78-2a8405fed887",
		SubnetName:  "subnet-name",
		DomainID:    "666da95112694b37b3efb0913de31337",
		DomainName:  "admin",
		ProjectID:   "0420083ad7d145dc9fdb9ccdb59ad5b6",
		ProjectName: "admin-net-infra",
		Created:     "2019-08-12 09:22:54",
		LastChanged: "2023-08-10 14:26:51",
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestListProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListIPAddressesSuccessfully(t)

	opts := ip.ListOpts{Limit: 1}

	p, err := ip.List(fakeclient.ServiceClient(), opts).AllPages(context.TODO())
	th.AssertNoErr(t, err)
	actual, err := ip.Extract(p)
	th.AssertNoErr(t, err)

	expected := []ip.IPAddress{
		{
			IP:          "127.0.0.1",
			PortUUID:    "9cf53dfa-8a72-1337-bf69-523d11ffccb9",
			Description: "test",
			Status:      "ACTIVE",
			DeviceID:    "dhcp5d784bae-4201-530e-90df-393914f8601b-1563904c-ac3d-4281-994a-676d9a1716c6",
			DeviceOwner: "network:dhcp",
			NetworkID:   "1563904c-ac3d-1337-994a-676d9a1716c6",
			NetworkName: "network-name",
			SubnetID:    "e6f6ff0c-42fa-1337-9e78-2a8405fed887",
			SubnetName:  "subnet-name",
			DomainID:    "666da95112694b37b3efb0913de31337",
			DomainName:  "admin",
			ProjectID:   "0420083ad7d145dc9fdb9ccdb59ad5b6",
			ProjectName: "admin-net-infra",
			Created:     "2019-08-12 09:22:54",
			LastChanged: "2023-08-10 14:26:51",
		}, {
			IP:          "192.0.0.1",
			PortUUID:    "9cf53dfa-8a72-1337-bf69-523d11ffccb9",
			Description: "dummy",
			Status:      "ACTIVE",
			DeviceID:    "dhcp5d784bae-4201-530e-90df-393914f8601b-1563904c-ac3d-4281-994a-676d9a1716c6",
			DeviceOwner: "network:dhcp",
			NetworkID:   "1563904c-ac3d-1337-994a-676d9a1716c6",
			NetworkName: "network-name",
			SubnetID:    "e6f6ff0c-42fa-1337-9e78-2a8405fed887",
			SubnetName:  "subnet-name",
			DomainID:    "666da95112694b37b3efb0913de31337",
			DomainName:  "admin",
			ProjectID:   "0420083ad7d145dc9fdb9ccdb59ad5b6",
			ProjectName: "admin-net-infra",
			Created:     "2019-08-12 09:22:54",
			LastChanged: "2023-08-10 14:26:51",
		}}

	th.CheckDeepEquals(t, expected, actual)
}
