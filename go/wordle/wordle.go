package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
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
	if ok, err := checkWord(str); ok && err == nil {
		fmt.Println(str, "is a valid Wordle word")
	}
}

type chrPresent map[int8]bool

func main() {

	//terminal()

	http.HandleFunc("/wordler", func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		notValid := make([]chrPresent, WORD_LENGTH)

		mMap := make(chrPresent)
		mMap['M'] = true
		notValid[3] = mMap

		bannedString := r.URL.Query().Get("banned")
		curStr := r.URL.Query().Get("current")
		if len(bannedString) == 0 || len(curStr) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "no parameters")
			return
		}
		bytes, err := base64.StdEncoding.DecodeString(curStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Error decoding current:", err)
			return
		}
		bytes = []byte(html.EscapeString(string(bytes)))
		if len(bytes) < 5 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "no parameters - current : ", string(bytes), " length : ", len(bytes))
			return
		}
		slate := [5]byte{bytes[0], bytes[1], bytes[2], bytes[3], bytes[4]}
		wordle([]byte(bannedString), slate, notValid, w)
	})

	port := ":8765"
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func terminal() bool {
	bannedString := flag.String("banned", "", "list of characters banned")
	curStr := flag.String("current", "", "current string in base 64")
	flag.Parse()
	if *bannedString == "" || *curStr == "" {
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := base64.StdEncoding.DecodeString(*curStr)
	if err != nil {
		fmt.Println("Error decoding environment variable:", err)
		return true
	}
	slate := [5]byte{bytes[0], bytes[1], bytes[2], bytes[3], bytes[4]}
	wordle([]byte(*bannedString), slate, os.Stdout)
	return false
}

func wordle(banned []byte, slate [5]byte, notValid []chrPresent, w io.Writer) {
	for i := 0; i < 5; i++ {
		if slate[i] >= byte('a') && slate[i] <= byte('z') {
			continue
		}
		for strt := 'a'; strt <= 'z'; strt++ {
			if bytes.Contains(banned, []byte(string(strt))) {
				continue
			}
			slate[i] = byte(strt)
			wordle(banned, slate, notValid, w)
		}
	}
	check(string(slate[:]), notValid)
}
