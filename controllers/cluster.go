package controllers

import (
	"fmt"

	"github.com/couchbaselabs/cbbootstrap/cbcluster"
	"github.com/couchbaselabs/cbbootstrap/goa/app"
	"github.com/goadesign/goa"
)

// ClusterController implements the cluster resource.
type ClusterController struct {
	*goa.Controller
}

// NewClusterController creates a cluster controller.
func NewClusterController(service *goa.Service) *ClusterController {
	return &ClusterController{Controller: service.NewController("ClusterController")}
}

// CreateOrJoin runs the create_or_join action.
func (c *ClusterController) CreateOrJoin(ctx *app.CreateOrJoinClusterContext) error {
	// ClusterController_CreateOrJoin: start_implement

	dynamoDb := cbcluster.CreateDynamoDbSession()

	// create a new CouchbaseNode
	cbNode := cbcluster.NewCouchbaseNode(
		*ctx.Payload.ClusterID,
		*ctx.Payload.NodeIPAddrOrHostname,
		dynamoDb,
	)

	err := cbNode.CreateOrJoinCuster()
	if err != nil {
		return ctx.OK([]byte(err.Error()))
	}

	return ctx.OK([]byte(fmt.Sprintf("Got cbNode: %+v", cbNode)))

}
