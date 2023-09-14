package main

import (
	"fmt"
	"log"
	"net/http"
)

func findFastest(urls []string) string {
	ch := make(chan string, 0)
	for _, url := range urls {

		go func(url string, ch chan string) {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			if resp.StatusCode == http.StatusOK {
				ch <- url
			}
		}(url, ch)

	}
	return <-ch
}
func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.tesla.com",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, findFastest(urls))

	})
	http.ListenAndServe(":8080", nil)
}
