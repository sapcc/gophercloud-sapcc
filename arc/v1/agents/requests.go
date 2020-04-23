package agents

import (
	"io/ioutil"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAgentListQuery() (string, error)
}

// ListOpts allows the filtering of paginated collections through the API.
// Filtering is achieved by passing in filter value. Page and PerPage are used
// for pagination.
type ListOpts struct {
	Page    int `q:"page"`
	PerPage int `q:"per_page"`
	// E.g. '@os = "darwin" OR (landscape = "staging" AND pool = "green")'
	// where:
	// @fact - fact
	// tag - tag
	Filter string `q:"q"`
}

// ToAgentListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAgentListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// agents. It accepts a ListOpts struct, which allows you to filter the
// returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToAgentListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := AgentPage{pagination.MarkerPageBase{PageResult: r}}
		p.MarkerPageBase.Owner = p
		return p
	})
}

// Get retrieves a specific agent based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// InitOptsBuilder allows extensions to add additional parameters to the
// Init request.
type InitOptsBuilder interface {
	ToAgentInitMap() (map[string]string, error)
}

// InitOpts represents the attributes used when initializing a new agent.
type InitOpts struct {
	// Valid options:
	// * application/json
	// * text/x-shellscript
	// * text/x-powershellscript
	// * text/cloud-config
	Accept string `h:"Accept" required:"true"`
}

// ToAgentInitMap formats a InitOpts into a map of headers.
func (opts InitOpts) ToAgentInitMap() (map[string]string, error) {
	return gophercloud.BuildHeaders(opts)
}

// Init accepts an InitOpts struct and initializes a new agent using the values
// provided.
func Init(c *gophercloud.ServiceClient, opts InitOptsBuilder) (r InitResult) {
	h, err := opts.ToAgentInitMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Request("POST", initURL(c), &gophercloud.RequestOpts{
		MoreHeaders:      h,
		OkCodes:          []int{200},
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

// Delete accepts a unique ID and deletes the agent associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateTagsBuilder allows extensions to add additional parameters to
// the CreateTags request.
type CreateTagsBuilder interface {
	ToTagsCreateMap() (map[string]string, error)
}

// ToTagsCreateMap converts a Tags into a request body.
func (opts Tags) ToTagsCreateMap() (map[string]string, error) {
	return opts, nil
}

// CreateTags adds/updates tags for a given agent.
func CreateTags(client *gophercloud.ServiceClient, agentID string, opts Tags) (r TagsErrResult) {
	b, err := opts.ToTagsCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(tagsURL(client, agentID), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetTags lists tags for a given agent.
func GetTags(client *gophercloud.ServiceClient, agentID string) (r TagsResult) {
	resp, err := client.Get(tagsURL(client, agentID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteTag deletes an individual tag from an agent.
func DeleteTag(client *gophercloud.ServiceClient, agentID string, key string) (r TagsErrResult) {
	resp, err := client.Delete(deleteTagURL(client, agentID, key), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetFacts lists tags for a given agent.
func GetFacts(client *gophercloud.ServiceClient, agentID string) (r FactsResult) {
	resp, err := client.Get(factsURL(client, agentID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
