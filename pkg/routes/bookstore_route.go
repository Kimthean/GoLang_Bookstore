package routes

import (
	"gihub.com/Thean-Arch/go_library/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
