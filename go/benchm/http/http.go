package main

import (
	"log"
	"net/http"

	"github.com/wooknight/going_in_circles/go/benchm/http/handlers"
)

func main() {
	handlers.Routes()
	log.Println("Listener :started listening on http://localhost:4000")
	log.Fatalln(http.ListenAndServe(":4000", nil))
}
