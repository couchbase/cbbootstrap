A REST API intendended to be run as a public service to help bootstrap Couchbase distributed database clusters, in other words, to make it easy to stand up a Couchbase cluster from scratch by providing an API for Couchbase nodes to:

- Discover whether they are the first node in the cluster, and need to run cluster-init so other nodes can join them

or

- If the cluster is already initialized, discover the IP address / hostname of the node they should join.

## Deployment Architecture

![architecture](docs/aws-deployment-architecture.png)

## Sample Use

See [cbbootstrap.py](https://github.com/couchbaselabs/sg-autoscale/blob/master/src/cbbootstrap.py) -- this is intended to be called from a `user-data.sh` script on EC2 instance launch, but can be used in other scenarios.

## Backstory

Needed for the dynamic and automated [sg-autoscale](http://github.com/couchbaselabs/sg-autoscale) project to be able to start a CloudFormation and have the Couchbase Server cluster initialize itself based on parameters.

This could also be useful when trying to run Couchbase in container orchestration platforms like Kubernetes or Docker Swarm.

## REST API Definition

See [swagger.yaml](goa/swagger/swagger.yaml)

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
$ curl -X POST -d '{"cluster_id": "mycluster", "node_ip_addr_or_hostname": "mynode_ip"}' https://5e61vqxs5f.execute-api.us-east-1.amazonaws.com/Prod/cluster
```
