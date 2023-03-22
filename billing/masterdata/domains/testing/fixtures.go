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

package testing

const GetResponse = `
{
  "domain_id": "707c94677ac741ecb1f2cabc804c1285",
  "iid": 123,
  "domain_name": "master",
  "description": "example domain",
  "cost_object": {
    "name": "1234567",
    "type": "IO",
    "projects_can_inherit": false
  },
  "responsible_primary_contact_id": null,
  "responsible_primary_contact_email": null,
  "additional_information": null,
  "changed_by": "c48b0ce218848fd0bc78c8367ae9c40512024e2fc39451f47d9a62ad3ff41c26",
  "changed_at": "2019-01-29T09:37:58.792",
  "collector": "billing.region.local",
  "region": "region",
  "is_complete": false,
  "missing_attributes": "Primary contact not specified"
}
`

const ListResponse = `
[
  {
    "domain_id": "707c94677ac741ecb1f2cabc804c1285",
    "iid": 123,
    "domain_name": "master",
    "description": "example domain",
    "cost_object": {
      "name": "1234567",
      "type": "IO",
      "projects_can_inherit": false
    },
    "responsible_primary_contact_id": null,
    "responsible_primary_contact_email": null,
    "additional_information": null,
    "changed_by": "c48b0ce218848fd0bc78c8367ae9c40512024e2fc39451f47d9a62ad3ff41c26",
    "changed_at": "2019-01-29T09:37:58.792",
    "collector": "billing.region.local",
    "region": "region",
    "is_complete": false,
    "missing_attributes": "Primary contact not specified"
  }
]
`

const UpdateRequest = `
{
  "description" : "new example domain",
  "additional_information": "",
  "responsible_primary_contact_email" : "example@mail.com",
  "cost_object" : {
     "projects_can_inherit" : true
  },
  "domain_id" : "707c94677ac741ecb1f2cabc804c1285",
  "domain_name" : "master",
  "responsible_primary_contact_id" : "D123456",
  "collector": "billing.region.local",
  "region": "region"
}
`

const UpdateResponse = `
{
  "domain_id": "707c94677ac741ecb1f2cabc804c1285",
  "iid": 123,
  "domain_name": "master",
  "description": "new example domain",
  "cost_object": {
    "projects_can_inherit": true
  },
  "responsible_primary_contact_id": "D123456",
  "responsible_primary_contact_email": null,
  "additional_information": null,
  "changed_by": "c48b0ce218848fd0bc78c8367ae9c40512024e2fc39451f47d9a62ad3ff41c26",
  "changed_at": "2019-01-29T09:37:58.792",
  "collector": "billing.region.local",
  "region": "region",
  "is_complete": true,
  "missing_attributes": null
}
`
