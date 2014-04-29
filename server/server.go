package main

import (
	"github.com/collinglass/bookAPI/server/ctrl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Server")

	r := mux.NewRouter()
	r.HandleFunc("/api/books/", ctrl.GetBookList()).Methods("GET")
	r.HandleFunc("/api/books/{id:[0-9]+}", ctrl.GetBook()).Methods("GET")

	http.Handle("/api/", r)
	http.Handle("/", logHandler(http.FileServer(http.Dir("../app/"))))

	log.Println("Listening...")
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func logHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
