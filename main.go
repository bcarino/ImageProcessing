package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	router := NewRouter()

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
