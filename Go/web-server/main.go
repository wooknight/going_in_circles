package main

import (
	"net/http"
	webhook "github.com/wooknight/GoingInCircles/Go/web-server/webhook"
)

func main() {
	http.ListenAndServe(webhook.processData)
}
