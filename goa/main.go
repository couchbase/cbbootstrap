//go:generate goagen bootstrap -d github.com/couchbaselabs/cbbootstrap/design

package main

import (
	"github.com/couchbaselabs/cbbootstrap/goa/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("cbbootstrap")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "cluster" controller
	c := NewClusterController(service)
	app.MountClusterController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
