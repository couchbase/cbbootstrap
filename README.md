A REST API intendended to be run as a public service to help bootstrap Couchbase distributed database clusters, in other words, to make it easy to stand up a Couchbase cluster from scratch by providing an API for Couchbase nodes to:

- Discover whether they are the first node in the cluster, and need to run cluster-init so other nodes can join them

or

- If the cluster is already initialized, discover the IP address / hostname of the node they should join.

## Backstory

Needed for the highly dynamic and automated [sg-autoscale](http://github.com/couchbaselabs/sg-autoscale) project, has the requirement to be able to easily resize Couchbase server clusters with minimal human intervention.

This is also useful when trying to run Couchbase in container orchestration platforms like Kubernetes or Docker Swarm.

## REST API Definition

Endpoints

- POST /cluster/add_or_create

This creates or updates a cluster object in the system.  If it’s the first node in this cluster (defined by cluster_id), then 

Request

```
{
    “cluster_id”: “safdasdf3234",
    “node_ip_addr_or_hostname”: “ip-233.11.2.5"
}   
```

Response

```
{
    “cluster_already_initialized”: true | false,  // if false, then this node becomes the initial node that other nodes try to join
    “initial_node_ip_addr_or_hostname”: “ip-233.11.2.5”,  // empty if cluster_already_initialized
    "all_known_nodes": [
        {
		“node_ip_addr_or_hostname”: “ip-233.11.2.5",
		"last_seen": <date>
        },
        {
		“node_ip_addr_or_hostname”: “ip-233.11.2.18",
		"last_seen": <date>
        },
    ]
    
}
```

## Deploy to AWS Lambda

Package the lambda function 

```
$ wget -O Makefile https://github.com/eawsy/aws-lambda-go-shim/raw/master/src/Makefile.example
$ make
```

Deploy cloudformation with lambda function

```
$ aws cloudformation package \
  --template-file aws_serverless_application_model.yaml  \
  --output-template-file aws_serverless_application_model.out.yaml \
  --s3-bucket cf-templates-1m70kn8ou9eql-us-east-1
$ aws cloudformation deploy \
  --template-file aws_serverless_application_model.out.yaml \
  --capabilities CAPABILITY_IAM \
  --stack-name CBBootstrapExperiment \
  --region us-east-1
```

Get REST API endpoint

```
$ aws cloudformation describe-stacks \
  --stack-name CBBootstrapExperiment \
  --region us-east-1 | grep -i OutputValue

"OutputValue": "https://5e61vqxs5f.execute-api.us-east-1.amazonaws.com/Prod"
```

Test endpoint

```
$ curl https://5e61vqxs5f.execute-api.us-east-1.amazonaws.com/Prod
```


Call API

```
$ curl -X POST -d '{"cluster_id": "foo15", "node_ip_addr_or_hostname": "bar2"}' http://localhost:8080/cluster
```
