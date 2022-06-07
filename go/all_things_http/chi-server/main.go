package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})
	r.Route("/tink", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Route("/{choot}", func(r chi.Router) {
			r.Use(chootHandler)
			r.Get("/", getChoot)
			r.Put("/", putChoot)
			r.Delete("/", deleteChoot)
		})
	})
	http.ListenAndServe(":8080", r)
}

func chootHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if chootID := chi.URLParam(r, "choot"); chootID != "" {
			fmt.Println(chootID, "  is what we found ")
			ctx := context.WithValue(r.Context(), "choot", chootID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.NotFound(w, r)
			return
		}
	})

}

func deleteChoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	choot, _ := ctx.Value("choot").(*int)
	fmt.Fprintln(w, "delete choot #", choot)

}

func putChoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	choot, _ := ctx.Value("choot").(*int)
	fmt.Fprintln(w, "put choot #", choot)
}

func getChoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	choot, _ := ctx.Value("choot").(string)
	fmt.Fprintln(w, "get choot #", choot)
}
