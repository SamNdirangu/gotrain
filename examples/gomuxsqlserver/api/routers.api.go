package api

import (
	"training/examples/gomuxsqlserver/api/controllers"

	"github.com/gorilla/mux"
)

var BooksAPIRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	//Creating and Updating
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
