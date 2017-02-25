//************************************************************************//
// API "cbbootstrap": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/couchbaselabs/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbaselabs/cbbootstrap/goa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// CreateOrJoinClusterContext provides the cluster create_or_join action context.
type CreateOrJoinClusterContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateOrJoinClusterPayload
}

// NewCreateOrJoinClusterContext parses the incoming request URL and body, performs validations and creates the
// context used by the cluster controller create_or_join action.
func NewCreateOrJoinClusterContext(ctx context.Context, service *goa.Service) (*CreateOrJoinClusterContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateOrJoinClusterContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateOrJoinClusterContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
