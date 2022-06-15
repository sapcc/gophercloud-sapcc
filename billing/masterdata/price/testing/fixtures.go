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
    "SEND_CC": 123456789,
    "COST_ELEMENT": 123456,
    "PRICE_LOC": "0.123456",
    "PRICE_SEC": "0.000000",
    "VALID_FROM": "2019-05-01T00:00:00",
    "VALID_TO": "9999-12-31T00:00:00",
    "METRIC_TYPE": "foo",
    "REGION": "region",
    "VALID_FOR_PROJECT_TYPE": "quotaUsage",
    "OBJECT_TYPE": "object"
  },
  {
    "SEND_CC": 123456789,
    "COST_ELEMENT": 123457,
    "PRICE_LOC": "0.023456",
    "PRICE_SEC": "0.000000",
    "VALID_FROM": "2019-05-01T00:00:00",
    "VALID_TO": "9999-12-31T00:00:00",
    "METRIC_TYPE": "bar",
    "REGION": "region",
    "VALID_FOR_PROJECT_TYPE": "quotaUsage",
    "OBJECT_TYPE": "object"
  }
]
`
