// Copyright 2024 SAP SE
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

import "github.com/sapcc/gophercloud-sapcc/v2/networking/v2/bgpvpn/interconnections"

const ListInterconnectionsResponse = `
{
  "interconnections": [
    {
      "id": "a943ab0b-8b32-47dd-805b-4d17b7e15359",
      "project_id": "efc95e26964c46568de64dec41ee9204",
      "name": "interconnection1",
      "type": "bgpvpn",
      "state": "ACTIVE",
      "local_resource_id": "b1421282-a0e6-4add-9b6a-423aaabd67d2",
      "remote_resource_id": "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
      "remote_region": "region2",
      "remote_interconnection_id": "a341bc06-05f0-40a4-ad19-084fbf1b1c79",
      "local_parameters": {
        "project_id": [
          "efc95e26964c46568de64dec41ee9204"
        ]
      },
      "remote_parameters": {
        "project_id": [
          "c406568b6b144188af37c4bc7c5155dd"
        ]
      },
      "tenant_id": "efc95e26964c46568de64dec41ee9204"
    },
    {
      "id": "09b14c95-aaed-45b5-be5c-3012991e1b11",
      "project_id": "efc95e26964c46568de64dec41ee9204",
      "name": "interconnection2",
      "type": "bgpvpn",
      "state": "ACTIVE",
      "local_resource_id": "eb399c8f-7281-482b-bcd5-d6ee5ae91fac",
      "remote_resource_id": "fb08c60d-1a21-49bc-9c55-bd3ac6c30c09",
      "remote_region": "region2",
      "remote_interconnection_id": "c63cac82-24ed-440a-90da-b83ea3663e83",
      "local_parameters": {
        "project_id": [
          "efc95e26964c46568de64dec41ee9204"
        ]
      },
      "remote_parameters": {
        "project_id": [
          "c406568b6b144188af37c4bc7c5155dd"
        ]
      },
      "tenant_id": "efc95e26964c46568de64dec41ee9204"
    },
    {
      "id": "0be30220-a2d3-4fca-bdc6-b8aa861875a9",
      "project_id": "efc95e26964c46568de64dec41ee9204",
      "name": "",
      "type": "bgpvpn",
      "state": "ACTIVE",
      "local_resource_id": "11dcc528-58ac-4bba-b961-f3aa2b290510",
      "remote_resource_id": "aba7ca96-42d3-45c4-9287-48c7705f3c1b",
      "remote_region": "region2",
      "remote_interconnection_id": "66d46618-f451-4d8a-9503-3f48a7dc35e0",
      "local_parameters": {
        "project_id": [
          "efc95e26964c46568de64dec41ee9204"
        ]
      },
      "remote_parameters": {
        "project_id": [
          "c406568b6b144188af37c4bc7c5155dd"
        ]
      },
      "tenant_id": "efc95e26964c46568de64dec41ee9204"
    }
  ]
}
`

var interconnectionsList = []interconnections.Interconnection{
	{
		ID:                      "a943ab0b-8b32-47dd-805b-4d17b7e15359",
		ProjectID:               "efc95e26964c46568de64dec41ee9204",
		Name:                    "interconnection1",
		Type:                    "bgpvpn",
		State:                   "ACTIVE",
		LocalResourceID:         "b1421282-a0e6-4add-9b6a-423aaabd67d2",
		RemoteResourceID:        "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
		RemoteRegion:            "region2",
		RemoteInterconnectionID: "a341bc06-05f0-40a4-ad19-084fbf1b1c79",
		LocalParameters: interconnections.Parameters{
			ProjectID: []string{"efc95e26964c46568de64dec41ee9204"},
		},
		RemoteParameters: interconnections.Parameters{
			ProjectID: []string{"c406568b6b144188af37c4bc7c5155dd"},
		},
		TenantID: "efc95e26964c46568de64dec41ee9204",
	},
	{
		ID:                      "09b14c95-aaed-45b5-be5c-3012991e1b11",
		ProjectID:               "efc95e26964c46568de64dec41ee9204",
		Name:                    "interconnection2",
		Type:                    "bgpvpn",
		State:                   "ACTIVE",
		LocalResourceID:         "eb399c8f-7281-482b-bcd5-d6ee5ae91fac",
		RemoteResourceID:        "fb08c60d-1a21-49bc-9c55-bd3ac6c30c09",
		RemoteRegion:            "region2",
		RemoteInterconnectionID: "c63cac82-24ed-440a-90da-b83ea3663e83",
		LocalParameters: interconnections.Parameters{
			ProjectID: []string{"efc95e26964c46568de64dec41ee9204"},
		},
		RemoteParameters: interconnections.Parameters{
			ProjectID: []string{"c406568b6b144188af37c4bc7c5155dd"},
		},
		TenantID: "efc95e26964c46568de64dec41ee9204",
	},
	{
		ID:                      "0be30220-a2d3-4fca-bdc6-b8aa861875a9",
		ProjectID:               "efc95e26964c46568de64dec41ee9204",
		Name:                    "",
		Type:                    "bgpvpn",
		State:                   "ACTIVE",
		LocalResourceID:         "11dcc528-58ac-4bba-b961-f3aa2b290510",
		RemoteResourceID:        "aba7ca96-42d3-45c4-9287-48c7705f3c1b",
		RemoteRegion:            "region2",
		RemoteInterconnectionID: "66d46618-f451-4d8a-9503-3f48a7dc35e0",
		LocalParameters: interconnections.Parameters{
			ProjectID: []string{"efc95e26964c46568de64dec41ee9204"},
		},
		RemoteParameters: interconnections.Parameters{
			ProjectID: []string{"c406568b6b144188af37c4bc7c5155dd"},
		},
		TenantID: "efc95e26964c46568de64dec41ee9204",
	},
}

