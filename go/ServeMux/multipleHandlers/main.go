package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type UUID struct {
}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		c := 10
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			log.Fatal("Error %v", err)

		}
		fmt.Fprintf(w, fmt.Sprintf("%x", b))
		return
	}
	http.NotFound(w, r)
	return
}

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})
	newMux.HandleFunc("/heartBeat", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Fprintln(w, currentTime.String())
	})
	go func() {
		mux := &UUID{}
		http.ListenAndServe(":9002", mux)
	}()

	http.ListenAndServe(":9001", newMux)

}
