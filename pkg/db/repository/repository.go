package repository

import "github.com/l4rma/todo-go/pkg/db/entity"

type TaskRepository interface {
	Save(task *entity.Task) (*entity.Task, error)
	FindById(id string) (*entity.Task, error)
	// FindAll() ([]*entity.Task, error)
	//UpdateById(id int64) (*entity.Task, error)
	// DeleteById(id int64) error
	// InsertDummyData(repo TaskRepository)
}
