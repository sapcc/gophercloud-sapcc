// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const ListResponse = `
{
  "resources": {
    "nfs-shares": {
      "checked": {"error": "cannot connect to OpenStack"},
      "asset_count": 42,
      "low_threshold": {"usage_percent": 20.0, "delay_seconds": 3600},
      "high_threshold": {"usage_percent": 80.0, "delay_seconds": 1800},
      "critical_threshold": {"usage_percent": 95.0},
      "size_constraints": {"minimum": 10, "maximum": 2000},
      "size_steps": {"percent": 20.0}
    },
    "smb-shares": {
      "checked": {"error": "cannot connect to OpenStack"},
      "asset_count": 42,
      "low_threshold": {"usage_percent": 10.0, "delay_seconds": 3600},
      "high_threshold": {"usage_percent": 50.0, "delay_seconds": 1800},
      "critical_threshold": {"usage_percent": 90.0},
      "size_constraints": {"minimum": 20, "maximum": 2000},
      "size_steps": {"percent": 10.0}
    }
  }
}
`

const GetResponse = `
{
  "checked": {"error": "cannot connect to OpenStack"},
  "asset_count": 42,
  "low_threshold": {"usage_percent": 20.0, "delay_seconds": 3600},
  "high_threshold": {"usage_percent": 80.0, "delay_seconds": 1800},
  "critical_threshold": {"usage_percent": 95.0},
  "size_constraints": {"minimum": 10, "maximum": 2000},
  "size_steps": {"percent": 20.0}
}
`

const DeleteResponse = ``

const CreateRequest = `
{
  "low_threshold": {"usage_percent": 20.0, "delay_seconds": 3600},
  "high_threshold": {"usage_percent": 80.0, "delay_seconds": 1800},
  "critical_threshold": {"usage_percent": 95.0},
  "size_constraints": {"minimum": 10, "maximum": 2000},
  "size_steps": {"percent": 20.0}
}
`

const CreateResponse = ``
