package main

// /* Required by eawsy/aws-lambda-go-net */
import "C"

import (
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/couchbaselabs/cbbootstrap/goa/app"
	"github.com/couchbaselabs/cbbootstrap/controllers"
)

func main() {

	/*log.Printf("hello world")

	dynamoDb := cbcluster.CreateDynamoDbSession()

	// create a new CouchbaseNode
	cbNode := cbcluster.NewCouchbaseNode("foo3", "127.0.0.1", dynamoDb)
	err := cbNode.CreateOrJoinCuster()
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}

	log.Printf("done")*/

	// Create service
	service := goa.New("CBBootstrap REST API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "cluster" controller
	c := controllers.NewClusterController(service)
	app.MountClusterController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

//func init() {
//	ln := net.Listen()
//
//	// Amazon API Gateway Binary support out of the box.
//	Handle = apigatewayproxy.New(ln, []string{"image/png"}).Handle
//
//	// Any Go framework complying with the Go http.Handler interface can be used.
//	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, etc.
//	go http.Serve(ln, http.HandlerFunc(handle))
//}

//func handle(w http.ResponseWriter, r *http.Request) {
//
//	dynamoDb := cbcluster.CreateDynamoDbSession()
//
//	// create a new CouchbaseNode
//	cbNode := cbcluster.NewCouchbaseNode("foo4", "127.0.0.1", dynamoDb)
//	err := cbNode.CreateOrJoinCuster()
//	if err != nil {
//		w.Write([]byte(err.Error()))
//	}
//
//	w.Write([]byte(fmt.Sprintf("Got cbNode: %+v", cbNode)))
//
//}
