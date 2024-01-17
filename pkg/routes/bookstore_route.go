package routes

import (
	"github.com/gorilla/mux"

	"github.com/thean/go_library/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
