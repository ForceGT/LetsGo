package routes

import (
	"github.com/ForceGT/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
