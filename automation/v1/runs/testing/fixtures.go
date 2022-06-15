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
    "id": "1",
    "log": "Selecting nodes using filter @identity='88e5cad3-38e6-454f-b412-662cda03e7a1':\n88e5cad3-38e6-454f-b412-662cda03e7a1 automation-node\nUsing exiting artifact for revision a7af74be592a4637ae5de390b8e8888022130e63\nScheduled 1 job:\n61915ce7-f719-4b23-a163-cd1132668110\nScheduled 1 job:\n61915ce7-f719-4b23-a163-cd1132668110\n",
    "created_at": "2019-03-05T19:45:40.057Z",
    "updated_at": "2019-03-05T19:45:57.041Z",
    "repository_revision": "a7af74be592a4637ae5de390b8e8888022130e63",
    "state": "completed",
    "jobs": [
      "61915ce7-f719-4b23-a163-cd1132668110"
    ],
    "owner": {
      "id": "b81eec56-5db9-49ae-8775-880b75d38a1a",
      "name": "user",
      "domain_id": "6c2feb1a-1d38-4541-aba4-93ed61f2ccca",
      "domain_name": "project"
    },
    "automation_id": "2",
    "automation_name": "chef",
    "selector": "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
    "automation_attributes": {
      "name": "chef",
      "debug": true,
      "timeout": 3600,
      "run_list": [
        "recipe[application::app]"
      ],
      "repository": "https://github.com/org/chef.git",
      "chef_version": "12.22.5",
      "repository_revision": "master"
    }
  }
]
`

const GetResponse = `
{
  "id": "1",
  "log": "Selecting nodes using filter @identity='88e5cad3-38e6-454f-b412-662cda03e7a1':\n88e5cad3-38e6-454f-b412-662cda03e7a1 automation-node\nUsing exiting artifact for revision a7af74be592a4637ae5de390b8e8888022130e63\nScheduled 1 job:\n61915ce7-f719-4b23-a163-cd1132668110\nScheduled 1 job:\n61915ce7-f719-4b23-a163-cd1132668110\n",
  "created_at": "2019-03-05T19:45:40.057Z",
  "updated_at": "2019-03-05T19:45:57.041Z",
  "repository_revision": "a7af74be592a4637ae5de390b8e8888022130e63",
  "state": "completed",
  "jobs": [
    "61915ce7-f719-4b23-a163-cd1132668110"
  ],
  "owner": {
    "id": "b81eec56-5db9-49ae-8775-880b75d38a1a",
    "name": "user",
    "domain_id": "6c2feb1a-1d38-4541-aba4-93ed61f2ccca",
    "domain_name": "project"
  },
  "automation_id": "2",
  "automation_name": "chef",
  "selector": "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
  "automation_attributes": {
    "name": "chef",
    "debug": true,
    "timeout": 3600,
    "run_list": [
      "recipe[application::app]"
    ],
    "repository": "https://github.com/org/chef.git",
    "chef_version": "12.22.5",
    "repository_revision": "master"
  }
}
`

const CreateRequest = `
{
  "automation_id": "2",
  "selector": "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'"
}
`

const CreateResponse = `
{
  "id": "2",
  "log": null,
  "created_at": "2019-03-05T20:03:16.954Z",
  "updated_at": "2019-03-05T20:03:16.954Z",
  "repository_revision": null,
  "state": "preparing",
  "jobs": null,
  "owner": {
    "id": "b81eec56-5db9-49ae-8775-880b75d38a1a",
    "name": "user",
    "domain_id": "6c2feb1a-1d38-4541-aba4-93ed61f2ccca",
    "domain_name": "project"
  },
  "automation_id": "1",
  "automation_name": "chef",
  "selector": "@identity='88e5cad3-38e6-454f-b412-662cda03e7a1'",
  "automation_attributes": null
}
`
