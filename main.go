package main

// /* Required by eawsy/aws-lambda-go-net */
import "C"

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"fmt"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"net/http"
)

/*

   #import cb_bootstrap

   # Wrapper around bootstrap.couchbase.io REST API service which has global view of
   # cluster and can track which node is the boostrap

   #couchbase_cluster = cb_bootstrap.CouchbaseCluster(cluster_token, node_id)
   #couchbase_cluster.SetAdminUser("Administrator")
   #couchbase_cluster.SetAdminPassword("Password")
   #couchbase_cluster.SetCouchbaseServerName(socket.gethostname())  # how to get the public ip?
   #couchbase_cluster.WireUp()  # blocks until it either sets up as initial node or joins other nodes
   #couchbase_cluster.AddBucketIfMissing(
   #   Name="data-bucket",
   #   PercentRam=0.50,
   #)
   #couchbase_cluster.AddBucketIfMissing(
   #   Name="index-bucket",
   #   PercentRam=0.50,
   #)
 */


type CouchbaseCluster struct {
	ClusterId string // Something to uniquely identify the cluster
	DynamoDb dynamodbiface.DynamoDBAPI // DynamoDB driver or mock
}

type CouchbaseNode struct {
	CouchbaseCluster
	IpAddrOrHostname string // The ip address or hostname for this Couchbase Node
}

func NewCouchbaseNode(clusterId, iPAddrOrHostname string, dynamoDb dynamodbiface.DynamoDBAPI) *CouchbaseNode {
	return &CouchbaseNode{
		CouchbaseCluster: CouchbaseCluster{
			ClusterId: clusterId,
			DynamoDb: dynamoDb,
		},
		IpAddrOrHostname: iPAddrOrHostname,
	}

}

func (cbnode *CouchbaseNode) CreateOrJoinCuster() error {

	putItemInput := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(cbnode.ClusterId),
			},
		},
		TableName:           aws.String("cb-bootstrap"),
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}

	// Create a new cluster item, or retrieve existing
	clusterAlreadyExists := false
	putItemOutPut, err := cbnode.DynamoDb.PutItem(putItemInput)
	if err != nil {
		if err.Error() == dynamodb.ErrCodeConditionalCheckFailedException {
			log.Printf("Cluster already exists!  Err: %+v PutItemOutput: %+v", err, putItemOutPut)
			clusterAlreadyExists = true
		}
		// Unexpected error
		return err
	}

	if !clusterAlreadyExists {
		log.Printf("Cluster created: %+v", putItemOutPut)
	}

	return nil

}


func main() {

	log.Printf("hello world")

	// connect to dynamodb
	awsSession := session.New()
	dynamoDb := dynamodb.New(awsSession)

	// create a new CouchbaseNode
	cbNode := NewCouchbaseNode("foo3", "127.0.0.1", dynamoDb)
	err := cbNode.CreateOrJoinCuster()
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}

	log.Printf("done")

}


// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {
	ln := net.Listen()

	// Amazon API Gateway Binary support out of the box.
	Handle = apigatewayproxy.New(ln, []string{"image/png"}).Handle

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, etc.
	go http.Serve(ln, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

