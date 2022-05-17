package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wooknight/going_in_circles/go/microservice/config"
	"github.com/wooknight/going_in_circles/go/microservice/handler"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/business/list", handler.Repo.ListData)
	mux.Get("/business/add", handler.Repo.AddData)
	return mux
}
