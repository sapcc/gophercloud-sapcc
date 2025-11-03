// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"os"
	"testing"

	"github.com/sapcc/gophercloud-sapcc/v2/internal/acceptance/tools"

	th "github.com/gophercloud/gophercloud/v2/testhelper"

	"github.com/sapcc/gophercloud-sapcc/v2/castellum/v1/resources"
)

func TestResourceList(t *testing.T) {
	client, err := NewCastellumV1Client(t.Context())
	th.AssertNoErr(t, err)

	projectID := os.Getenv("OS_PROJECT_ID")
	if projectID == "" {
		t.Skip("OS_PROJECT_ID must be set for this acceptance test")
	}

	allResources, err := resources.List(t.Context(), client, projectID, resources.ListOpts{}).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, allResources)

	if len(allResources) == 0 {
		t.Log("No resources configured for autoscaling in this project")
	}
}
