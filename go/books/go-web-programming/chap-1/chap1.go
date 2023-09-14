package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world %v", time.Now())
	})

	http.ListenAndServe(":8080", nil)
}
