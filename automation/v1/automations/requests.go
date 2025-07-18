// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package automations

import (
	"context"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAutomationListQuery() (string, error)
}

// ListOpts allows the listing of paginated collections through the API. Page
// and PerPage are used for pagination.
type ListOpts struct {
	Page    int `q:"page"`
	PerPage int `q:"per_page"`
}

// ToAutomationListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAutomationListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// automations. It accepts a ListOpts struct.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToAutomationListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := AutomationPage{pagination.MarkerPageBase{PageResult: r}}
		p.Owner = p
		return p
	})
}

// Get retrieves a specific automation based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Get(ctx, getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAutomationCreateMap() (map[string]any, error)
}

// CreateOpts represents the attributes used when creating a new automation.
type CreateOpts struct {
	Name       string `json:"name" required:"true"`
	Repository string `json:"repository" required:"true"`
	// RepositoryRevision defaults to master, when Type is Chef
	RepositoryRevision string `json:"repository_revision,omitempty"`
	// RepositoryCredentials credentials needed to access the repository.
	// e.g.: git token or ssh key
	RepositoryCredentials string `json:"repository_credentials,omitempty"`
	// Timeout defaults to 3600. Must be within 1-86400
	Timeout int `json:"timeout,omitempty"`
	// Tags don't work
	Tags map[string]string `json:"tags,omitempty"`
	Type string            `json:"type" required:"true"`

	// RunList is required only, when Type is Chef
	RunList []string `json:"run_list,omitempty"`
	// ChefAttributes can be set only, when Type is Chef
	ChefAttributes map[string]any `json:"chef_attributes,omitempty"`
	LogLevel       string         `json:"log_level,omitempty"`
	// Debug can be set only, when Type is Chef
	Debug bool `json:"debug,omitempty"`
	// ChefVersion can be set only, when Type is Chef
	ChefVersion string `json:"chef_version,omitempty"`

	// Path is required only, when Type is Script
	Path string `json:"path,omitempty"`
	// Path can be set only, when Type is Script
	Arguments []string `json:"arguments,omitempty"`
	// Environment can be set only, when Type is Script
	Environment map[string]string `json:"environment,omitempty"`
}

// ToAutomationCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToAutomationCreateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and creates a new automation using the
// values provided.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToAutomationCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Post(ctx, createURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToAutomationUpdateMap() (map[string]any, error)
}

// UpdateOpts represents the attributes used when updating an existing
// automation.
type UpdateOpts struct {
	Name       string `json:"name,omitempty"`
	Repository string `json:"repository,omitempty"`
	// Repository revision can be unset to empty only for Script Type
	RepositoryRevision *string `json:"repository_revision,omitempty"`
	// RepositoryCredentials credentials needed to access the repository.
	// e.g.: git token or ssh key
	RepositoryCredentials *string `json:"repository_credentials,omitempty"`
	// Timeout defaults to 3600. Must be within 1-86400
	Timeout int `json:"timeout,omitempty"`
	// Tags don't work
	Tags map[string]string `json:"tags,omitempty"`

	// Chef
	RunList        []string       `json:"run_list,omitempty"`
	ChefAttributes map[string]any `json:"chef_attributes,omitempty"`
	LogLevel       *string        `json:"log_level,omitempty"`
	Debug          *bool          `json:"debug,omitempty"`
	ChefVersion    *string        `json:"chef_version,omitempty"`

	// Script
	Path        *string           `json:"path,omitempty"`
	Arguments   []string          `json:"arguments,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
}

// ToAutomationUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToAutomationUpdateMap() (map[string]any, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.Tags != nil && len(opts.Tags) == 0 {
		b["tags"] = nil
	}

	if opts.RunList != nil && len(opts.RunList) == 0 {
		b["run_list"] = nil
	}

	if opts.ChefAttributes != nil && len(opts.ChefAttributes) == 0 {
		b["chef_attributes"] = nil
	}

	if opts.Arguments != nil && len(opts.Arguments) == 0 {
		b["arguments"] = nil
	}

	if opts.Environment != nil && len(opts.Environment) == 0 {
		b["environment"] = nil
	}

	return b, nil
}

// Update accepts a UpdateOpts struct and updates an existing automation using
// the values provided.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToAutomationUpdateMap()
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

// Delete accepts a unique ID and deletes the automation associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	//nolint:bodyclose // already handled by gophercloud
	resp, err := c.Delete(ctx, deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
