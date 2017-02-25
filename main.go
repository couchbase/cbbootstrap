package main

// /* Required by eawsy/aws-lambda-go-net */
import "C"

import (

	"log"
	"fmt"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"net/http"
	"github.com/couchbaselabs/cbbootstrap/cbcluster"
	"github.com/goadesign/goa"
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

	service := goa.New("CBBootstrap REST API")

	accountController := controllers.NewClusterController(service, coreService)
	app.MountAccountController(service, accountController)

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

