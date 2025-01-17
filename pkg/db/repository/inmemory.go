package repository

import "github.com/l4rma/todo-go/pkg/db/entity"

type inMemoryRepository struct {
	db []entity.Task
}

func NewInMemoryRepository() TaskRepository {
	return &inMemoryRepository{
		db: []entity.Task{},
	}
}

func (repo *inMemoryRepository) Save(task *entity.Task) (*entity.Task, error) {
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
