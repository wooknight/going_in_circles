package main

import (
	"GoingInCircles/Go/benchm/http/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()
	log.Println("Listener :started listening on http://localhost:4000")
	log.Fatalln(http.ListenAndServe(":4000", nil))
}
