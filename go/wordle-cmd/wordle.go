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
var wordCount int
var mapWords map[string]bool

func checkWord(word string) (bool, error) {
	for {
		if (wordCount % 1000) == 0 {
			fmt.Printf("processed %d words . Currently processing %s . %c is the %d letter \r", wordCount, word, word[0], (word[0]-'A')+1)
		}
		wordCount++

		if _, ok := mapWords[word]; ok {
			return false, nil
		}
		mapWords[word] = true
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
		fmt.Printf("\n%s is a valid Wordle word\n", str)
	}
}

type chrPresent map[int8]bool

func main() {
	mapWords = make(map[string]bool)
	chrMapPos := make(map[int]byte)
	chrMapPos[1] = 'I'
	chrMapPos[3] = 'A'
	chrsNotPresent := []byte("EUPDLCVNM")

	notValid := make([]chrPresent, WORD_LENGTH)
	mMap := make(chrPresent)
	mMap['S'] = true
	notValid[3] = mMap
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
