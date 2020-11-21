package eventsrepository

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"wakelet-challenge/nasa"
)

func Create(events []nasa.Event) error {
	if len(events) == 0 {
		return errors.New("We need some events")
	}

	var dynamo *dynamodb.DynamoDB

	dynamo = connectDynamoDb()

	//add one nasa event
	av, err := dynamodbattribute.MarshalMap(events[0])

	if err != nil {
		fmt.Println("Got error marshalling event")
		fmt.Println(err.Error())
	}

	tableName := "events_table"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamo.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	return nil
}

func connectDynamoDb() (db *dynamodb.DynamoDB) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}))

	return dynamodb.New(sess)
}
