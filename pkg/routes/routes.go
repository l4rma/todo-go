package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/l4rma/todo-go/pkg/db/entity"
	"github.com/l4rma/todo-go/pkg/db/repository"
	"github.com/l4rma/todo-go/pkg/service"
)

var (
	ErrorMethodNotAllowed string                    = "method not allowed"
	taskRepository        repository.TaskRepository = repository.NewDynamoDBRepository()
	taskService           service.TaskService       = service.NewTaskService(taskRepository)
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Incomming %s request", request.HTTPMethod)

	switch request.HTTPMethod {
	case "GET":
		id := request.QueryStringParameters["id"]

		task, err := taskService.FindbyId(id)
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
