package main

//Installing ligature fonts
// brew tap homebrew/cask-fonts
// brew search font- | grep cascadia
//brew install --cask font-fira-code
//brew install --cask font-cascadia-code

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const bsaEndpoint = "http://api.bart.gov/api/bsa.aspx?cmd=bsa&key=MW9S-E7SL-26DU-VV8V&json=y"

type bsaResponse struct {
	Root struct {
		Advisories []BSA `json:"bsa"`
	}
}

// BSA is a BART service advisory
type BSA struct {
	Station     string
	Description struct {
		Text string `json:"#cdata-section"`
	}
}

func poll() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for tm := <-ticker.C; tm.Second() > 5; {
		resp, err := http.Get(bsaEndpoint)
		if err != nil {
			log.Println("ERROR: could not GET bsaEndpoint:", err)
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("ERROR: could not parse response body:", err)
			}
			var br bsaResponse
			err = json.Unmarshal(b, &br)
			if err != nil {
				log.Println("ERROR: json.Unmarshal:", err)
			}
			if len(br.Root.Advisories) > 0 {
				for _, adv := range br.Root.Advisories {
					log.Println(adv.Station, adv.Description.Text)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		go func() {
			log.Println("hello, world 1")
		}()
	})
	go poll()
	log.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}
