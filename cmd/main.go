package main

import (
	"log"
	"net/http"
	"web1/internal/controller/taskcontroller"
	"web1/internal/store/taskstore"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	taskstore := taskstore.NewTaskStore()
	taskcontroller := taskcontroller.NewTaskController(taskstore)
	taskcontroller.Init(router)

	log.Fatal(http.ListenAndServe(":80", router))
}
