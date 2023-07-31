package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssuloglu/simple-library-management-with-go/packages/routes"
)

func main() {
	r := mux.NewRouter()
	routes.LibraryRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8089", r))
}
