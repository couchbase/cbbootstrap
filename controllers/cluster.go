package controllers

import (

	"github.com/couchbase/cbbootstrap/cbcluster"
	"github.com/couchbase/cbbootstrap/goa/app"
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


// Status runs the status action.
func (c *ClusterController) Status(ctx *app.StatusClusterContext) error {
	// ClusterController_Status: start_implement

	cbClusterReturnVal, err := getClusterStatus(ctx.ClusterID)
	if err != nil {
		ctx.ResponseData.WriteHeader(500)
		_, err2 := ctx.ResponseData.Write([]byte(err.Error()))
		return err2
	}

	return ctx.OK(&cbClusterReturnVal)

}

// GetStatus runs the get_status action.
func (c *ClusterController) GetStatus(ctx *app.GetStatusClusterContext) error {
	// ClusterController_GetStatus: start_implement

	cbClusterReturnVal, err := getClusterStatus(ctx.Payload.ClusterID)
	if err != nil {
		ctx.ResponseData.WriteHeader(500)
		_, err2 := ctx.ResponseData.Write([]byte(err.Error()))
		return err2
	}

	return ctx.OK(&cbClusterReturnVal)
}

func getClusterStatus(ClusterId string) (app.Couchbasecluster, error) {

	dynamoDb := cbcluster.CreateDynamoDbSession()

	cbcluster := cbcluster.CouchbaseCluster{
		ClusterId: ClusterId,
		DynamoDb: dynamoDb,
	}

	// Load existing
	cbNode := cbcluster.NewCouchbaseNode()

	// Load from DB
	err := cbNode.DBLoad()
	if err != nil {
		return app.Couchbasecluster{}, err
	}

	cbNode.IsInitialNode = true

	cbClusterReturnVal := app.Couchbasecluster{
		ClusterID: cbNode.CouchbaseCluster.ClusterId,
		InitialNodeIPAddrOrHostname: cbNode.IpAddrOrHostname,
		IsInitialNode: cbNode.IsInitialNode,
	}

	return cbClusterReturnVal, nil

}
