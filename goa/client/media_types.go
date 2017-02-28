//************************************************************************//
// API "cbbootstrap": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/couchbaselabs/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbaselabs/cbbootstrap/goa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A CouchbaseCluster (default view)
//
// Identifier: application/vnd.couchbasecluster+json; view=default
type Couchbasecluster struct {
	// The cluster id
	ClusterID string `form:"cluster_id" json:"cluster_id" xml:"cluster_id"`
	// The initial node ip address or host that can be used to join cluster
	InitialNodeIPAddrOrHostname string `form:"initial_node_ip_addr_or_hostname" json:"initial_node_ip_addr_or_hostname" xml:"initial_node_ip_addr_or_hostname"`
}

// Validate validates the Couchbasecluster media type instance.
func (mt *Couchbasecluster) Validate() (err error) {
	if mt.ClusterID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "cluster_id"))
	}
	if mt.InitialNodeIPAddrOrHostname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "initial_node_ip_addr_or_hostname"))
	}

	return
}

// DecodeCouchbasecluster decodes the Couchbasecluster instance encoded in resp body.
func (c *Client) DecodeCouchbasecluster(resp *http.Response) (*Couchbasecluster, error) {
	var decoded Couchbasecluster
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
