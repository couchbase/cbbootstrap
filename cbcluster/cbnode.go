package cbcluster

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

type CouchbaseNode struct {
	CouchbaseCluster *CouchbaseCluster
	IpAddrOrHostname string // The ip address or hostname for this Couchbase Node
	InitialNode      bool   // Whether this is the initial node that others can join
}


func (cbnode *CouchbaseNode) DBLoad() error {

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
