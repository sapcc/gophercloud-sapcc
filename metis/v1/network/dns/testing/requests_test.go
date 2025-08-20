// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/metis/v1/network/dns"
)

func TestGetDNSZone(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleGetDNSZoneSuccessfully(t, fakeServer)

	actual, err := dns.Get(t.Context(), client.ServiceClient(fakeServer), "004321e142604754a789dd9b23db3242").Extract()
	th.AssertNoErr(t, err)

	expected := &dns.Zone{
		UUID:            "004321e142604754a789dd9b23db3242",
		Name:            "test-regression.germany.cloud.de.",
		Email:           "test.user@cloud.de",
		Serial:          1680176469,
		ParentZoneID:    "831905c1640146358f69d58437a2a042",
		ParentZoneName:  "germany.cloud.de.",
		Pool:            "default",
		PoolDescription: "Bind9 Pool",
		TTL:             7200,
		Status:          "ACTIVE",
		Action:          "NONE",
		Type:            "PRIMARY",
		Attributes: map[string]string{
			"pool_id": "794ccc2c-d75e-1337-b57f-8894c9f5c842",
		},
		SharedWithProjects: []string{
			"97a92f9cea1337c8a3c3bbe01caa842e",
		},
		ProjectID:   "97a92f9cea4644c8a3c3bbe01caa842e",
		ProjectName: "regression",
		DomainID:    "8395b48e21337b8a827cc76b5fcf1c8",
		DomainName:  "test",
		CreatedAt:   "2022-04-04 07:41:42",
		UpdatedAt:   "2023-03-30 11:41:41",
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestListDNSZones(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()
	HandleListDNSZonesSuccessfully(t, fakeServer)

	opts := dns.ListOpts{Limit: 1}

	p, err := dns.List(client.ServiceClient(fakeServer), opts).AllPages(t.Context())
	th.AssertNoErr(t, err)
	actual, err := dns.Extract(p)
	th.AssertNoErr(t, err)

	expected := []dns.Zone{
		{
			UUID:            "004321e142604754a789dd9b23db3242",
			Name:            "test-regression.germany.cloud.de.",
			Email:           "test.user@cloud.de",
			Serial:          1680176469,
			ParentZoneID:    "831905c1640146358f69d58437a2a042",
			ParentZoneName:  "germany.cloud.de.",
			Pool:            "default",
			PoolDescription: "Bind9 Pool",
			TTL:             7200,
			Status:          "ACTIVE",
			Action:          "NONE",
			Type:            "PRIMARY",
			Attributes: map[string]string{
				"pool_id": "794ccc2c-d75e-1337-b57f-8894c9f5c842",
			},
			ProjectID:   "97a92f9cea4644c8a3c3bbe01caa842e",
			ProjectName: "regression",
			DomainID:    "8395b48e21337b8a827cc76b5fcf1c8",
			DomainName:  "test",
			CreatedAt:   "2022-04-04 07:41:42",
			UpdatedAt:   "2023-03-30 11:41:41",
		},
		{
			UUID:            "17374100fd4b4e72b94353fc1931a920",
			Name:            "hermestest.test.com.",
			Description:     "hermes test",
			Email:           "user.test@cloud.de",
			Serial:          1688691690,
			Pool:            "default",
			PoolDescription: "Bind9 Pool",
			TTL:             7200,
			Status:          "ACTIVE",
			Action:          "NONE",
			Type:            "PRIMARY",
			Attributes: map[string]string{
				"pool_id": "794ccc2c-d751-44fe-b57f-1337c9f5c842",
			},
			SharedWithProjects: []string{
				"97a92f9cea1337c8a3c3bbe01caa842e",
			},
			ProjectID:   "e9141fb24eee4b3e9f25ae69cda31142",
			ProjectName: "demo",
			DomainID:    "2bac466eed364d8a92e477459e901337",
			DomainName:  "admin",
			CreatedAt:   "2019-10-14 14:15:18",
			UpdatedAt:   "2023-07-07 01:01:23",
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}
