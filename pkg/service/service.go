package service

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/l4rma/todo-go/internal/db/entity"
	"github.com/l4rma/todo-go/internal/db/repository"
)

var (
	taskRepo repository.TaskRepository
)

type TaskService interface {
	// Validate(book *entity.Task) error
	Create(book *entity.Task) (*entity.Task, error)
	// FindAll() ([]*entity.Task, error)
	FindbyId(id string, title string) (*entity.Task, error)
	// Delete(id int64) error
}

type service struct{}

func NewTaskService(repo repository.TaskRepository) TaskService {
	taskRepo = repo
	return &service{}
}

func (*service) FindbyId(id string, title string) (*entity.Task, error) {
	return taskRepo.FindById(id, title)
}

func (*service) Create(task *entity.Task) (*entity.Task, error) {
	return taskRepo.Save(task)
}

func response(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	jsonBody, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: statusCode}, nil
}
