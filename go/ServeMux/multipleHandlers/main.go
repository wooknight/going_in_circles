package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
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
			log.Fatalf("Error %v", err)

		}
		fmt.Fprintf(w, "%s", fmt.Sprintf("%x", b))
		return
	}
	http.NotFound(w, r)
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
