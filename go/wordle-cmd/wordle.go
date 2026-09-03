package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

const dictionaryPath = "/usr/share/dict/words"

const WORD_LENGTH = 5

// declaring the variable which is an ASCII value of A
const A byte = 'A'

var wordleWords string
var wordCount int
var dictionary map[string]bool

func toUpperByte(b byte) byte {
	return byte(strings.ToUpper(string(b))[0])
}

func isUpperAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return false
		}
	}
	return true
}

func loadDictionary() error {
	file, err := os.Open(dictionaryPath)
	if err != nil {
		return err
	}
	defer file.Close()

	dictionary = make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToUpper(strings.TrimSpace(scanner.Text()))
		if len(word) != WORD_LENGTH || !isUpperAlpha(word) {
			continue
		}
		dictionary[word] = true
	}
	return scanner.Err()
}

func checkWord(word string) (bool, error) {
	if (wordCount % 1000) == 0 {
		fmt.Printf("processed %d words . Currently processing %s . %c is the %d letter \r", wordCount, word, word[0], (word[0]-'A')+1)
	}
	wordCount++

	return dictionary[word], nil
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
	if err := loadDictionary(); err != nil {
		fmt.Printf("failed to load dictionary from %s: %v\n", dictionaryPath, err)
		os.Exit(1)
	}

	chrMapPos := make(map[int]byte)
	chrMapPos[1] = toUpperByte('o')
	chrMapPos[2] = toUpperByte('i')
	chrMapPos[4] = toUpperByte('T')
	chrsNotPresent := []byte("EUPADN")
	for key, ch := range chrsNotPresent {
		chrsNotPresent[key] = toUpperByte(ch)
	}
	fmt.Printf("chrMapPos: %v\n", chrMapPos)
	for _, ch := range chrsNotPresent {

		if ch < 'A' || ch > 'Z' {
			fmt.Printf("Invalid character in chrsNotPresent: %c\n", ch)
			os.Exit(1)
		}
	}

	notValid := make([]chrPresent, WORD_LENGTH)
	notValid[0] = make(chrPresent)
	// notValid[0]['A'] = true
	// notValid[0]['I'] = true
	notValid[1] = make(chrPresent)
	// notValid[1]['O'] = true
	// notValid[1]['I'] = true
	notValid[2] = make(chrPresent)
	notValid[3] = make(chrPresent)
	// notValid[3]['O'] = true
	// notValid[3]['U'] = true
	notValid[4] = make(chrPresent)
	// notValid[4]['T'] = true
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
	fmt.Printf("\nTotal words processed: %d\n-------------------\n", wordCount)
}
