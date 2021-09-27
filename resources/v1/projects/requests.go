package projects

import (
	"io/ioutil"

	"github.com/gophercloud/gophercloud"
	"github.com/sapcc/limes"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToProjectListParams() (map[string]string, string, error)
}

// ListOpts contains parameters for filtering a List request.
type ListOpts struct {
	Cluster  string `h:"X-Limes-Cluster-Id"`
	Detail   bool   `q:"detail"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
	Rates    string `q:"rates"`
}

// ToProjectListParams formats a ListOpts into a map of headers and a query string.
func (opts ListOpts) ToProjectListParams() (map[string]string, string, error) {
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		return nil, "", err
	}

	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return nil, "", err
	}

	return h, q.String(), nil
}

// List enumerates the projects in a specific domain.
func List(c *gophercloud.ServiceClient, domainID string, opts ListOptsBuilder) (r CommonResult) {
	url := listURL(c, domainID)
	headers := make(map[string]string)
	if opts != nil {
		h, q, err := opts.ToProjectListParams()
		if err != nil {
			r.Err = err
			return
		}
		headers = h
		url += q
	}

	resp, err := c.Get(url, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: headers,
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetOptsBuilder allows extensions to add additional parameters to the Get request.
type GetOptsBuilder interface {
	ToProjectGetParams() (map[string]string, string, error)
}

// GetOpts contains parameters for filtering a Get request.
type GetOpts struct {
	Cluster  string `h:"X-Limes-Cluster-Id"`
	Detail   bool   `q:"detail"`
	Area     string `q:"area"`
	Service  string `q:"service"`
	Resource string `q:"resource"`
	Rates    string `q:"rates"`
}

// ToProjectGetParams formats a GetOpts into a map of headers and a query string.
func (opts GetOpts) ToProjectGetParams() (map[string]string, string, error) {
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		return nil, "", err
	}

	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return nil, "", err
	}

	return h, q.String(), nil
}

// Get retrieves details on a single project, by ID.
func Get(c *gophercloud.ServiceClient, domainID string, projectID string, opts GetOptsBuilder) (r CommonResult) {
	url := getURL(c, domainID, projectID)
	headers := make(map[string]string)
	if opts != nil {
		h, q, err := opts.ToProjectGetParams()
		if err != nil {
			r.Err = err
			return
		}
		headers = h
		url += q
	}

	resp, err := c.Get(url, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: headers,
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToProjectUpdateMap() (map[string]string, map[string]interface{}, error)
}

// UpdateOpts contains parameters to update a project.
type UpdateOpts struct {
	Cluster  string             `h:"X-Limes-Cluster-Id"`
	Services limes.QuotaRequest `json:"services"`
}

// ToProjectUpdateMap formats a UpdateOpts into a map of headers and a request body.
func (opts UpdateOpts) ToProjectUpdateMap() (map[string]string, map[string]interface{}, error) {
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		return nil, nil, err
	}

	b, err := gophercloud.BuildRequestBody(opts, "project")
	if err != nil {
		return nil, nil, err
	}

	return h, b, nil
}

// Update modifies the attributes of a project and returns the response body which contains non-fatal error messages.
func Update(c *gophercloud.ServiceClient, domainID string, projectID string, opts UpdateOptsBuilder) (r UpdateResult) {
	url := updateURL(c, domainID, projectID)
	h, b, err := opts.ToProjectUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(url, b, nil, &gophercloud.RequestOpts{
		OkCodes:          []int{202},
		MoreHeaders:      h,
		KeepResponseBody: true,
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	if r.Err != nil {
		return
	}
	defer resp.Body.Close()
	r.Body, r.Err = ioutil.ReadAll(resp.Body)
	return
}

// SyncOptsBuilder allows extensions to add additional parameters to the Sync request.
type SyncOptsBuilder interface {
	ToProjectSyncParams() (map[string]string, error)
}

// SyncOpts contains parameters for filtering a Sync request.
type SyncOpts struct {
	Cluster string `h:"X-Limes-Cluster-Id"`
}

// ToProjectSyncParams formats a SyncOpts into a map of headers.
func (opts SyncOpts) ToProjectSyncParams() (map[string]string, error) {
	return gophercloud.BuildHeaders(opts)
}

// Sync schedules a sync task that pulls a project's data from the backing services
// into Limes' local database.
func Sync(c *gophercloud.ServiceClient, domainID string, projectID string, opts SyncOptsBuilder) (r SyncResult) {
	url := syncURL(c, domainID, projectID)
	headers := make(map[string]string)
	if opts != nil {
		h, err := opts.ToProjectSyncParams()
		if err != nil {
			r.Err = err
			return
		}
		headers = h
	}

	resp, err := c.Post(url, nil, nil, &gophercloud.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: headers,
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
