package main

import (
	"log"
	"net/http"

	"github.com/wooknight/GoingInCircles/Go/web-server/webhook"
)

func main() {
	http.HandleFunc("/orders", webhook.ProcessData)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
