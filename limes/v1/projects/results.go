package projects

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/sapcc/limes/pkg/reports"
)

type ProjectPage struct {
	pagination.SinglePageBase
}

func (r ProjectPage) IsEmpty() (bool, error) {
	addresses, err := ExtractProjects(r)
	return len(addresses) == 0, err
}

func ExtractProjects(r pagination.Page) ([]reports.Project, error) {
	var s struct {
		Projects []reports.Project `json:"projects"`
	}

	err := (r.(ProjectPage)).ExtractInto(&s)
	return s.Projects, err
}

// GetResult represents the result of a get operation.
type GetResult struct {
	gophercloud.Result
}

// Extract is a function that extracts a service from a GetResult.
func (r GetResult) Extract() (*reports.Project, error) {
	var s struct {
		Project *reports.Project `json:"project"`
	}
	err := r.ExtractInto(&s)
	return s.Project, err
}
