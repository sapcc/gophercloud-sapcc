// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	fake "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/networking/v2/bgpvpn/interconnections"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	fields := []string{"id", "name"}
	listOpts := interconnections.ListOpts{
		Fields: fields,
	}
	th.Mux.HandleFunc("/v2.0/interconnection/interconnections",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, http.MethodGet)
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

			err := r.ParseForm()
			th.AssertNoErr(t, err)
			th.AssertDeepEquals(t, r.Form["fields"], fields)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, ListInterconnectionsResponse)
		})
	count := 0

	err := interconnections.List(fake.ServiceClient(), listOpts).EachPage(
		t.Context(),
		func(ctx context.Context, page pagination.Page) (bool, error) {
			count++
			actual, err := interconnections.ExtractInterconnections(page)
			if err != nil {
				t.Errorf("Failed to extract Interconnections: %v", err)
				return false, nil
			}

			th.CheckDeepEquals(t, interconnectionsList, actual)

			return true, nil
		})
	th.AssertNoErr(t, err)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	id := "a943ab0b-8b32-47dd-805b-4d17b7e15359"
	th.Mux.HandleFunc("/v2.0/interconnection/interconnections/"+id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, GetInterconnectionResponse)
	})

	s, err := interconnections.Get(t.Context(), fake.ServiceClient(), id).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, *s, interconnectionGet)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/interconnection/interconnections", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateInterconnectionRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, GetInterconnectionResponse)
	})

	r, err := interconnections.Create(t.Context(), fake.ServiceClient(), interconnectionCreate).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, *r, interconnectionGet)
}

func TestDelete(t *testing.T) {
	id := "0f9d472a-908f-40f5-8574-b4e8a63ccbf0"
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/interconnection/interconnections/"+id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := interconnections.Delete(t.Context(), fake.ServiceClient(), id).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUpdate(t *testing.T) {
	id := "a943ab0b-8b32-47dd-805b-4d17b7e15359"
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/interconnection/interconnections/"+id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateInterconnectionRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateInterconnectionResponse)
	})

	name := "interconnection2"
	state := "WAITING_REMOTE"
	opts := interconnections.UpdateOpts{
		Name:  &name,
		State: &state,
	}

	r, err := interconnections.Update(t.Context(), fake.ServiceClient(), id, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, r.Name, name)
	th.AssertDeepEquals(t, r.State, state)
}
