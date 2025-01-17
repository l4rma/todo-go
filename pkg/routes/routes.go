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
	ErrorMethodNotAllowed string = "method not allowed"
	//taskRepository        repository.TaskRepository = repository.NewDynamoDBRepository()
	taskRepository repository.TaskRepository = repository.NewInMemoryRepository()
	taskService    service.TaskService       = service.NewTaskService(taskRepository)
	handler        Handler                   = Handler{}
)

func HandleRequest() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handler.FindAll)
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /task", handler.FindById)
	mux.HandleFunc("PATCH /task", handler.UpdateTask)

	log.Printf("Server started on port 8080")
	http.ListenAndServe(":8080", mux)
}

func OldHandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Incomming %s request", request.HTTPMethod)

	switch request.HTTPMethod {
	case "GET":
		id := request.QueryStringParameters["id"]

		if id == "" {
			tasks, err := taskService.FindAll()
			if err != nil {
				log.Printf("Error: %v", err)
			}

			return response(200, tasks)
		}

		task, err := taskService.FindById(id)
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
