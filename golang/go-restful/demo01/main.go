package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/hello", hello)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: serveMux}
	log.Fatal(server.ListenAndServe())
}

func hello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "world")
}