const GetInterconnectionResponse = `
{
    "interconnection": {
      "id": "a943ab0b-8b32-47dd-805b-4d17b7e15359",
      "project_id": "efc95e26964c46568de64dec41ee9204",
      "name": "interconnection1",
      "type": "bgpvpn",
      "state": "ACTIVE",
      "local_resource_id": "b1421282-a0e6-4add-9b6a-423aaabd67d2",
      "remote_resource_id": "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
      "remote_region": "region2",
      "remote_interconnection_id": "a341bc06-05f0-40a4-ad19-084fbf1b1c79",
      "local_parameters": {
        "project_id": [
          "efc95e26964c46568de64dec41ee9204"
        ]
      },
      "remote_parameters": {
        "project_id": [
          "c406568b6b144188af37c4bc7c5155dd"
        ]
      },
      "tenant_id": "efc95e26964c46568de64dec41ee9204"
    }
}
`

var interconnectionGet = interconnections.Interconnection{
	ID:                      "a943ab0b-8b32-47dd-805b-4d17b7e15359",
	ProjectID:               "efc95e26964c46568de64dec41ee9204",
	Name:                    "interconnection1",
	Type:                    "bgpvpn",
	State:                   "ACTIVE",
	LocalResourceID:         "b1421282-a0e6-4add-9b6a-423aaabd67d2",
	RemoteResourceID:        "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
	RemoteRegion:            "region2",
	RemoteInterconnectionID: "a341bc06-05f0-40a4-ad19-084fbf1b1c79",
	LocalParameters: interconnections.Parameters{
		ProjectID: []string{"efc95e26964c46568de64dec41ee9204"},
	},
	RemoteParameters: interconnections.Parameters{
		ProjectID: []string{"c406568b6b144188af37c4bc7c5155dd"},
	},
	TenantID: "efc95e26964c46568de64dec41ee9204",
}

const CreateInterconnectionRequest = `
{
  "interconnection": {
    "name": "interconnection1",
    "type": "bgpvpn",
    "local_resource_id": "b1421282-a0e6-4add-9b6a-423aaabd67d2",
    "remote_resource_id": "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
    "remote_region": "region2"
  }
}
`

var interconnectionCreate = interconnections.CreateOpts{
	Name:             "interconnection1",
	Type:             "bgpvpn",
	LocalResourceID:  "b1421282-a0e6-4add-9b6a-423aaabd67d2",
	RemoteResourceID: "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
	RemoteRegion:     "region2",
}

const CreateInterconnectionResponse = `
{
  "interconnection": {
    "id": "a943ab0b-8b32-47dd-805b-4d17b7e15359",
    "project_id": "efc95e26964c46568de64dec41ee9204",
    "name": "interconnection1",
    "type": "bgpvpn",
    "state": "ACTIVE",
    "local_resource_id": "b1421282-a0e6-4add-9b6a-423aaabd67d2",
    "remote_resource_id": "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
    "remote_region": "region2",
    "remote_interconnection_id": "a341bc06-05f0-40a4-ad19-084fbf1b1c79",
    "local_parameters": {
      "project_id": [
        "efc95e26964c46568de64dec41ee9204"
      ]
    },
    "remote_parameters": {
      "project_id": [
        "c406568b6b144188af37c4bc7c5155dd"
      ]
    },
    "tenant_id": "efc95e26964c46568de64dec41ee9204"
  }
}
`

const UpdateInterconnectionRequest = `
{
  "interconnection": {
    "name": "interconnection2",
    "state": "WAITING_REMOTE"
  }
}
`

const UpdateInterconnectionResponse = `
{
  "interconnection": {
    "id": "a943ab0b-8b32-47dd-805b-4d17b7e15359",
    "project_id": "efc95e26964c46568de64dec41ee9204",
    "name": "interconnection2",
    "type": "bgpvpn",
    "state": "WAITING_REMOTE",
    "local_resource_id": "b1421282-a0e6-4add-9b6a-423aaabd67d2",
    "remote_resource_id": "6a985ed2-cb97-45ae-a6dc-07feb0f36110",
    "remote_region": "region2",
    "remote_interconnection_id": "a341bc06-05f0-40a4-ad19-084fbf1b1c79",
    "local_parameters": {
      "project_id": [
        "efc95e26964c46568de64dec41ee9204"
      ]
    },
    "remote_parameters": {
      "project_id": [
        "c406568b6b144188af37c4bc7c5155dd"
      ]
    },
    "tenant_id": "efc95e26964c46568de64dec41ee9204"
  }
}
`
