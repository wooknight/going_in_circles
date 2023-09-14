package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func corsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func query(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://ec2-35-164-145-227.us-west-2.compute.amazonaws.com/cors")
	if err != nil {
		log.Fatal("Encountered an error while querying")
	}
	p := make([]byte, resp.ContentLength)
	dt, err := resp.Body.Read(p)
	if err != nil && err != io.EOF {
		log.Fatal("Could not read", err)
	}

	fmt.Fprintln(w, dt, p)
}

func main() {
	http.HandleFunc("/cors", corsHandler)
	http.HandleFunc("/q", query)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
