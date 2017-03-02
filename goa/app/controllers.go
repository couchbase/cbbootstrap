// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/couchbase/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbase/cbbootstrap/goa
// --version=v1.0.0
//
// API "cbbootstrap": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ClusterController is the controller interface for the Cluster actions.
type ClusterController interface {
	goa.Muxer
	CreateOrJoin(*CreateOrJoinClusterContext) error
	Status(*StatusClusterContext) error
}

// MountClusterController "mounts" a Cluster resource controller on the given service.
func MountClusterController(service *goa.Service, ctrl ClusterController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateOrJoinClusterContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateOrJoinClusterPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateOrJoin(rctx)
	}
	service.Mux.Handle("POST", "/cluster", ctrl.MuxHandler("CreateOrJoin", h, unmarshalCreateOrJoinClusterPayload))
	service.LogInfo("mount", "ctrl", "Cluster", "action", "CreateOrJoin", "route", "POST /cluster")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStatusClusterContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Status(rctx)
	}
	service.Mux.Handle("GET", "/cluster/:cluster_id", ctrl.MuxHandler("Status", h, nil))
	service.LogInfo("mount", "ctrl", "Cluster", "action", "Status", "route", "GET /cluster/:cluster_id")
}

// unmarshalCreateOrJoinClusterPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateOrJoinClusterPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createOrJoinClusterPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
