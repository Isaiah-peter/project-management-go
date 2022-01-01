package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port != "" {
		port = "5000"
	}

	router := mux.NewRouter()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(port, router))
}
