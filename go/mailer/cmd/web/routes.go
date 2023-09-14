package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("starting web server...")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Panic(err)
	}
}

func (app *Config) routes() http.Handler {
	//create router
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)
	//application routes
	mux.Get("/", app.HomePage)
	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount)
	mux.Mount("/members", app.authRouter())
	return mux

}

func (app *Config) authRouter() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.Auth)
	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)
	return mux
}
