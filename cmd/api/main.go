package main

import (
	"log"
	"net/http"
	"telem.kmani/internal/api"
)

func main() {
	router := api.NewRouter()
	log.Println("Starting API server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
