package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type response struct {
	FastestURL string
	Latency    time.Duration
}

// MirrorList is list of debian mirror sites
var MirrorList = []string{
	"http://ftp.am.debian.org/debian/", "http://ftp.au.debian.org/debian/",
	"http://ftp.at.debian.org/debian/", "http://ftp.by.debian.org/debian/",
	"http://ftp.be.debian.org/debian/", "http://ftp.br.debian.org/debian/",
	"http://ftp.bg.debian.org/debian/", "http://ftp.ca.debian.org/debian/",
	"http://ftp.cl.debian.org/debian/", "http://ftp2.cn.debian.org/debian/",
	"http://ftp.cn.debian.org/debian/", "http://ftp.hr.debian.org/debian/",
	"http://ftp.cz.debian.org/debian/", "http://ftp.dk.debian.org/debian/",
	"http://ftp.sv.debian.org/debian/", "http://ftp.ee.debian.org/debian/",
	"http://ftp.fr.debian.org/debian/", "http://ftp2.de.debian.org/debian/",
	"http://ftp.de.debian.org/debian/", "http://ftp.gr.debian.org/debian/",
	"http://ftp.hk.debian.org/debian/", "http://ftp.hu.debian.org/debian/",
	"http://ftp.is.debian.org/debian/", "http://ftp.it.debian.org/debian/",
	"http://ftp.jp.debian.org/debian/", "http://ftp.kr.debian.org/debian/",
	"http://ftp.lt.debian.org/debian/", "http://ftp.mx.debian.org/debian/",
	"http://ftp.md.debian.org/debian/", "http://ftp.nl.debian.org/debian/",
	"http://ftp.nc.debian.org/debian/", "http://ftp.nz.debian.org/debian/",
	"http://ftp.no.debian.org/debian/", "http://ftp.pl.debian.org/debian/",
	"http://ftp.pt.debian.org/debian/", "http://ftp.ro.debian.org/debian/",
	"http://ftp.ru.debian.org/debian/", "http://ftp.sg.debian.org/debian/",
	"http://ftp.sk.debian.org/debian/", "http://ftp.si.debian.org/debian/",
	"http://ftp.es.debian.org/debian/", "http://ftp.fi.debian.org/debian/",
	"http://ftp.se.debian.org/debian/", "http://ftp.ch.debian.org/debian/",
	"http://ftp.tw.debian.org/debian/", "http://ftp.tr.debian.org/debian/",
	"http://ftp.uk.debian.org/debian/", "http://ftp.us.debian.org/debian/",
}

func main() {
	fmt.Println("Staring the server")
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(MirrorList)
		respJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You git!! you git the rest GET"))
	})
	port := ":8080"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	for _, url := range urls {
		mirrorURL := url
		go func() {
			log.Println("Started probing: ", mirrorURL)
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if nil == err {
				urlChan <- mirrorURL
				latencyChan <- latency

			}
			log.Printf("Mirror :%s \nLatency %s", mirrorURL, latency)
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func post() {
	//add a new category
	fmt.Printf("Posting")
}

func get() {
	//retreive the category from mongo and return the object
	fmt.Printf("Get")
}

func patch() {
	//update the object
	fmt.Printf("pathc")
}

func delete() {
	//delete the object
	fmt.Printf("delete")
}

func postTouchpoint(category string) {
	//add a new category
	fmt.Printf("Posting")
}

func getTouchpoints() {
	//retreive the category from mongo and return the object
	fmt.Printf("Get")
}

func patchTouchpoints() {
	//update the object
	fmt.Printf("pathc")
}

func delete() {
	//delete the object
	fmt.Printf("delete")
}
