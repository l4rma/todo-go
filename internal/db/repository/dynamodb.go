package repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/l4rma/todo-go/internal/db/entity"
)

var (
	db        *dynamodb.DynamoDB
	TableName string = "tasks"
)

type dynamodbRepository struct{}

func NewDynamoDBRepository() TaskRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ddbClient := dynamodb.New(sess)
	db = ddbClient

	return &dynamodbRepository{}
}

func (*dynamodbRepository) Save(task *entity.Task) (*entity.Task, error) {
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
		TableName: aws.String(TableName),
	}

	_, err = db.PutItem(input)

	msg := "Task saved in DB. Title: " + task.Title + ", ID: " + task.Id
	log.Printf(msg)

	return task, err
}

func (*dynamodbRepository) FindById(id string, title string) (*entity.Task, error) {
	task := &entity.Task{}

	log.Printf("Getting item from table: '%s' with id: '%s' and title: '%s'", TableName, id, title)

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
			"title": {
				S: aws.String(title),
			},
		},
		TableName: aws.String(TableName),
	}

	result, err := db.GetItem(input)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &task)
	if err != nil {
		log.Fatalf("Failed to unmarshal event: %v", err)
	}
	return task, err
}
