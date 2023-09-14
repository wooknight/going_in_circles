package main

import (
	"compress/gzip"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
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
type SafeAttribCount struct {
	mu      sync.Mutex
	attribs map[string]int
}

func PrintAllHTTOAttributes() {
	var counter SafeCount
	var attribCounter SafeAttribCount
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
	attribCounter.attribs = make(map[string]int)
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
					attribCounter.mu.Lock()
					attribCounter.attribs[a.Key]++
					attribCounter.mu.Unlock()
				}
				for c := h.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
			f(doc)

		}(urls[i], &wg)
	}
	wg.Wait()
	for key, val := range counter.attribs {
		fmt.Println("Key", key, "Value", val)
	}
	for key, val := range attribCounter.attribs {
		fmt.Println("Key", key, "Value", val)
	}

}

type wanker struct {
	Name         string
	Public_Repos int
}

func getGitHubInfo(login string) (string, int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Get called failed . Error = %v\n", err)
		return "", 0, err
	}
	var r wanker
	if resp.StatusCode != 200 {
		log.Printf("Get called failed . Status code = %d\n", resp.StatusCode)
		return "", 0, err

	}
	defer resp.Body.Close()
	decode := json.NewDecoder(resp.Body)
	if err = decode.Decode(&r); err != nil {
		log.Printf("Decoding failed . Error : %v : Reply : %#v\n", err, r)
		return "", 0, err
	}
	return r.Name, r.Public_Repos, nil
}

func getSha1(fileName string) (string, error) {
	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, fileName, nil)
	if err != nil {
		log.Printf("could not create request  :%s . Error : %v", fileName, err)
		return "", err
	}
	request.Header.Add("Accept-Encoding", "gzip")
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("could not open http file :%s . Error : %v", fileName, err)
		return "", err
	}
	defer resp.Body.Close()
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		log.Printf("got Content-Encoding gzip\n")
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Printf("could not open reader  . Error : %v", err)
			return "", err
		}
		defer resp.Body.Close()
	default:
		log.Printf("could not get Content-Encoding \n")
		reader = resp.Body

	}
	gunzipper, err := gzip.NewReader(reader)
	if err != nil {
		log.Printf("could not open zip. Error : %v", err)
		return "", err
	}

	shaNew := sha1.New()
	if _, err := io.Copy(shaNew, gunzipper); err != nil {
		log.Printf("could not write to sha. Error : %v", err)
		return "", err
	}
	sig := shaNew.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

func main() {
	// PrintAllHTTOAttributes()
	// conf, num, err := getGitHubInfo("wooknight")
	sha1, err := getSha1("https://www.353solutions.com/c/znga/data/http.log.gz")
	fmt.Println(sha1, err)
}
