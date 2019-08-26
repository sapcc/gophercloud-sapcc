package projects

import (
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List returns a Pager which allows you to iterate over a collection of
// projects. It accepts a ListOpts struct.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return ProjectPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific project based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
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
	// A project ID
	ProjectID *string `json:"project_id" required:"true"`
	// Human-readable name for the project. Might not be unique.
	ProjectName *string `json:"project_name,omitempty"`
	// Technical of the domain in which the project is contained
	DomainID *string `json:"domain_id" required:"true"`
	// Name of the domain
	DomainName *string `json:"domain_name,omitempty"`
	// Description of the project
	Description *string `json:"description,omitempty"`
	// A project parent ID
	ParentID *string `json:"parent_id,omitempty"`
	// A project type
	ProjectType *string `json:"project_type,omitempty"`
	// SAP-User-Id of primary contact for the project
	ResponsiblePrimaryContactID *string `json:"responsible_primary_contact_id,omitempty"`
	// Email-address of primary contact for the project
	ResponsiblePrimaryContactEmail *string `json:"responsible_primary_contact_email,omitempty"`
	// SAP-User-Id of the person who is responsible for operating the project
	ResponsibleOperatorID *string `json:"responsible_operator_id,omitempty"`
	// Email-address or DL of the person/group who is operating the project
	ResponsibleOperatorEmail *string `json:"responsible_operator_email,omitempty"`
	// SAP-User-Id of the person who is responsible for the security of the project
	ResponsibleSecurityExpertID *string `json:"responsible_security_expert_id,omitempty"`
	// Email-address or DL of the person/group who is responsible for the security of the project
	ResponsibleSecurityExpertEmail *string `json:"responsible_security_expert_email,omitempty"`
	// SAP-User-Id of the product owner
	ResponsibleProductOwnerID *string `json:"responsible_product_owner_id,omitempty"`
	// Email-address or DL of the product owner
	ResponsibleProductOwnerEmail *string `json:"responsible_product_owner_email,omitempty"`
	// SAP-User-Id of the controller who is responsible for the project / the costobject
	ResponsibleControllerID *string `json:"responsible_controller_id,omitempty"`
	// Email-address or DL of the person/group who is controlling the project / the costobject
	ResponsibleControllerEmail *string `json:"responsible_controller_email,omitempty"`
	// Indicating if the project is directly or indirectly creating revenue
	// Allowed values: [generating, enabling, other]
	RevenueRelevance *string `json:"revenue_relevance,omitempty"`
	// Indicates how important the project for the business is. Possible values: [dev,test,prod]
	// Allowed values: [dev, test, prod]
	BusinessCriticality *string `json:"business_criticality,omitempty"`
	// If the number is unclear, always provide the lower end --> means always > number_of_endusers (-1 indicates that it is infinite)
	NumberOfEndusers *int `json:"number_of_endusers,omitempty"`
	// Freetext field for additional information for project
	AdditionalInformation *string `json:"additional_information,omitempty"`
	// The cost object structure
	CostObject *CostObject `json:"cost_object,omitempty"`

	// The date, when the project was created.
	CreatedAt *time.Time `json:"-"`
	// The date, when the project was updated.
	ChangedAt *time.Time `json:"-"`
	// The ID of the user, who did the last change.
	ChangedBy *string `json:"changed_by,omitempty"`

	// Only contained in Server response: True, if the given masterdata are complete; Otherwise false
	IsComplete *bool `json:"is_complete,omitempty"`
	// Only contained in Server response: Human readable text, showing, what information are missing
	MissingAttributes *string `json:"missing_attributes,omitempty"`
	// Only contained in Server response: Collector of the project
	Collector *string `json:"collector,omitempty"`
	// Only contained in Server response: Region of the project
	Region *string `json:"region,omitempty"`
}

// ToProjectUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToProjectUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.CreatedAt != nil && *opts.CreatedAt != (time.Time{}) {
		b["created_at"] = opts.CreatedAt.Format(gophercloud.RFC3339MilliNoZ)
	}

	if opts.ChangedAt != nil && *opts.ChangedAt != (time.Time{}) {
		b["changed_at"] = opts.ChangedAt.Format(gophercloud.RFC3339MilliNoZ)
	}

	return b, nil
}

// Update accepts a UpdateOpts struct and updates an existing project using
// the values provided.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToProjectUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
