// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const ListPendingResponse = `
{
  "pending_operations": [
    {
      "project_id": "88e5cad3-38e6-454f-b412-662cda03e7a1",
      "asset_type": "nfs-shares",
      "asset_id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
      "state": "confirmed",
      "reason": "high",
      "old_size": 100,
      "new_size": 120,
      "created": {"at": 1700000000, "usage_percent": 82.0},
      "confirmed": {"at": 1700003600},
      "greenlit": {"at": 1700007200}
    }
  ]
}
`

const ListRecentlyFailedResponse = `
{
  "recently_failed_operations": [
    {
      "project_id": "88e5cad3-38e6-454f-b412-662cda03e7a1",
      "asset_type": "nfs-shares",
      "asset_id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
      "state": "failed",
      "reason": "high",
      "old_size": 100,
      "new_size": 120,
      "created": {"at": 1700000000, "usage_percent": 82.0},
      "confirmed": {"at": 1700003600},
      "greenlit": {"at": 1700007200},
      "finished": {"at": 1700010800, "error": "quota exceeded"}
    }
  ]
}
`

const ListRecentlySucceededResponse = `
{
  "recently_succeeded_operations": [
    {
      "project_id": "88e5cad3-38e6-454f-b412-662cda03e7a1",
      "asset_type": "nfs-shares",
      "asset_id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
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
