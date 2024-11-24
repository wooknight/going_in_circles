package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

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
			fmt.Println(url)
			return true, nil
		}

		// If not, it's not a valid word
		if resp.StatusCode == http.StatusNotFound {
			return false, nil
		}
		fmt.Println(url, resp.StatusCode)
	}
}

func main() {

	bannedString := flag.String("banned", "", "list of characters banned")
	curStr := flag.String("current", "", "current string in base 64")
	flag.Parse()
	if *bannedString == "" || *curStr == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Decode the Base64 string into a byte slice
	bytes, err := base64.StdEncoding.DecodeString(*curStr)
	if err != nil {
		fmt.Println("Error decoding environment variable:", err)
		return
	}
	wordle([]byte(*bannedString), [5]byte{bytes[0], bytes[1], bytes[2], bytes[3], bytes[4]})
}

func wordle(banned []byte, slate [5]byte) {
	for i := 0; i < 5; i++ {
		if slate[i] >= byte('a') && slate[i] <= byte('z') {
			continue
		}
		for strt := 'a'; strt <= 'z'; strt++ {
			if bytes.Contains(banned, []byte(string(strt))) {
				continue
			}
			slate[i] = byte(strt)
			wordle(banned, slate)
		}
	}
	valid, err := checkWord(string(slate[:]))
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	} else if valid {
		fmt.Println(string(slate[:]))
	}
}
