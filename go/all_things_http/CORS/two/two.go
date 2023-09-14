package main

import (
	"fmt"
	"log"
	"net/http"
)

func corsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/cors", corsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
