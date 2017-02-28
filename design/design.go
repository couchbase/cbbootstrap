package design

import (
. "github.com/goadesign/goa/design"
. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cbbootstrap", func() {
	Title("REST API to enable bootstrapping Couchbase Server clusters")
	Description("REST API to enable bootstrapping Couchbase Server clusters")
	Host("localhost:8080")
	Scheme("http", "https")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("cluster", func() {

	BasePath("/cluster") // Gets appended to the API base path
	DefaultMedia(CouchbaseClusterJson)

	Action("create_or_join", func() {

		Routing(POST(""))
		Description("Create a new Couchbase Cluster")
		Payload(CreateOrJoinClusterPayload, func() {
			Required("cluster_id", "node_ip_addr_or_hostname")
		})
		Response(OK)
	})


})

var CreateOrJoinClusterPayload = Type("CreateOrJoinClusterPayload", func() {
	Attribute("cluster_id", func() {
		MinLength(1)
	})
	Attribute("node_ip_addr_or_hostname", func() {
		MinLength(1)
	})
})

var CouchbaseClusterJson = MediaType("application/vnd.couchbasecluster+json", func() {
	Description("A CouchbaseCluster")
	Attributes(func() {
		Attribute("cluster_id", String, "The cluster id", func() {
			Example("FooAWSStack123")
		})
		Attribute("initial_node_ip_addr_or_hostname", String, "The initial node ip address or host that can be used to join cluster", func() {
			Example("10.1.1.1")
		})

		Required("cluster_id", "initial_node_ip_addr_or_hostname")
	})

	View("default", func() {
		Attribute("cluster_id")
		Attribute("initial_node_ip_addr_or_hostname")
	})

})