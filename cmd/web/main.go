package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/snippet", showSnippet)
	router.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", router)

	log.Fatal(err)
}
