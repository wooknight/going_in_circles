package main

//go get github.com/oxequa/realize;rm -rf .realize.yaml;realize start --run routers.go
import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const FQDN = "fqdn"
const CATEGORY = "category"
const BUSINESS = "business"
const RATING = "rating"

func main() {
	var wait time.Duration
	r := mux.NewRouter()
	cat := r.PathPrefix("/" + CATEGORY).Subrouter()
	cat.HandleFunc("/", categoryHandler)
	business := r.PathPrefix("/" + BUSINESS).Subrouter()
	business.HandleFunc("/", businessHandler)
	rating := r.PathPrefix("/" + RATING).Subrouter()
	rating.HandleFunc("/", ratingsHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
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
