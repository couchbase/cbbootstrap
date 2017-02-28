//************************************************************************//
// API "cbbootstrap": cluster TestHelpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/couchbaselabs/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbaselabs/cbbootstrap/goa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package test

import (
	"bytes"
	"fmt"
	"github.com/couchbaselabs/cbbootstrap/goa/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// CreateOrJoinClusterOK runs the method CreateOrJoin of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateOrJoinClusterOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ClusterController, payload *app.CreateOrJoinClusterPayload) (http.ResponseWriter, *app.Couchbasecluster) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cluster"),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ClusterTest"), rw, req, prms)
	createOrJoinCtx, err := app.NewCreateOrJoinClusterContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createOrJoinCtx.Payload = payload

	// Perform action
	err = ctrl.CreateOrJoin(createOrJoinCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Couchbasecluster
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Couchbasecluster)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Couchbasecluster", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
