package main

import (
	"log"
	"net/http"

	"github.com/wooknight/going_in_circles/go/web-server/webhook"
)

func main() {
	http.Handle("/orders", &webhook.SecretTokenHandler{
		Next:      webhook.NewUptimeHandler(),
		SecretJWT: "mongo",
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
