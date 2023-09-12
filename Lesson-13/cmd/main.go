package main

import (
	"net/http"
	"thinknetica/Lesson-13/pkg/storage"
	"thinknetica/Lesson-13/pkg/webapp"

	"github.com/gorilla/mux"
)

const depth = 2

var needle string

func main() {
	storage := storage.NewInMemoryStorage()
	api := webapp.New(storage)
	router := setupRoutes(api)
	http.ListenAndServe("localhost:8081", router)
}

func setupRoutes(api *webapp.API) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/add", api.Add)
	router.HandleFunc("/delete/{id}", api.Remove)
	router.HandleFunc("/show/{queryText}", api.FindByQueryText)
	router.HandleFunc("/update/{id}", api.UpdateById)
	return router
}
