package taskcontroller

import (
	"encoding/json"
	"net/http"
	"time"
	"web1/internal/store/taskstore"

	"github.com/go-chi/chi/v5"
)

func NewTaskController(store taskstore.TaskStore) *TaskController {
	return &TaskController{
		store: store,
	}
}

type TaskController struct {
	store taskstore.TaskStore
}

func (t *TaskController) Init(mux *chi.Mux) {
	mux.Post("/tasks", t.createTaskHandler)
	mux.Get("/tasks", t.getAllTasksHandler)
	mux.Get("/tasks/{id}", t.getTaskByIdHandler)
	mux.Get("/tasks/tag/{tag}", t.getTaskByTagHandler)
	mux.Get("/tasks/due/{due}", t.getTaskByDueHandler)
	mux.Delete("/tasks/{id}", t.deleteTaskHandler)
}

func (t *TaskController) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	type taskRequest struct {
		Text string   `json:"text"`
		Tags []string `json:"tags"`
		Due  string   `json:"due"`
	}

	var req taskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	layout := "2006-01-02 15:04:05"
	due, err := time.Parse(layout, req.Due)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	t.store.CreateTask(req.Text, req.Tags, due)

	responce := struct {
		Data string `json:"data"`
	}{
		Data: "ok",
	}

	renderJSON(w, responce)
}

func renderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (t *TaskController) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	renderJSON(w, t.store.GetAllTasks())
}

func (t *TaskController) getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {

}

func (t *TaskController) getTaskByTagHandler(w http.ResponseWriter, r *http.Request) {

}

func (t *TaskController) getTaskByDueHandler(w http.ResponseWriter, r *http.Request) {

}

func (t *TaskController) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}
