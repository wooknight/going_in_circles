package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wooknight/going_in_circles/go/web-server/webhook"
)

func main() {
	// http.Handle("/orders", &webhook.SecretTokenHandler{
	// 	Next:      webhook.NewUptimeHandler(),
	// 	SecretJWT: "mongo",
	// })
	// log.Fatal(http.ListenAndServe(":8000", nil))
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", webhook.ArticleHandler)
	srv := http.Server{Handler: r, Addr: ":8000", WriteTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second}
	log.Fatalln(srv.ListenAndServe())
}
