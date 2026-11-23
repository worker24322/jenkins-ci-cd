package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("GET /hello", hello)

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
