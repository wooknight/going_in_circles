package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJson() {
	r := httptest.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)
	var u struct {
		Name  string
		Email string
	}
	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Printf("\t%s\tShould be able to decode the response [%v]", failed, err)
	}
	fmt.Println(u)
	//Output:
	//{Ramesh ramesh@naidu.net}
}
