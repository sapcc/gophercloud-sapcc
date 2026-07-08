// SPDX-FileCopyrightText: 2026 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"

	"github.com/sapcc/gophercloud-sapcc/v2/audit/v1/dataplaneconfig"
)

const testProjectID = "abcd1234efgh5678ijkl9012mnop3456"

var expectedConfig = dataplaneconfig.DataplaneConfig{
	ProjectID:    testProjectID,
	Enabled:      true,
	TargetBucket: "audit-bucket",
	UpdatedAt:    time.Date(2026, 7, 8, 10, 0, 0, 0, time.UTC),
	UpdatedBy:    "u-42",
}

func TestGet(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/"+testProjectID+"/dataplane-config", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	cfg, err := dataplaneconfig.Get(t.Context(), client.ServiceClient(fakeServer), testProjectID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedConfig, *cfg)
}

func TestPut(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/"+testProjectID+"/dataplane-config", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, PutRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, PutResponse)
	})

	opts := dataplaneconfig.PutOpts{
		Enabled:      true,
		TargetBucket: "audit-bucket",
	}

	cfg, err := dataplaneconfig.Put(t.Context(), client.ServiceClient(fakeServer), testProjectID, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedConfig, *cfg)
}

func TestPutDisabledOmitsBucket(t *testing.T) {
	// When Enabled=false and TargetBucket="", the request body must omit
	// target_bucket entirely thanks to `omitempty`.
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/"+testProjectID+"/dataplane-config", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"enabled": false}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, `{
  "project_id": "`+testProjectID+`",
  "enabled": false,
  "updated_at": "2026-07-08T10:00:00Z",
  "updated_by": "u-42"
}`)
	})

	_, err := dataplaneconfig.Put(t.Context(), client.ServiceClient(fakeServer), testProjectID, dataplaneconfig.PutOpts{}).Extract()
	th.AssertNoErr(t, err)
}

func TestDelete(t *testing.T) {
	fakeServer := th.SetupHTTP()
	defer fakeServer.Teardown()

	fakeServer.Mux.HandleFunc("/projects/"+testProjectID+"/dataplane-config", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	res := dataplaneconfig.Delete(t.Context(), client.ServiceClient(fakeServer), testProjectID)
	th.AssertNoErr(t, res.Err)
}
