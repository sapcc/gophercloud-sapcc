// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const ListResponse = `
{
  "assets": [
    {
      "id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
      "size": 100,
      "usage_percent": 75.5,
      "min_size": 10,
      "max_size": 1000,
      "checked": {"error": ""},
      "stale": false
    },
    {
      "id": "5d7f5c1c-3f2e-4b0a-9e6d-8a1b2c3d4e5f",
      "size": 200,
      "usage_percent": 20.0,
      "stale": false
    }
  ]
}
`

const GetResponse = `
{
  "id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
  "size": 100,
  "usage_percent": 75.5,
  "min_size": 10,
  "max_size": 1000,
  "checked": {"error": ""},
  "stale": false,
  "pending_operation": {
    "state": "confirmed",
    "reason": "high",
    "old_size": 100,
    "new_size": 120,
    "created": {"at": 1700000000, "usage_percent": 82.0},
    "confirmed": {"at": 1700003600},
    "greenlit": {"at": 1700007200}
  }
}
`

const GetHistoryResponse = `
{
  "id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
  "size": 100,
  "usage_percent": 75.5,
  "stale": false,
  "finished_operations": [
    {
      "state": "succeeded",
      "reason": "high",
      "old_size": 80,
      "new_size": 100,
      "created": {"at": 1699000000, "usage_percent": 85.0},
      "confirmed": {"at": 1699003600},
      "greenlit": {"at": 1699007200},
      "finished": {"at": 1699010800}
    }
  ]
}
`
