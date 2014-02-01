package main

import (
	//"encoding/json"
	"github.com/gorilla/handlers"
	//"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Server")
	//initialize()
	//mainRouter := mux.NewRouter()
	//mainRouter.HandleFunc("/api/v0.1/message", MessageHandler).Methods("GET")

	//http.Handle("/api/", mainRouter)
	http.Handle("/", logHandler(http.FileServer(http.Dir("./app/"))))

	log.Println("Listening...")
	panic(http.ListenAndServe(":3000", nil))
}

func logHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
