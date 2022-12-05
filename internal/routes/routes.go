package routes

import (
	"github.com/gorilla/mux"
	"github.com/harsh/project/internal/service"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", service.GetUsers).Methods("GET")
	router.HandleFunc("/api/", service.CreateUser).Methods("POST")
	router.HandleFunc("/api/{id}", service.GetUser).Methods("GET")
	router.HandleFunc("/api/{id}", service.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/{id}", service.DeleteUser).Methods("DELETE")

	return router
}
