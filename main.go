package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("GET /hello", hello)
	http.HandleFunc("GET /hello2", hello2)
	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
