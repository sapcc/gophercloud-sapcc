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

const ListResponse = `
[
  {
    "year": 2019,
    "month": 8,
    "region": "region",
    "project_id": "1a894ddae4274a32a81eee43e4e5d67e",
    "object_id": "1a894ddae4274a32a81eee43e4e5d67e",
    "cost_object": "1234567",
    "cost_object_type": "IO",
    "co_inherited": false,
    "allocation_type": "usable",
    "service": "blockStorage",
    "measure": "capacity",
    "amount": 671930,
    "amount_unit": "GiBh",
    "duration": 671.93,
    "duration_unit": "h",
    "price_loc": 67.193,
    "price_sec": 0,
    "currency": "EUR"
  },
  {
    "year": 2019,
    "month": 8,
    "region": "region",
    "project_id": "1a894ddae4274a32a81eee43e4e5d67e",
    "object_id": "29940f04-961a-4903-a4c5-d91e750acc7f",
    "cost_object": "1234567",
    "cost_object_type": "IO",
    "co_inherited": false,
    "allocation_type": "provisioned",
    "service": "virtual",
    "measure": "os_suse",
    "amount": 1.00299,
    "amount_unit": "pieceh",
    "duration": 1.00299,
    "duration_unit": "h",
    "price_loc": 0.0001,
    "price_sec": 0,
    "currency": "EUR"
  }
]
`
