package main

import (
	"flag"
	"log"
	"strings"
)

func main() {

	l := flag.String("l", "", "count lines")
	w := flag.String("w", "", "count words")
	c := flag.String("c", "", "count bytes")
	flag.Parse()
	if len(*l) > 0 {
		log.Println("counting lines")
		arr := strings.Split(*l, "\n")
		log.Println("number of words is ", len(arr))

	}
	if len(*w) > 0 {
		log.Println("counting words")
		arr := strings.Split(*w, " ")
		log.Println("number of words is ", len(arr))

	}
	if len(*c) > 0 {
		log.Println("counting bytes")

		log.Println("number of words is ", len(*c))

	}
}
