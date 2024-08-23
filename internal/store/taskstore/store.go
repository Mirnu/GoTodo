package taskstore

import (
	"time"
	"web1/internal/model"
)

type TaskStore interface {
	CreateTask(text string, tags []string, due time.Time) int
	GetTask(id int) (model.Task, error)
	DeleteTask(id int) error
	GetAllTasks() []model.Task
	GetTasksByTag(tag string) []model.Task
	GetTasksByDueDate(year int, month time.Month, day int) []model.Task
}

func NewTaskStore() TaskStore {
	return &taskStoreInMemory{
		tasks:  make(map[int]model.Task),
		nextId: 1,
	}
}
