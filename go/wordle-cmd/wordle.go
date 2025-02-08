package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const WORD_LENGTH = 5

// declaring the variable which is an ASCII value of A
const A byte = 'A'

var wordleWords string

func checkWord(word string) (bool, error) {
	for {
		url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
		req, _ := http.NewRequest("GET", url, nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusTooManyRequests {
			time.Sleep(1 * time.Second)
			continue
		}
		// If the status code is 200, the word is valid
		if resp.StatusCode == http.StatusOK {
			return true, nil
		}
		// If not, it's not a valid word
		if resp.StatusCode == http.StatusNotFound {
			return false, nil
		}
	}
}

func check(str string, present []chrPresent) {

	for i, ch := range str {
		if _, ok := present[i][int8(ch)]; ok {
			return
		}
	}
	chrsToCheckFor := make(chrPresent)
	for _, mapper := range present {
		for key := range mapper {
			chrsToCheckFor[key] = false
		}

	}
	for _, ch := range str {
		if _, ok := chrsToCheckFor[int8(ch)]; ok {
			chrsToCheckFor[int8(ch)] = true
		}
	}
	for _, exists := range chrsToCheckFor {
		if !exists {
			return
		}
	}
	if ok, err := checkWord(str); ok && err == nil {
		fmt.Println(str, "is a valid Wordle word")
	}
}

type chrPresent map[int8]bool

func main() {
	chrMapPos := make(map[int]byte)
	chrMapPos[0] = 'S'
	chrMapPos[2] = 'A'
	chrMapPos[3] = 'T'
	chrsNotPresent := []byte("ERUIOPDKLCN")

	notValid := make([]chrPresent, WORD_LENGTH)

	cMap := make(chrPresent)
	cMap['T'] = true
	notValid[4] = cMap
	notValid[1] = cMap
	slate := []byte{}
	var gen func(int, []byte)
	gen = func(pos int, slate []byte) {
		if pos == WORD_LENGTH {
			check(string(slate), notValid)
			return
		}
		if val, ok := chrMapPos[pos]; ok {
			slate = append(slate, val)
			gen(pos+1, slate)
			return
		}
		for ch := 'A'; ch <= 'Z'; ch++ {
			if !bytes.Contains(chrsNotPresent, append([]byte{}, byte(ch))) {
				gen(pos+1, append(slate, byte(ch)))
			}
		}
	}
	gen(0, slate)
}
