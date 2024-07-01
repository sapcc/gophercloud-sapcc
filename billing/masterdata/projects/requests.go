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

package projects

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

// ListOpts is a structure that holds options for listing project masterdata.
type ListOpts struct {
	CheckCOValidity bool      `q:"checkCOValidity"`
	ExcludeDeleted  bool      `q:"excludeDeleted"`
	From            time.Time `q:"-"`
}

// ToProjectListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProjectListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	params := q.Query()

	if opts.From != (time.Time{}) {
		params.Add("from", opts.From.Format(gophercloud.RFC3339MilliNoZ))
	}

	q = &url.URL{RawQuery: params.Encode()}

	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// projects
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	serviceURL := listURL(c)
	if opts != (ListOpts{}) {
		query, err := opts.ToProjectListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		serviceURL += query
	}
	return pagination.NewPager(c, serviceURL, func(r pagination.PageResult) pagination.Page {
		return ProjectPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific project based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToProjectUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents the attributes used when updating an existing
// project.
type UpdateOpts struct {
	// Description of the project
	Description string `json:"description"`
	// SAP-User-Id of primary contact for the project
	ResponsiblePrimaryContactID string `json:"responsible_primary_contact_id"`
	// Email-address of primary contact for the project
	ResponsiblePrimaryContactEmail string `json:"responsible_primary_contact_email"`
	// SAP-User-Id of the person who is responsible for operating the project
	ResponsibleOperatorID string `json:"responsible_operator_id"`
	// Email-address or DL of the person/group who is operating the project
	ResponsibleOperatorEmail string `json:"responsible_operator_email"`
	// ID of the Person/entity responsible to correctly maintain assets in SAP's Global DC HW asset inventory SISM/CCIR
	ResponsibleInventoryRoleID string `json:"responsible_inventory_role_id"`
	// Email of the Person/entity responsible to correctly maintain assets in SAP's Global DC HW asset inventory SISM/CCIR
	ResponsibleInventoryRoleEmail string `json:"responsible_inventory_role_email"`
	// ID of the infrastructure coordinator
	ResponsibleInfrastructureCoordinatorID string `json:"responsible_infrastructure_coordinator_id"`
	// Email address of the infrastructure coordinator
	ResponsibleInfrastructureCoordinatorEmail string `json:"responsible_infrastructure_coordinator_email"`
	// Indicating if the project is directly or indirectly creating revenue
	// Allowed values: [generating, enabling, other]
	RevenueRelevance string `json:"revenue_relevance"`
	// Indicates how important the project for the business is. Possible values: [dev,test,prod]
	// Allowed values: [dev, test, prod]
	BusinessCriticality string `json:"business_criticality"`
	// If the number is unclear, always provide the lower end --> means always > number_of_endusers (-1 indicates that it is infinite)
	NumberOfEndusers int `json:"number_of_endusers"`
	// Name of the Customer (CCIR/BPD Key)
	Customer string `json:"customer"`
	// Freetext field for additional information for project
	AdditionalInformation string `json:"additional_information"`
	// The cost object structure
	CostObject CostObject `json:"cost_object" required:"true"`
	// Build environment of the project
	// Allowed values: [Prod,QA,Admin,DEV,Demo,Train,Sandbox,Lab,Test]
	Environment string `json:"environment"`
	// Software License Mode
	// Allowed values: [Revenue Generating,Training & Demo,Development,Test & QS,Administration,Make,Virtualization-Host,Productive]
	SoftLicenseMode string `json:"soft_license_mode"`
	// Input parameter for KRITIS flag in CCIR
	// Allowed values: [SAP Business Process,Customer Cloud Service,Customer Business Process,Training & Demo Cloud]
	TypeOfData string `json:"type_of_data"`
	// Uses GPUs
	GPUEnabled bool `json:"-"`
	// Required to harmonize with CALM and for further calculation of CIA
	ContainsPIIDPPHR bool `json:"-"`
	// Required to harmonize with CALM and for further calculation of CIA
	ContainsExternalCustomerData bool `json:"-"`
	// Information about whether there is any external certification present in this project
	ExtCertification *ExtCertification `json:"ext_certification,omitempty"`
	// when the token is not scoped
	// A project ID
	ProjectID string `json:"project_id,omitempty"`
	// Human-readable name for the project. Might not be unique.
	ProjectName string `json:"project_name,omitempty"`
	// Technical of the domain in which the project is contained
	DomainID string `json:"domain_id,omitempty"`
	// Name of the domain
	DomainName string `json:"domain_name,omitempty"`
	// A project parent ID
	ParentID string `json:"parent_id,omitempty"`
	// A project type
	ProjectType string `json:"project_type,omitempty"`
}

func (opts *UpdateOpts) UnmarshalJSON(b []byte) error {
	type tmp UpdateOpts
	var s struct {
		tmp
		GPUEnabled                   int `json:"gpu_enabled"`
		ContainsPIIDPPHR             int `json:"contains_pii_dpp_hr"`
		ContainsExternalCustomerData int `json:"contains_external_customer_data"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*opts = UpdateOpts(s.tmp)
	opts.GPUEnabled = s.GPUEnabled != 0
	opts.ContainsPIIDPPHR = s.ContainsPIIDPPHR != 0
	opts.ContainsExternalCustomerData = s.ContainsExternalCustomerData != 0

	return nil
}

func (opts UpdateOpts) MarshalJSON() ([]byte, error) {
	type tmp UpdateOpts
	res := struct {
		tmp
		GPUEnabled                   int `json:"gpu_enabled"`
		ContainsPIIDPPHR             int `json:"contains_pii_dpp_hr"`
		ContainsExternalCustomerData int `json:"contains_external_customer_data"`
	}{
		tmp:                          tmp(opts),
		GPUEnabled:                   b2i(opts.GPUEnabled),
		ContainsPIIDPPHR:             b2i(opts.ContainsPIIDPPHR),
		ContainsExternalCustomerData: b2i(opts.ContainsExternalCustomerData),
	}
	return json.Marshal(res)
}

// ToProjectUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToProjectUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update accepts a UpdateOpts struct and updates an existing project using
// the values provided.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToProjectUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Put(ctx, updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func ProjectToUpdateOpts(project *Project) UpdateOpts {
	return UpdateOpts{
		Description:                               project.Description,
		ResponsiblePrimaryContactID:               project.ResponsiblePrimaryContactID,
		ResponsiblePrimaryContactEmail:            project.ResponsiblePrimaryContactEmail,
		ResponsibleOperatorID:                     project.ResponsibleOperatorID,
		ResponsibleOperatorEmail:                  project.ResponsibleOperatorEmail,
		ResponsibleInventoryRoleID:                project.ResponsibleInventoryRoleID,
		ResponsibleInventoryRoleEmail:             project.ResponsibleInventoryRoleEmail,
		ResponsibleInfrastructureCoordinatorID:    project.ResponsibleInfrastructureCoordinatorID,
		ResponsibleInfrastructureCoordinatorEmail: project.ResponsibleInfrastructureCoordinatorEmail,
		RevenueRelevance:                          project.RevenueRelevance,
		BusinessCriticality:                       project.BusinessCriticality,
		NumberOfEndusers:                          project.NumberOfEndusers,
		Customer:                                  project.Customer,
		AdditionalInformation:                     project.AdditionalInformation,
		CostObject:                                project.CostObject,
		Environment:                               project.Environment,
		SoftLicenseMode:                           project.SoftLicenseMode,
		TypeOfData:                                project.TypeOfData,
		GPUEnabled:                                project.GPUEnabled,
		ContainsPIIDPPHR:                          project.ContainsPIIDPPHR,
		ContainsExternalCustomerData:              project.ContainsExternalCustomerData,
		ExtCertification:                          project.ExtCertification,
		// unscoped tokens
		ProjectID:   project.ProjectID,
		ProjectName: project.ProjectName,
		DomainID:    project.DomainID,
		DomainName:  project.DomainName,
		ParentID:    project.ParentID,
		ProjectType: project.ProjectType,
	}
}
