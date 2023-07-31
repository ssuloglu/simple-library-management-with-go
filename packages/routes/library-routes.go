package routes

import (
	"github.com/gorilla/mux"
	"github.com/ssuloglu/simple-library-management-with-go/packages/controllers"
)

var LibraryRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/author/{author}", controllers.GetBooksByAuthor).Methods("GET")
}
