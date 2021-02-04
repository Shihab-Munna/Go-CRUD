package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/users/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users", middleware.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/users/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/users/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
