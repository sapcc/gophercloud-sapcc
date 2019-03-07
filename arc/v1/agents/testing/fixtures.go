package testing

const ListResponse = `
[
  {
    "display_name": "instance1",
    "agent_id": "88e5cad3-38e6-454f-b412-662cda03e7a1",
    "project": "3946cfbc1fda4ce19561da1df5443c86",
    "organization": "41aac04ce58c428b9ed2262798d0d336",
    "created_at": "2018-03-13T10:44:27.432827Z",
    "updated_at": "2019-03-06T10:02:19.062626Z",
    "updated_with": "a76ddf9f-748d-421e-bcd1-c1a6afc922e4",
    "updated_by": "linux"
  },
  {
    "display_name": "instance2",
    "agent_id": "7bf82bb6-61a6-4d01-aa50-16e19d1dca99",
    "project": "3946cfbc1fda4ce19561da1df5443c86",
    "organization": "41aac04ce58c428b9ed2262798d0d336",
    "created_at": "2018-11-12T10:17:12.455872Z",
    "updated_at": "2018-12-03T08:48:57.40089Z",
    "updated_with": "40db2bb1-e6b9-4c64-8353-fae5553a0092",
    "updated_by": "linux"
  }
]
`

const GetResponse = `
{
  "display_name": "instance1",
  "agent_id": "88e5cad3-38e6-454f-b412-662cda03e7a1",
  "project": "3946cfbc1fda4ce19561da1df5443c86",
  "organization": "41aac04ce58c428b9ed2262798d0d336",
  "created_at": "2018-03-13T10:44:27.432827Z",
  "updated_at": "2019-03-06T10:02:19.062626Z",
  "updated_with": "a76ddf9f-748d-421e-bcd1-c1a6afc922e4",
  "updated_by": "linux"
}
`

const InitResponseJson = `
{
  "token": "4d523051-089f-41ce-aaf7-727fee19c28a",
  "url": "https://arc.example.com/api/v1/agents/init/4d523051-089f-41ce-aaf7-727fee19c28a",
  "endpoint_url": "tls://arc-broker.example.com:8883",
  "update_url": "https://stable.arc.example.com",
  "renew_cert_url": "https://example.com/renew"
}
`

const InitResponseCloudConfig = "#cloud-config"

const InitResponseShell = "#!/bin/sh"

const InitResponsePowerShell = "#ps1_sysnative"

const GetTagsResponse = `
{
  "pool": "green",
  "landscape": "staging"
}
`

const GetFactsResponse = `
{
  "agents": {
    "chef": "enabled",
    "execute": "enabled",
    "rpc": "enabled"
  },
  "os": "linux",
  "platform": "ubuntu",
  "platform_family": "debian",
  "platform_version": "16.04",
  "project": "3946cfbc1fda4ce19561da1df5443c86"
}
`
