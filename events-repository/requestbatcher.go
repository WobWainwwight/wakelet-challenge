package eventsrepository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"wakelet-challenge/nasa"
)

const (
	EventsTableName = "events_table"
)

func createEventBatches(events []nasa.Event, batches *[]dynamodb.BatchWriteItemInput) {
	var requests [][]*dynamodb.WriteRequest

	// Each batch can add 25 items so first we create slices of requests
	// of length 25 or less (depending on len(events))
	i := 0
	for {
		if (i + 24) < (len(events) - 1) {
			requests = append(requests, getRequestsFromEvents(events[i:i+25]))
		} else {
			requests = append(requests, getRequestsFromEvents(events[i:]))
			break
		}

		i += 25
	}

	for _, requestBatch := range requests {
		var m = map[string][]*dynamodb.WriteRequest{
			EventsTableName: requestBatch,
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
