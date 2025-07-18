// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"

	v1 "github.com/sapcc/gophercloud-sapcc/v2/metis/v1"
)

type IPAddress struct {
	IP          string `json:"ipaddress"`
	PortUUID    string `json:"port"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	DeviceID    string `json:"deviceID,omitempty"`
	DeviceOwner string `json:"deviceOwner,omitempty"`
	FixedPortID string `json:"fixedPortID,omitempty"`
	FixedIP     string `json:"fixedIP,omitempty"`
	NetworkID   string `json:"networkID,omitempty"`
	NetworkName string `json:"networkName,omitempty"`
	SubnetID    string `json:"subnetID,omitempty"`
	SubnetName  string `json:"subnetName,omitempty"`
	DomainID    string `json:"domainID,omitempty"`
	DomainName  string `json:"domainName,omitempty"`
	ProjectID   string `json:"projectID,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	Created     string `json:"created,omitempty"`
	LastChanged string `json:"lastChanged,omitempty"`
}

// Extract accepts a Page struct, specifically an v1.CommonPage struct,
// and extracts the elements into a slice of IPAddress structs.
func Extract(r pagination.Page) ([]IPAddress, error) {
	var s struct {
		IPAdresses []IPAddress `json:"items"`
	}
	if err := (r.(v1.CommonPage)).ExtractInto(&s); err != nil {
		return nil, err
	}
	return s.IPAdresses, nil
}

// GetResult represents the result of a get operation. Call its Extract method
// to interpret it as a IPAddress.
type GetResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a IPAddress
// resource.
func (r GetResult) Extract() (*IPAddress, error) {
	var s struct {
		Data struct {
			Item IPAddress `json:"item"`
		} `json:"data"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data.Item, nil
}

func (r GetResult) ExtractInto(v any) error {
	return r.ExtractIntoStructPtr(v, "")
}
