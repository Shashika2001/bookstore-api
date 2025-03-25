package routes

import (
	"github.com/gorilla/mux"
	"github.com/Shashika2001/bookstore-api/handlers"
)

// SetupRouter initializes the routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	return router
}
