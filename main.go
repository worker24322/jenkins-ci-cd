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

func hello3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! 3")
}
func main() {
	http.HandleFunc("GET /hello", hello)
	http.HandleFunc("GET /hello2", hello2)
	http.HandleFunc("GET /hello3", hello3)
	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
