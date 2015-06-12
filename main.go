package main

import (
	"log"
	"net/http"

	"github.com/jlope001/wasabi/bootstrap"
)

func main() {
	router := bootstrap.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
