package cbcluster

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateDynamoDbSession() *dynamodb.DynamoDB {
	// connect to dynamodb
	awsSession := session.New()
	dynamoDb := dynamodb.New(awsSession)
	return dynamoDb
}
