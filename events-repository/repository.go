package eventsrepository

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"wakelet-challenge/nasa"
)

var context *dynamodb.DynamoDB

func init() {
	context = connectDynamoDb()
}

func CreateMany(events []nasa.Event) error {
	if len(events) == 0 {
		return errors.New("We need some events")
	}

	//add one nasa event
	var batches []dynamodb.BatchWriteItemInput

	createEventBatches(events, &batches)

	for _, batchInput := range batches {
		for {
			errCount := 0
			_, err := context.BatchWriteItem(&batchInput)

			// Try five times or until we succesfully batch write
			if err == nil || errCount == 5 {
				break
			}
			errCount++
		}

	}

	return nil
}

func GetEvents(orderBy string) ([]nasa.Event, error) {
	var events []nasa.Event

	if orderBy == "title" {
		byTitleErr := orderByTitle(&events)
		if byTitleErr != nil {
			return nil, byTitleErr
		}
	} else if orderBy == "time" {
		return nil, errors.New("Not handled time yet")
	} else {
		return nil, errors.New("Can only order by time or name")
	}

	return events, nil
}

func orderByTitle(events *[]nasa.Event) error {
	keyCondition := expression.Key("id").Equal(expression.Value("nasa_event"))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		fmt.Print("expression builder error")
		return err
	}

	var limit int64 = 50
	input := dynamodb.QueryInput{
		TableName:                 aws.String(EventsTableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		Limit:                     &limit,
	}
	output, queryErr := context.Query(&input)

	if queryErr != nil {
		return queryErr
	}

	if len(output.Items) == 0 {
		return nil
	}

	for _, item := range output.Items {
		var event nasa.Event
		unmarshallErr := dynamodbattribute.UnmarshalMap(item, &event)

		if unmarshallErr == nil {
			*events = append(*events, event)
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
