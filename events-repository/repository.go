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

const eventsTable = "events_table"

func CreateMany(events []nasa.Event) error {
	if len(events) == 0 {
		return errors.New("We need some events")
	}

	var dynamo *dynamodb.DynamoDB

	dynamo = connectDynamoDb()

	//add one nasa event
	var batches []dynamodb.BatchWriteItemInput

	createEventBatches(events, &batches)

	for _, batchInput := range batches {
		for {
			errCount := 0
			_, err := dynamo.BatchWriteItem(&batchInput)

			// Try five times or until we succesfully batch write
			if err == nil || errCount == 5 {
				break
			}
			errCount++
		}

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

func createEventBatches(events []nasa.Event, batches *[]dynamodb.BatchWriteItemInput) {
	var requests [][]*dynamodb.WriteRequest

	// Each batch can add 25 items so first we create slices of requests
	// of length 25 or less (depending on len(events))
	i := 0
	for {
		if (i + 24) < (len(events) - 1) {
			requests = append(requests, getRequestsFromEvents(events[i:i+24]))
		} else {
			requests = append(requests, getRequestsFromEvents(events[i:]))
			break
		}

		i += 25
	}

	for _, requestBatch := range requests {
		var m = map[string][]*dynamodb.WriteRequest{
			eventsTable: requestBatch,
		}
		var batch = dynamodb.BatchWriteItemInput{
			RequestItems: m,
		}
		*batches = append(*batches, batch)
	}

}

func getRequestsFromEvents(events []nasa.Event) []*dynamodb.WriteRequest {
	var requests []dynamodb.WriteRequest
	var requestPointers []*dynamodb.WriteRequest

	for i, event := range events {
		av, err := dynamodbattribute.MarshalMap(event)

		if err != nil {
			fmt.Print("Error in marshalling event %b", i)
			continue
		}

		putRequest := dynamodb.PutRequest{
			Item: av,
		}

		request := dynamodb.WriteRequest{
			PutRequest: &putRequest,
		}

		requests = append(requests, request)
		requestPointers = append(requestPointers, &request)

	}
	fmt.Printf("Added %v", len(requests))

	return requestPointers
}
