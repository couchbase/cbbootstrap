
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
   
