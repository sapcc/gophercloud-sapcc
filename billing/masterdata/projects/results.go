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
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a project
// resource.
func (r commonResult) Extract() (*Project, error) {
	var s Project
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// GetResult represents the result of a get operation. Call its Extract method
// to interpret it as a Project.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Project.
type UpdateResult struct {
	commonResult
}

// Project represents a Billing Project.
type Project struct {
	// A project ID
	ProjectID string `json:"project_id"`
	// Human-readable name for the project. Might not be unique.
	ProjectName string `json:"project_name"`
	// Technical of the domain in which the project is contained
	DomainID string `json:"domain_id"`
	// Name of the domain
	DomainName string `json:"domain_name"`
	// Description of the project
	Description string `json:"description"`
	// A project parent ID
	ParentID string `json:"parent_id"`
	// A project type
	ProjectType string `json:"project_type"`
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
	CostObject CostObject `json:"cost_object"`
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
	// The date, when the project was created.
	CreatedAt time.Time `json:"-"`
	// The date, when the project was updated.
	ChangedAt time.Time `json:"-"`
	// The ID of the user, who did the last change.
	ChangedBy string `json:"changed_by"`
	// Only contained in Server response: True, if the given masterdata are complete; Otherwise false
	IsComplete bool `json:"is_complete"`
	// Only contained in Server response: Human readable text, showing, what information are missing
	MissingAttributes string `json:"missing_attributes"`
	// Only contained in Server response: Collector of the project
	Collector string `json:"collector"`
	// Only contained in Server response: Region of the project
	Region string `json:"region"`
}

// The cost object structure
type CostObject struct {
	// Shows, if the CO is inherited. Mandatory, if name/type not set
	Inherited bool `json:"inherited"`
	// Name of the costobject. Mandatory, if inherited not true
	Name string `json:"name,omitempty"`
	// Costobject-Type Type of the costobject. Mandatory, if inherited not true
	// IO, CC, WBS, SO
	Type string `json:"type,omitempty"`
}

type ExtCertification struct {
	// C5 is a government-backed verification framework implemented by the German Federal Office for Information Security (BSI)
	C5 bool `json:"-"`
	// An ISO certification describes the process that confirms that ISO standards are being followed
	ISO bool `json:"-"`
	// PCI certification ensures the security of card data at your business through a set of requirements established by the PCI SSC
	PCI bool `json:"-"`
	// SOCx is a type of audit report that attests to the trustworthiness of services provided by a service organization
	SOC1 bool `json:"-"`
	SOC2 bool `json:"-"`
	// The law mandates strict reforms to improve financial disclosures from corporations and prevent accounting fraud
	SOX bool `json:"-"`
}

func (r *ExtCertification) UnmarshalJSON(b []byte) error {
	var s struct {
		C5   int `json:"c5"`
		ISO  int `json:"iso"`
		PCI  int `json:"pci"`
		SOC1 int `json:"soc1"`
		SOC2 int `json:"soc2"`
		SOX  int `json:"sox"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	r.C5 = s.C5 != 0
	r.ISO = s.ISO != 0
	r.PCI = s.PCI != 0
	r.SOC1 = s.SOC1 != 0
	r.SOC2 = s.SOC2 != 0
	r.SOX = s.SOX != 0

	return nil
}

func (r ExtCertification) MarshalJSON() ([]byte, error) {
	res := struct {
		C5   int `json:"c5"`
		ISO  int `json:"iso"`
		PCI  int `json:"pci"`
		SOC1 int `json:"soc1"`
		SOC2 int `json:"soc2"`
		SOX  int `json:"sox"`
	}{
		C5:   b2i(r.C5),
		ISO:  b2i(r.ISO),
		PCI:  b2i(r.PCI),
		SOC1: b2i(r.SOC1),
		SOC2: b2i(r.SOC2),
		SOX:  b2i(r.SOX),
	}
	return json.Marshal(res)
}

func (r *Project) UnmarshalJSON(b []byte) error {
	type tmp Project
	var s struct {
		tmp
		CreatedAt                    gophercloud.JSONRFC3339MilliNoZ `json:"created_at"`
		ChangedAt                    gophercloud.JSONRFC3339MilliNoZ `json:"changed_at"`
		GPUEnabled                   int                             `json:"gpu_enabled"`
		ContainsPIIDPPHR             int                             `json:"contains_pii_dpp_hr"`
		ContainsExternalCustomerData int                             `json:"contains_external_customer_data"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Project(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.ChangedAt = time.Time(s.ChangedAt)
	r.GPUEnabled = s.GPUEnabled != 0
	r.ContainsPIIDPPHR = s.ContainsPIIDPPHR != 0
	r.ContainsExternalCustomerData = s.ContainsExternalCustomerData != 0

	return nil
}

// ProjectPage is the page returned by a pager when traversing over a collection
// of projects.
type ProjectPage struct {
	pagination.SinglePageBase
}

// ExtractProjects accepts a Page struct, specifically a ProjectPage
// struct, and extracts the elements into a slice of Project structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractProjects(r pagination.Page) ([]Project, error) {
	var s []Project
	err := ExtractProjectsInto(r, &s)
	return s, err
}

func ExtractProjectsInto(r pagination.Page, v interface{}) error {
	return r.(ProjectPage).Result.ExtractIntoSlicePtr(v, "")
}
