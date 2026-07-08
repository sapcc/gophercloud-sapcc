// SPDX-FileCopyrightText: 2026 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const GetResponse = `
{
  "project_id": "abcd1234efgh5678ijkl9012mnop3456",
  "enabled": true,
  "target_bucket": "audit-bucket",
  "updated_at": "2026-07-08T10:00:00Z",
  "updated_by": "u-42"
}
`

const PutRequest = `
{
  "enabled": true,
  "target_bucket": "audit-bucket"
}
`

const PutResponse = GetResponse
