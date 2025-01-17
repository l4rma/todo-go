package repository

import (
	"strconv"

	"github.com/l4rma/todo-go/pkg/db/entity"
)

type inMemoryRepository struct {
	db []entity.Task
	id int
}

func NewInMemoryRepository() TaskRepository {
	return &inMemoryRepository{
		db: []entity.Task{},
		id: 0,
	}
}

func (repo *inMemoryRepository) Save(task *entity.Task) (*entity.Task, error) {
	repo.id++
	task.Id = string(strconv.Itoa(repo.id))
	repo.db = append(repo.db, *task)
	return task, nil
}

func (repo *inMemoryRepository) FindById(id string) (*entity.Task, error) {
	for _, task := range repo.db {
		if task.Id == id {
			return &task, nil
		}
	}
	return nil, nil
}

func (repo *inMemoryRepository) FindAll() ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)
	for _, task := range repo.db {
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (repo *inMemoryRepository) UpdateById(task *entity.Task) (*entity.Task, error) {
	for i, t := range repo.db {
		if t.Id == task.Id {
			repo.db[i] = *task
			return task, nil
		}
	}
	return nil, nil
}
