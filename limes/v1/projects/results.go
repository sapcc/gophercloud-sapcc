package projects

import (
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
