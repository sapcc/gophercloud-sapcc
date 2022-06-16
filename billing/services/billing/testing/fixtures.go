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
    "REGION": "region",
    "PROJECT_ID": "1a894ddae4274a32a81eee43e4e5d67e",
    "PROJECT_NAME": "my-project",
    "OBJECT_ID": "1a894ddae4274a32a81eee43e4e5d67e",
    "METRIC_TYPE": "compute_ram_quota",
    "AMOUNT": "12.1688",
    "DURATION": "2.968",
    "PRICE_LOC": "4.1000000000",
    "PRICE_SEC": "0.0000000000",
    "COST_OBJECT": null,
    "COST_OBJECT_TYPE": null,
    "CO_INHERITED": 1,
    "SEND_CC": 123456789
  },
  {
    "REGION": "region",
    "PROJECT_ID": "1a894ddae4274a32a81eee43e4e5d67e",
    "PROJECT_NAME": "my-project",
    "OBJECT_ID": "1a894ddae4274a32a81eee43e4e5d67e",
    "METRIC_TYPE": "network_loadbalancers_quota",
    "AMOUNT": "3.2648",
    "DURATION": "2.968",
    "PRICE_LOC": "1.1000000000",
    "PRICE_SEC": "0.0000000000",
    "COST_OBJECT": "123",
    "COST_OBJECT_TYPE": "CC",
    "CO_INHERITED": 0,
    "SEND_CC": 123456789
  }
]
`
