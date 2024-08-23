package taskstore

import (
	"fmt"
	"sync"
	"time"
	"web1/internal/model"
)

type taskStoreInMemory struct {
	m      sync.RWMutex
	tasks  map[int]model.Task
	nextId int
}

// CreateTask implements TaskStore.
func (t *taskStoreInMemory) CreateTask(text string, tags []string, due time.Time) int {
	t.m.Lock()
	defer t.m.Unlock()

	task := model.Task{
		Id:   t.nextId,
		Text: text,
		Tags: tags,
		Due:  due,
	}
	t.nextId++
	t.tasks[task.Id] = task

	return task.Id
}

func (t *taskStoreInMemory) DeleteTask(id int) error {
	t.m.Lock()
	defer t.m.Unlock()

	if _, ok := t.tasks[id]; !ok {
		return fmt.Errorf("task %d not found", id)
	}
	delete(t.tasks, id)
	return nil
}

// GetAllTasks implements TaskStore.
func (t *taskStoreInMemory) GetAllTasks() []model.Task {
	t.m.RLock()
	defer t.m.RUnlock()

	tasks := make([]model.Task, 0, len(t.tasks))

	for _, task := range t.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTask implements TaskStore.
func (t *taskStoreInMemory) GetTask(id int) (model.Task, error) {
	t.m.RLock()
	defer t.m.RUnlock()

	task, ok := t.tasks[id]
	if !ok {
		return model.Task{}, fmt.Errorf("task %d not found", id)
	}
	return task, nil
}

// GetTasksByDueDate implements TaskStore.
func (t *taskStoreInMemory) GetTasksByDueDate(year int, month time.Month, day int) []model.Task {
	t.m.RLock()
	defer t.m.RUnlock()

	tasks := make([]model.Task, 0, len(t.tasks))
	for _, task := range t.tasks {
		if task.Due.Year() == year && task.Due.Month() == month && task.Due.Day() == day {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTasksByTag implements TaskStore.
func (t *taskStoreInMemory) GetTasksByTag(tag string) []model.Task {
	t.m.RLock()
	defer t.m.RUnlock()

	tasks := make([]model.Task, 0, len(t.tasks))
	for _, task := range t.tasks {
		for _, t := range task.Tags {
			if t == tag {
				tasks = append(tasks, task)
			}
		}
	}
	return tasks
}
