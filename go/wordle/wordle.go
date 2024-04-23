package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const WORD_LENGTH = 5

// declaring the variable which is an ASCII value of A
const A byte = 'A'

var wordleWords string

func check(str string) {
	if strings.Contains(wordleWords, str) {
		fmt.Println(str, "is a valid Wordle word")
	}
}

func main() {
	chrMapPos := make(map[int]byte)
	chrsPresentButWrongPos := make([]byte, 0)
	chrsNotPresent := make([]byte, 26)
	chrsNotPresent[0] = 'R'
	chrsNotPresent[1] = 'E'
	chrsNotPresent[2] = 'A'
	chrsNotPresent[3] = 'T'
	chrsNotPresent[4] = 'C'
	chrsNotPresent[5] = 'K'
	chrsNotPresent[6] = 'M'
	chrsNotPresent[7] = 'U'
	chrsNotPresent[8] = 'S'
	chrMapPos[1] = 'U'
	chrMapPos[2] = 'S'
	chrsPresentButWrongPos = append(chrsPresentButWrongPos, 'I')

	//read wordle into string
	resp, err := http.Get("https://www.wordunscrambler.net/word-list/wordle-word-list")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	wordlebytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	wordleWords = string(wordlebytes)
	slate := []byte{}
	var gen func(int, []byte)
	gen = func(pos int, slate []byte) {
		if pos == WORD_LENGTH {
			check(string(slate))
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
