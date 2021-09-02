package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

type att struct {
	key      string
	val      string
	attrType string
	data     string
}
type SafeCount struct {
	mu      sync.Mutex
	attribs map[att]int
}

func main() {
	var counter SafeCount
	var wg sync.WaitGroup
	resp, err := http.Get("https://www.debian.org/mirror/list")
	if err != nil {
		log.Fatalln("Failed to connect ", err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalln("Failed to parse ", err)
	}
	urls := make([]string, 0)
	var AddUrls func(n *html.Node)
	AddUrls = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, c := range n.Attr {
				if c.Key == "href" {
					urls = append(urls, c.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			AddUrls(c)
		}
	}
	AddUrls(doc)
	counter.attribs = make(map[att]int)
	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {

		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error while opening %s  Error:%v\n", url, err)
				return
			}
			defer resp.Body.Close()
			doc, err := html.Parse(resp.Body)
			if err != nil {
				fmt.Printf("Error while parsing %s  Error:%v\n", url, err)
				return
			}

			var f func(h *html.Node)
			f = func(h *html.Node) {
				for _, a := range h.Attr {
					counter.mu.Lock()
					counter.attribs[att{attrType: fmt.Sprintf("%d", h.Type), data: h.Data, key: a.Key, val: a.Val}]++
					counter.mu.Unlock()
				}
				for c := h.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
			f(doc)

		}(urls[i], &wg)
	}
	wg.Wait()
	fmt.Println(counter.attribs)
}
