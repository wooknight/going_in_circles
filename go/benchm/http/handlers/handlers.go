package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendjson", SendJson)
	http.HandleFunc("/sendxml", SendXML)
}

var u = struct {
	Name  string
	Email string
}{
	Name:  "Ramesh",
	Email: "ramesh@naidu.net",
}

func SendJson(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%+v\n", rw)
	json.NewEncoder(rw).Encode(&u)
}

func SendXML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/xml")
	w.WriteHeader(200)
	xml.NewEncoder(w).Encode(&u)
}
