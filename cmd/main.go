package main

import (
	"log"
	"net/http"
	"web1/internal/controller/taskcontroller"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	taskcontroller := taskcontroller.NewTaskController()
	taskcontroller.Init(router)

	log.Fatal(http.ListenAndServe(":80", router))
}
