package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
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

var (
	ErrorMethodNotAllowed string = "method not allowed"
	TableName             string = "tasks"
)

func main() {
	lambda.Start(handleRequest)
}
func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ddbClient := dynamodb.New(sess)

	log.Printf("Incomming request: %s", request.HTTPMethod)

	switch request.HTTPMethod {
	case "GET":
		return getTask(request, ddbClient)
	case "POST":
		return createTask(request, ddbClient)
	default:
		return response(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
	}
}

func getTask(request events.APIGatewayProxyRequest, ddbClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	id := request.QueryStringParameters["id"]
	title := request.QueryStringParameters["title"]

	log.Printf("Processing GET request with id: %s", id)

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

	log.Printf("Input: %v", input)

	result, err := ddbClient.GetItem(input)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	var task Task

	err = dynamodbattribute.UnmarshalMap(result.Item, &task)
	if err != nil {
		log.Fatalf("Failed to unmarshal event: %v", err)
	}
	return response(http.StatusOK, task)
}

func createTask(request events.APIGatewayProxyRequest, ddbClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	log.Print("Processing POST request")

	var task Task
	if err := json.Unmarshal([]byte(request.Body), &task); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return events.APIGatewayProxyResponse{Body: "Error: Failed to unmarshal event", StatusCode: 500}, err
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
		TableName: aws.String(TableName),
	}

	_, err = ddbClient.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	msg := "Created task: " + task.Title + ", with ID: " + task.Id
	log.Printf(msg)

	res, err := response(200, task)
	return res, err
}

func response(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	jsonBody, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: statusCode}, nil
}
