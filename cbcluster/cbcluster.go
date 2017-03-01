package cbcluster

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

type CouchbaseCluster struct {
	ClusterId string                    // Something to uniquely identify the cluster
	DynamoDb  dynamodbiface.DynamoDBAPI // DynamoDB driver or mock
}


func (cluster *CouchbaseCluster) CreateOrJoinCuster(iPAddrOrHostname string) (CouchbaseNode, error) {

	// Create a new cluster object from database, or retrieve existing
	// -------------------------------------------------------------------------------------------------------------
	err := cluster.DBCreate(iPAddrOrHostname)
	if err == nil {

		// Create succeeded -- which means iPAddrOrHostname successfully became the inital node

		cbNode := cluster.NewCouchbaseNode()

		// no error,
		cbNode.IsInitialNode = true
		cbNode.IpAddrOrHostname = iPAddrOrHostname

		return cbNode, nil
	}

	// Create failed -- if due to existing cluster, then fetch existing cluster details, or else raise error
	// -------------------------------------------------------------------------------------------------------------

	// We got an error.  If it was just a ErrCodeConditionalCheckFailedException,
	// then we should just do a GetItem call to get the value
	awsErr, ok := err.(awserr.Error)
	if !ok {
		// unexpected error
		log.Printf("Expected an awserr.Error, got: %+v", err)
		return CouchbaseNode{}, err
	}

	if awsErr.Code() != dynamodb.ErrCodeConditionalCheckFailedException {
		// unexpected error
		log.Printf("Expected an awserr.Error with dynamodb.ErrCodeConditionalCheckFailedException, got :%+v", awsErr)
		return CouchbaseNode{}, err
	}

	log.Printf("Cluster already exists!  Fetching existing initial node from db.")

	// now we need to do a fetch to get the initial node ip addr or host
	cbNode := cluster.NewCouchbaseNode()

	err = cbNode.DBLoad()
	if err != nil {
		return CouchbaseNode{}, err
	}

	log.Printf("Loaded cbnode from db: %+v", cbNode)

	cbNode.IsInitialNode = false

	return cbNode, nil

}


func (cluster *CouchbaseCluster) DBCreate(iPAddrOrHostname string) error {

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
	putItemOutPut, err := cluster.DynamoDb.PutItem(putItemInput)
	log.Printf("Cluster.DBCreate() PutItemOutput: %+v.  err: %v", putItemOutPut, err)
	return err

}

func (cluster *CouchbaseCluster) NewCouchbaseNode() CouchbaseNode {

	return CouchbaseNode{
		CouchbaseCluster: cluster,
	}

}

