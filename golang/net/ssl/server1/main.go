package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServeTLS(
		":8080",
		"../cert/server/server.pem",
		"../cert/server/server-key.pem",
		nil),
	)
}
