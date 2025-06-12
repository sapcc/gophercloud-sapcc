// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

const ListResponse = `
[
  {
    "version": 1,
    "sender": "linux",
    "request_id": "afa7a1d0-12a0-4848-ae4d-7bb7b01f126d",
    "to": "fe3da83f-919e-4b2e-8200-7acb1816c8d0",
    "timeout": 3600,
    "agent": "execute",
    "action": "tarball",
    "payload": "{\"path\":\"/\",\"environment\":{\"X\":\"y\"},\"url\":\"https://objectstore:443/v1/AUTH_3946cfbc1fda4ce19561da1df5443c86/path/to\"}",
    "status": "failed",
    "created_at": "2018-04-29T11:29:51.32156Z",
    "updated_at": "2018-04-29T11:29:51.385979Z",
    "project": "3946cfbc1fda4ce19561da1df5443c86",
    "user": {
      "domain_id": "123",
      "domain_name": "domain",
      "id": "123",
      "name": "user",
      "roles": [
        "automation_admin"
      ]
    }
  },
  {
    "version": 1,
    "sender": "linux",
    "request_id": "c6c5e3a4-9a6a-40c5-a46b-cc8f2482e3df",
    "to": "88e5cad3-38e6-454f-b412-662cda03e7a1",
    "timeout": 60,
    "agent": "execute",
    "action": "script",
    "payload": "echo \"Scritp start\"\n\nfor i in {1..10}\ndo\n\techo $i\n  sleep 1s\ndone\n\necho \"Script done\"",
    "status": "complete",
    "created_at": "2019-03-06T13:17:13.823592Z",
    "updated_at": "2019-03-06T13:17:23.853638Z",
    "project": "3946cfbc1fda4ce19561da1df5443c86",
    "user": {
      "domain_id": "123",
      "domain_name": "domain",
      "id": "123",
      "name": "user",
      "roles": [
        "automation_admin"
      ]
    }
  }
]
`

const GetResponse = `
{
  "version": 1,
  "sender": "linux",
  "request_id": "c6c5e3a4-9a6a-40c5-a46b-cc8f2482e3df",
  "to": "88e5cad3-38e6-454f-b412-662cda03e7a1",
  "timeout": 60,
  "agent": "execute",
  "action": "script",
  "payload": "echo \"Scritp start\"\n\nfor i in {1..10}\ndo\n\techo $i\n  sleep 1s\ndone\n\necho \"Script done\"",
  "status": "complete",
  "created_at": "2019-03-06T13:17:13.823592Z",
  "updated_at": "2019-03-06T13:17:23.853638Z",
  "project": "3946cfbc1fda4ce19561da1df5443c86",
  "user": {
    "domain_id": "123",
    "domain_name": "domain",
    "id": "123",
    "name": "user",
    "roles": [
      "automation_admin"
    ]
  }
}
`

const CreateRequest = `
{
  "to": "7ec336bd-fcd1-42af-a663-da578dd0b224",
  "timeout": 60,
  "agent": "execute",
  "action": "script",
  "payload": "echo \"Script start\"\n\nfor i in {1..10}\ndo\n\techo $i\n  sleep 1s\ndone\n\necho \"Script done\""
}
`

const CreateResponse = `
{
  "request_id": "2f550302-5567-44b6-8b99-fd563dc53c18"
}
`

const LogResponse = `Script start
1
2
3
4
5
6
7
8
9
10
Script done`
