package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type response struct {
	FastestUrl   string        `json:"fastest_url"`
	ResponseTime time.Duration `json:"response_time"`
}

func GetURLS(urls []string) []string {
	resp, err := http.Get("https://www.debian.org/mirror/list")
	if err != nil {
		fmt.Printf("Error while opening page %v\n", err)
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error while reading data %v\n", err)
	}
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.Contains(a.Val, "http") == true {
					urls = append(urls, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls
}
func FindFastest(urls []string) response {
	urlChan := make(chan string)
	timeChan := make(chan time.Duration)
	for _, val := range urls {
		currUrl := val
		go func(url string) {
			start := time.Now()
			// fmt.Println(url)
			_, err := http.Get(url + "README")
			latency := time.Since(start) / time.Millisecond
			if err == nil {
				urlChan <- url
				timeChan <- latency
			}
		}(currUrl)
	}
	return response{<-urlChan, <-timeChan}
}

func main() {
	urls := make([]string, 0)
	urls = GetURLS(urls)
	fmt.Println("URL Count ", len(urls))
	http.HandleFunc("/fastest_mirror", func(w http.ResponseWriter, r *http.Request) {
		response := FindFastest(urls)
		respJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)
	})
	port := ":8000"
	server := http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n\n", port)
	log.Fatal(server.ListenAndServe())
}
