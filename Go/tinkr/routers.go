package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const FQDN = "fqdn"
const CATEGORY = "category"
const BUSINESS = "business"
const RATING = "rating"

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/"+CATEGORY, categoryHandler)
	r.HandleFunc("/"+BUSINESS, businessHandler)
	r.HandleFunc("/"+RATING, ratingsHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
func categoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Category"))
}

func businessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Business"))
}

func ratingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rating"))
}
