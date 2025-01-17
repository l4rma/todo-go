package service

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/l4rma/todo-go/pkg/db/entity"
	"github.com/l4rma/todo-go/pkg/db/repository"
)

var (
	taskRepo repository.TaskRepository
)

type TaskService interface {
	// Validate(book *entity.Task) error
	Create(task *entity.Task) (*entity.Task, error)
	FindAll() ([]*entity.Task, error)
	FindById(id string) (*entity.Task, error)
	UpdateTask(task *entity.Task) (*entity.Task, error)
	// Delete(id int64) error
}

type service struct{}

func NewTaskService(repo repository.TaskRepository) TaskService {
	taskRepo = repo
	return &service{}
}

func (*service) FindAll() ([]*entity.Task, error) {
	return taskRepo.FindAll()
}

func (*service) FindById(id string) (*entity.Task, error) {
	return taskRepo.FindById(id)
}

func (*service) Create(task *entity.Task) (*entity.Task, error) {
	return taskRepo.Save(task)
}

func (*service) UpdateTask(task *entity.Task) (*entity.Task, error) {
	return taskRepo.UpdateById(task)
}

func response(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	jsonBody, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: statusCode}, nil
}
