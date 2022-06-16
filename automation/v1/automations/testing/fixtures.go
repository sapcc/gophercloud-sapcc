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
    "id": 1,
    "type": "Script",
    "name": "script",
    "project_id": "3946cfbc1fda4ce19561da1df5443c86",
    "repository": "https://github.com/org/script.git",
    "repository_revision": "master",
    "repository_authentication_enabled": false,
    "timeout": 3600,
    "tags": null,
    "created_at": "2018-04-29T11:39:13.412Z",
    "updated_at": "2018-04-29T11:39:13.412Z",
    "run_list": null,
    "chef_attributes": null,
    "log_level": null,
    "debug": false,
    "chef_version": null,
    "path": "/install_nginx.sh",
    "arguments": [],
    "environment": {
      "X": "y"
    }
  },
  {
    "id": 2,
    "type": "Chef",
    "name": "chef",
    "project_id": "3946cfbc1fda4ce19561da1df5443c86",
    "repository": "https://github.com/org/chef.git",
    "repository_revision": "master",
    "repository_authentication_enabled": false,
    "timeout": 3600,
    "tags": null,
    "created_at": "2018-12-27T14:20:08.521Z",
    "updated_at": "2018-12-28T13:05:52.241Z",
    "run_list": [
      "recipe[application::app]"
    ],
    "chef_attributes": {},
    "log_level": null,
    "debug": true,
    "chef_version": "12.22.5",
    "path": null,
    "arguments": null,
    "environment": null
  }
]
`

const GetResponse = `
{
  "id": 2,
  "type": "Chef",
  "name": "chef",
  "project_id": "3946cfbc1fda4ce19561da1df5443c86",
  "repository": "https://github.com/org/chef.git",
  "repository_revision": "master",
  "repository_authentication_enabled": false,
  "timeout": 3600,
  "tags": null,
  "created_at": "2018-12-27T14:20:08.521Z",
  "updated_at": "2018-12-28T13:05:52.241Z",
  "run_list": [
    "recipe[application::app]"
  ],
  "chef_attributes": {},
  "log_level": null,
  "debug": true,
  "chef_version": "12.22.5",
  "path": null,
  "arguments": null,
  "environment": null
}
`

const CreateRequest = `
{
  "type": "Chef",
  "name": "chef",
  "repository": "https://github.com/org/chef.git",
  "repository_revision": "master",
  "timeout": 3600,
  "run_list": [
    "recipe[application::app]"
  ],
  "debug": true,
  "chef_version": "12.22.5"
}
`

const CreateResponse = `
{
  "id": 2,
  "type": "Chef",
  "name": "chef",
  "project_id": "3946cfbc1fda4ce19561da1df5443c86",
  "repository": "https://github.com/org/chef.git",
  "repository_revision": "master",
  "timeout": 3600,
  "tags": null,
  "created_at": "2018-12-27T14:20:08.521Z",
  "updated_at": "2018-12-28T13:05:52.241Z",
  "run_list": [
    "recipe[application::app]"
  ],
  "chef_attributes": {},
  "log_level": null,
  "debug": true,
  "chef_version": "12.22.5",
  "path": null,
  "arguments": null,
  "environment": null
}
`

const UpdateRequest = `
{
  "debug": false
}
`

const UpdateResponse = `
{
  "id": 2,
  "type": "Chef",
  "name": "chef",
  "project_id": "3946cfbc1fda4ce19561da1df5443c86",
  "repository": "https://github.com/org/chef.git",
  "repository_revision": "master",
  "repository_authentication_enabled": false,
  "timeout": 3600,
  "tags": null,
  "created_at": "2018-12-27T14:20:08.521Z",
  "updated_at": "2018-12-28T13:05:52.241Z",
  "run_list": [
    "recipe[application::app]"
  ],
  "chef_attributes": {},
  "log_level": null,
  "debug": false,
  "chef_version": "12.22.5",
  "path": null,
  "arguments": null,
  "environment": null
}
`

const UpdateRequestCreds = `
{
  "repository_credentials": "foobar"
}
`

const UpdateResponseCreds = `
{
  "id": 2,
  "type": "Chef",
  "name": "chef",
  "project_id": "3946cfbc1fda4ce19561da1df5443c86",
  "repository": "https://github.com/org/chef.git",
  "repository_revision": "master",
  "repository_authentication_enabled": true,
  "timeout": 3600,
  "tags": null,
  "created_at": "2018-12-27T14:20:08.521Z",
  "updated_at": "2018-12-28T13:05:52.241Z",
  "run_list": [
    "recipe[application::app]"
  ],
  "chef_attributes": {},
  "log_level": null,
  "debug": true,
  "chef_version": "12.22.5",
  "path": null,
  "arguments": null,
  "environment": null
}
`

const RemoveRunListRequest = `
{
  "tags": null,
  "run_list": null
}
`

const RemoveRunListResponse = `
{
  "id": 2,
  "type": "Chef",
  "name": "chef",
  "project_id": "3946cfbc1fda4ce19561da1df5443c86",
  "repository": "https://github.com/org/chef.git",
  "repository_revision": "master",
  "timeout": 3600,
  "tags": null,
  "created_at": "2018-12-27T14:20:08.521Z",
  "updated_at": "2018-12-28T13:05:52.241Z",
  "run_list": null,
  "chef_attributes": {},
  "log_level": null,
  "debug": true,
  "chef_version": "12.22.5",
  "path": null,
  "arguments": null,
  "environment": null
}
`
