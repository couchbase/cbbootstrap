package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateOrJoinClusterPath computes a request path to the create_or_join action of cluster.
func CreateOrJoinClusterPath() string {
	return fmt.Sprintf("/cluster")
}

// Create a new Couchbase Cluster
func (c *Client) CreateOrJoinCluster(ctx context.Context, path string, payload *CreateOrJoinClusterPayload) (*http.Response, error) {
	req, err := c.NewCreateOrJoinClusterRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateOrJoinClusterRequest create the request corresponding to the create_or_join action endpoint of the cluster resource.
func (c *Client) NewCreateOrJoinClusterRequest(ctx context.Context, path string, payload *CreateOrJoinClusterPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
