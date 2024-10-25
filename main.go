package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event json.RawMessage) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	tableName := "tasks"

	var task Task
	if err := json.Unmarshal(event, &task); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return err
	}

	// If no id, generate random id
	if task.Id == "" {
		task.Id = uuid.New().String()
		log.Printf("Ready to put item with new id: %s", task.Id)
	}

	taskAttributeValue, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		log.Fatalf("Get error marshalling new task item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      taskAttributeValue,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	log.Printf("Created task '%s' with ID: %s", task.Title, task.Id)

	return nil
}
