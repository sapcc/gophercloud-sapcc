// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const GetResponse = `
{
  "project_id": "e9141fb24eee4b3e9f25ae69cda31132",
  "project_name": "project",
  "description": "Demos and Tests",
  "parent_id": "2bac466eed364d8a92e477459e908736",
  "domain_id": "2bac466eed364d8a92e477459e908736",
  "domain_name": "domain",
  "cost_object": {
    "name": "123456789",
    "type": "IO",
    "inherited": false
  },
  "project_type": "quota",
  "responsible_primary_contact_id": "D123456",
  "responsible_primary_contact_email": "example@mail.com",
  "responsible_operator_id": null,
  "responsible_operator_email": null,
  "responsible_security_expert_id": null,
  "responsible_security_expert_email": null,
  "responsible_product_owner_id": null,
  "responsible_product_owner_email": null,
  "responsible_controller_id": null,
  "responsible_controller_email": null,
  "revenue_relevance": "generating",
  "business_criticality": "dev",
  "number_of_endusers": 100,
  "additional_information": "info",
  "changed_by": "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
  "changed_at": "2019-08-20T14:39:39.786",
  "collector": "billing.region.local",
  "region": "region",
  "is_complete": true
}
`

const ListResponse = `
[
  {
    "project_id": "e9141fb24eee4b3e9f25ae69cda31132",
    "project_name": "project",
    "description": "Demos and Tests",
    "parent_id": "2bac466eed364d8a92e477459e908736",
    "domain_id": "2bac466eed364d8a92e477459e908736",
    "domain_name": "domain",
    "cost_object": {
      "name": "123456789",
      "type": "IO",
      "inherited": false
    },
    "project_type": "quota",
    "responsible_primary_contact_id": "D123456",
    "responsible_primary_contact_email": "example@mail.com",
    "responsible_operator_id": null,
    "responsible_operator_email": null,
    "responsible_security_expert_id": null,
    "responsible_security_expert_email": null,
    "responsible_product_owner_id": null,
    "responsible_product_owner_email": null,
    "responsible_controller_id": null,
    "responsible_controller_email": null,
    "revenue_relevance": "generating",
    "business_criticality": "dev",
    "number_of_endusers": 100,
    "additional_information": "info",
    "changed_by": "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
    "changed_at": "2019-08-20T14:39:39.786",
    "collector": "billing.region.local",
    "region": "region",
    "is_complete": true
  }
]
`

const UpdateRequest = `
{
  "additional_information": "",
  "revenue_relevance" : "generating",
  "description" : "Demos and Tests",
  "responsible_primary_contact_email" : "example@mail.com",
  "cost_object" : {
     "inherited" : true
  },
  "project_id" : "e9141fb24eee4b3e9f25ae69cda31132",
  "domain_id" : "2bac466eed364d8a92e477459e908736",
  "project_name" : "project",
  "number_of_endusers" : 99,
  "responsible_primary_contact_id" : "D123456",
  "responsible_primary_contact_email" : "example@mail.com",
  "responsible_operator_email": "",
  "responsible_operator_id": "",
  "responsible_inventory_role_id": "D123456",
  "responsible_inventory_role_email": "123@mail.com",
  "responsible_infrastructure_coordinator_id": "D123456",
  "responsible_infrastructure_coordinator_email": "123@mail.com",
  "customer": "123ABC",
  "gpu_enabled": 0,
  "contains_pii_dpp_hr": 1,
  "contains_external_customer_data": 0,
  "soft_license_mode": "",
  "type_of_data": "",
  "environment": "",
  "ext_certification": {
     "c5": 0,
     "iso": 1,
     "pci": 0,
     "soc1": 0,
     "soc2": 0,
     "sox": 0
  },
  "parent_id" : "2bac466eed364d8a92e477459e908736",
  "business_criticality" : "dev"
}
`

const UpdateResponse = `
{
  "project_id": "e9141fb24eee4b3e9f25ae69cda31132",
  "project_name": "project",
  "description": "Demos and Tests",
  "parent_id": "2bac466eed364d8a92e477459e908736",
  "domain_id": "2bac466eed364d8a92e477459e908736",
  "domain_name": "domain",
  "cost_object": {
    "inherited": true
  },
  "project_type": "quota",
  "responsible_primary_contact_id": "D123456",
  "responsible_primary_contact_email": "example@mail.com",
  "responsible_operator_id": null,
  "responsible_operator_email": null,
  "responsible_inventory_role_id": "D123456",
  "responsible_inventory_role_email": "123@mail.com",
  "responsible_infrastructure_coordinator_id": "D123456",
  "responsible_infrastructure_coordinator_email": "123@mail.com",
  "customer": "123ABC",
  "gpu_enabled": 0,
  "contains_pii_dpp_hr": 1,
  "contains_external_customer_data": null,
  "type_of_data": null,
  "ext_certification": {
     "c5": 0,
     "iso": 1,
     "pci": 0,
     "soc1": 0,
     "soc2": 0,
     "sox": 0
  },
  "revenue_relevance": "generating",
  "business_criticality": "dev",
  "number_of_endusers": 99,
  "additional_information": "",
  "changed_by": "41cab08d5af96b7c64b561c639be948dc16d9b2e263a3660bfa1e096422d522e",
  "changed_at": "2019-08-26T09:09:05.457",
  "collector": "billing.region.local",
  "region": "region",
  "is_complete": true
}
`
