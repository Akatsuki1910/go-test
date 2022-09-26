package main

import (
	handler "go-test/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
