package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/wooknight/going_in_circles/go/microservice/business"
	"github.com/wooknight/going_in_circles/go/microservice/config"
	"github.com/wooknight/going_in_circles/go/microservice/handler"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app = config.AppConfig{}
	handler.NewHandlers(handler.NewRepo(&app))
	app.IsProd = false
	session = scs.New()
	session.Lifetime = 5 * time.Minute
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProd
	app.Session = session
	tc, err := business.RenderTemplateTotal()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	business.NewTemplates(&app)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: Routes(&app),
	}
	log.Fatal(srv.ListenAndServe())
}
