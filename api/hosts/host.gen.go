// Code generated by "make api"; DO NOT EDIT.
package hosts

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type Host struct {
	Id                string                 `json:"id,omitempty"`
	HostCatalogId     string                 `json:"host_catalog_id,omitempty"`
	Scope             *scopes.ScopeInfo      `json:"scope,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Description       string                 `json:"description,omitempty"`
	CreatedTime       time.Time              `json:"created_time,omitempty"`
	UpdatedTime       time.Time              `json:"updated_time,omitempty"`
	Version           uint32                 `json:"version,omitempty"`
	Type              string                 `json:"type,omitempty"`
	HostSetIds        []string               `json:"host_set_ids,omitempty"`
	Attributes        map[string]interface{} `json:"attributes,omitempty"`
	AuthorizedActions []string               `json:"authorized_actions,omitempty"`

	response *api.Response
}

type HostReadResult struct {
	Item     *Host
	response *api.Response
}

func (n HostReadResult) GetItem() interface{} {
	return n.Item
}

func (n HostReadResult) GetResponse() *api.Response {
	return n.response
}

type (
	HostCreateResult = HostReadResult
	HostUpdateResult = HostReadResult
)

type HostDeleteResult struct {
	response *api.Response
}

// GetItem will always be nil for HostDeleteResult
func (n HostDeleteResult) GetItem() interface{} {
	return nil
}

func (n HostDeleteResult) GetResponse() *api.Response {
	return n.response
}

type HostListResult struct {
	Items    []*Host
	response *api.Response
}

func (n HostListResult) GetItems() interface{} {
	return n.Items
}

func (n HostListResult) GetResponse() *api.Response {
	return n.response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, hostCatalogId string, opt ...Option) (*HostCreateResult, error) {
	if hostCatalogId == "" {
		return nil, fmt.Errorf("empty hostCatalogId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["host_catalog_id"] = hostCatalogId

	req, err := c.client.NewRequest(ctx, "POST", "hosts", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(HostCreateResult)
	target.Item = new(Host)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, id string, opt ...Option) (*HostReadResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("hosts/%s", id), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(HostReadResult)
	target.Item = new(Host)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, id string, version uint32, opt ...Option) (*HostUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("hosts/%s", id), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(HostUpdateResult)
	target.Item = new(Host)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, id string, opt ...Option) (*HostDeleteResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("hosts/%s", id), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &HostDeleteResult{
		response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, hostCatalogId string, opt ...Option) (*HostListResult, error) {
	if hostCatalogId == "" {
		return nil, fmt.Errorf("empty hostCatalogId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["host_catalog_id"] = hostCatalogId

	req, err := c.client.NewRequest(ctx, "GET", "hosts", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(HostListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}
