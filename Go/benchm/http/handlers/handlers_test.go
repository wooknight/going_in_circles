package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wooknight/GoingInCircles/Go/benchm/http/handlers"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJson(t *testing.T) {
	url := "/sendjson"
	statusCode := http.StatusOK
	t.Log("Given the need to test the SendJson endpoint")
	{
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould receive a status code %d for the response. Received[%d]", failed, statusCode, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code %d for the response. Received[%d]", succeed, statusCode, w.Code)
			var u struct {
				Name  string
				Email string
			}
			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tShould be able to decode the response [%v]", failed, err)
			}
			t.Logf("\t%s\tShould be able to decode the response", succeed)

		}
	}
}
