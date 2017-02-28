package controllers

import (

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


	cbcluster := cbcluster.CouchbaseCluster{
		ClusterId: *ctx.Payload.ClusterID,
		DynamoDb: dynamoDb,
	}

	cbNode, err := cbcluster.CreateOrJoinCuster(*ctx.Payload.NodeIPAddrOrHostname)
	if err != nil {
		ctx.ResponseData.WriteHeader(500)
		_, err2 := ctx.ResponseData.Write([]byte(err.Error()))
		return err2
	}

	cbClusterReturnVal := app.Couchbasecluster{
		ClusterID: cbNode.CouchbaseCluster.ClusterId,
		InitialNodeIPAddrOrHostname: cbNode.IpAddrOrHostname,
		IsInitialNode: cbNode.IsInitialNode,
	}

	return ctx.OK(&cbClusterReturnVal)

}
