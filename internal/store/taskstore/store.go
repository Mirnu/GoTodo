package taskstore

import (
	"time"
	"web1/internal/model"
)

type TaskStore interface {
	CreateTask(text string, tags []string, due time.Time) int
	GetTask(id int) (model.Task, error)
	DeleteTask(id int) error
	DeleteAllTasks() error
	GetAllTasks() []model.Task
	GetTasksByTag(tag string) []model.Task
	GetTasksByDueDate(year int, month time.Month, day int) []model.Task
}
