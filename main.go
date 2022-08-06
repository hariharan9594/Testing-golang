package main

import (
	"io"
	"log"
	"os"

	//"test/hari/model"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	http.HandleFunc("/post", post)

	log.Print("Lisitening on Port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, This is hari..Nice to see you...")
}
func post(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This page gives u details about live market...")
}
