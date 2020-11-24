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

	for i, batchInput := range batches {
		fmt.Println(i)
		errCount := 0
		for {
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
	var query dynamodb.QueryInput

	if orderBy == "title" {
		byTitleErr := orderByTitle(&query)
		if byTitleErr != nil {
			return nil, byTitleErr
		}
	} else if orderBy == "time" {
		byTimeErr := orderByTime(&query)
		if byTimeErr != nil {
			return nil, byTimeErr
		}
	} else {
		return nil, errors.New("Can only order by time or name")
	}

	output, queryErr := context.Query(&query)

	if queryErr != nil {
		return nil, queryErr
	}
	if len(output.Items) == 0 {
		return nil, nil
	}
	for _, item := range output.Items {
		var event nasa.Event
		unmarshallErr := dynamodbattribute.UnmarshalMap(item, &event)

		if unmarshallErr == nil {
			events = append(events, event)
		}
	}
	return events, nil
}

func orderByTitle(query *dynamodb.QueryInput) error {
	keyCondition := expression.Key("id").Equal(expression.Value("nasa_event"))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		fmt.Print("expression builder error")
		return err
	}

	var limit int64 = 50

	*query = dynamodb.QueryInput{
		TableName:                 aws.String(EventsTableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		Limit:                     &limit,
	}

	return nil

}

func orderByTime(query *dynamodb.QueryInput) error {
	// Use global secondary indexes
	keyCondition := expression.Key("id").Equal(expression.Value("nasa_event"))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		fmt.Print("expression builder error")
		return err
	}

	var limit int64 = 50
	var scanForward bool = false
	indexName := "id-time-index"

	*query = dynamodb.QueryInput{
		TableName:                 aws.String(EventsTableName),
		IndexName:                 aws.String(indexName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		Limit:                     &limit,
		ScanIndexForward:          &scanForward,
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
