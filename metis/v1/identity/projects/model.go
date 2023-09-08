// Copyright 2023 SAP SE
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

// Project represents a Keystone project.
type Project struct {
	Name          string        `json:"name,omitempty"`
	UUID          string        `json:"uuid"`
	Description   string        `json:"description,omitempty"`
	DomainName    string        `json:"domainName,omitempty"`
	DomainUUID    string        `json:"domainUUID,omitempty"`
	CBRMasterdata CBRMasterdata `json:"cbrMasterdata,omitempty"`
	Users         []User        `json:"users"`
}

// CBRMasterdata represents the CBR masterdata of a Keystone project.
type CBRMasterdata struct {
	CostObjectName                  string                 `json:"costObjectName,omitempty"`
	CostObjectType                  string                 `json:"costObjectType,omitempty"`
	CostObjectInherited             bool                   `json:"costObjectInherited,omitempty"`
	BusinessCriticality             string                 `json:"businessCriticality,omitempty"`
	RevenueRelevance                string                 `json:"revenueRelevance,omitempty"`
	NumberOfEndusers                int                    `json:"numberOfEndusers,omitempty"`
	PrimaryContactUserID            string                 `json:"primaryContactUserID,omitempty"`
	PrimaryContactEmail             string                 `json:"primaryContactEmail,omitempty"`
	OperatorUserID                  string                 `json:"operatorUserID,omitempty"`
	OperatorEmail                   string                 `json:"operatorEmail,omitempty"`
	InventoryRoleUserID             string                 `json:"inventoryRoleUserID,omitempty"`
	InventoryRoleEmail              string                 `json:"inventoryRoleEmail,omitempty"`
	InfrastructureCoordinatorUserID string                 `json:"infrastructureCoordinatorUserID,omitempty"`
	InfrastructureCoordinatorEmail  string                 `json:"infrastructureCoordinatorEmail,omitempty"`
	ExternalCertifications          ExternalCertifications `json:"externalCertifications,omitempty"`
	GPUEnabled                      bool                   `json:"gpuEnabled,omitempty"`
	ContainsPIIDPPHR                bool                   `json:"containsPIIDPPHR,omitempty"`
	ContainsExternalCustomerData    bool                   `json:"containsExternalCustomerData,omitempty"`
}

// ExternalCertifications represents the external certifications of a Keystone project.
type ExternalCertifications struct {
	ISO  bool `json:"ISO,omitempty"`
	PCI  bool `json:"PCI,omitempty"`
	SOC1 bool `json:"SOC1,omitempty"`
	SOC2 bool `json:"SOC2,omitempty"`
	C5   bool `json:"C5,omitempty"`
	SOX  bool `json:"SOX,omitempty"`
}

// User represents a Keystone user.
type User struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
}
