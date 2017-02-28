package cbcluster

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
	ClusterId string                    // Something to uniquely identify the cluster
	DynamoDb  dynamodbiface.DynamoDBAPI // DynamoDB driver or mock
}

type CouchbaseNode struct {
	CouchbaseCluster *CouchbaseCluster
	IpAddrOrHostname string // The ip address or hostname for this Couchbase Node
	InitialNode      bool   // Whether this is the initial node that others can join
}

func (cluster *CouchbaseCluster) CreateOrJoinCuster(iPAddrOrHostname string) (CouchbaseNode, error) {

	cbNode := CouchbaseNode{
		CouchbaseCluster: cluster,
	}

	putItemInput := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"cluster_id": {
				S: aws.String(cluster.ClusterId),
			},
			"initial_node_ip_addr_or_hostname": {
				S: aws.String(iPAddrOrHostname),
			},
		},
		TableName:           aws.String("cb-bootstrap"),
		ConditionExpression: aws.String("attribute_not_exists(cluster_id)"),
	}

	// Create a new cluster item, or retrieve existing
	putItemOutPut, err := cluster.DynamoDb.PutItem(putItemInput)
	if err != nil {

		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				log.Printf("Cluster already exists!  Err: %+v PutItemOutput: %+v", err, putItemOutPut)

				// now we need to do a fetch to get the initial node ip addr or host
				err2 := cbNode.LoadFromDatabase()
				if err2 != nil {
					return cbNode, err2
				}

				log.Printf("Loaded cbnode from db: %+v", cbNode)

				cbNode.InitialNode = false

				return cbNode, nil
			} else {
				// unexpected error
				log.Printf("Unexpected errort: %v", err)
				return cbNode, err
			}
		} else {
			// unexpected error
			log.Printf("Unexpected errort: %v", err)
			return cbNode, err
		}


		// unexpected error
		log.Printf("Unexpected errort: %v", err)
		return cbNode, err

	}

	// if we got this far, then we successfully became the inital node
	cbNode.InitialNode = true
	cbNode.IpAddrOrHostname = iPAddrOrHostname

	return cbNode, nil


}


func (cbnode *CouchbaseNode) LoadFromDatabase() error {

	attribute := dynamodb.AttributeValue{S: aws.String(cbnode.CouchbaseCluster.ClusterId)}
	query := map[string]*dynamodb.AttributeValue{"cluster_id": &attribute}

	getItemInput := &dynamodb.GetItemInput{
		Key: query,
		ConsistentRead: aws.Bool(true),
		TableName:           aws.String("cb-bootstrap"),
	}
	getItemOutput, err := cbnode.CouchbaseCluster.DynamoDb.GetItem(getItemInput)
	if err != nil {
		return err
	}

	initialNodeIpOrHostnameAttribute := getItemOutput.Item["initial_node_ip_addr_or_hostname"]
	cbnode.IpAddrOrHostname = *initialNodeIpOrHostnameAttribute.S

	return nil

}


/*
func NewCouchbaseNode(clusterId, iPAddrOrHostname string, dynamoDb dynamodbiface.DynamoDBAPI) *CouchbaseNode {
	return &CouchbaseNode{
		CouchbaseCluster: CouchbaseCluster{
			ClusterId: clusterId,
			DynamoDb:  dynamoDb,
		},
		IpAddrOrHostname: iPAddrOrHostname,
	}

}*/

/*
func (cbnode *CouchbaseNode) CreateOrJoinCuster() error {

	putItemInput := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"cluster_id": {
				S: aws.String(cbnode.ClusterId),
			},
			"node_ip_addr_or_hostname": {
				S: aws.String(cbnode.IpAddrOrHostname),
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
		cbnode.LoadFromDB()
		// Unexpected error
		return err
	}

	if !clusterAlreadyExists {
		log.Printf("Cluster created: %+v", putItemOutPut)
	}

	return nil

}
*/

func CreateDynamoDbSession() *dynamodb.DynamoDB {
	// connect to dynamodb
	awsSession := session.New()
	dynamoDb := dynamodb.New(awsSession)
	return dynamoDb
}
