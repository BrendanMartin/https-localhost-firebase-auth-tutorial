package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServeTLS(
		":8082",
		"localhost.crt",
		"localhost.key",
		router(),
	)
	if err != nil {
		log.Fatalf("Server failed to start. Error: %s", err)
	}
}
