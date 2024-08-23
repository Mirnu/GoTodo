package taskcontroller

import "github.com/go-chi/chi/v5"

func NewTaskController() *TaskController {
	return &TaskController{}
}

type TaskController struct{}

func (t *TaskController) Init(mux *chi.Mux) {

}
