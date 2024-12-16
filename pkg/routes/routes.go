package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/l4rma/todo-go/internal/db/entity"
	"github.com/l4rma/todo-go/internal/service"
)

var (
	ErrorMethodNotAllowed string = "method not allowed"
	taskService           service.TaskService
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Incomming %s request", request.HTTPMethod)

	switch request.HTTPMethod {
	case "GET":
		id := request.QueryStringParameters["id"]
		title := request.QueryStringParameters["title"]

		task, err := taskService.FindbyId(id, title)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		return response(200, task)
	case "POST":
		task := &entity.Task{}
		if err := json.Unmarshal([]byte(request.Body), &task); err != nil {
			log.Printf("Failed to unmarshal event: %v", err)
			return events.APIGatewayProxyResponse{Body: "Error: Failed to unmarshal event", StatusCode: 500}, err
		}

		task, err := taskService.Create(task)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		return response(200, task)
	default:
		return response(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
	}
}

func response(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	jsonBody, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: statusCode}, nil
}
