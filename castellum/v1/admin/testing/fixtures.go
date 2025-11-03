// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const ResourceScrapeErrorsResponse = `
{
  "resource_scrape_errors": [
    {
      "asset_type": "nfs-shares",
      "checked": {"error": "cannot connect to backend"},
      "domain_id": "d7a35a2e-3b6a-4b3c-8d5e-9f0a1b2c3d4e",
      "project_id": "88e5cad3-38e6-454f-b412-662cda03e7a1"
    }
  ]
}
`

const AssetScrapeErrorsResponse = `
{
  "asset_scrape_errors": [
    {
      "asset_id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
      "asset_type": "nfs-shares",
      "checked": {"error": "share not found"},
      "domain_id": "d7a35a2e-3b6a-4b3c-8d5e-9f0a1b2c3d4e",
      "project_id": "88e5cad3-38e6-454f-b412-662cda03e7a1"
    }
  ]
}
`

const AssetResizeErrorsResponse = `
{
  "asset_resize_errors": [
    {
      "asset_id": "05620cba-c0c1-4e75-a5e9-b5decf643dc7",
      "asset_type": "nfs-shares",
      "domain_id": "d7a35a2e-3b6a-4b3c-8d5e-9f0a1b2c3d4e",
      "project_id": "88e5cad3-38e6-454f-b412-662cda03e7a1",
      "old_size": 100,
      "new_size": 120,
      "finished": {"at": 1700010800, "error": "quota exceeded"}
    }
  ]
}
`
