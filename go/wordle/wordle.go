package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var cache map[string]string

const filePath = "data.json"

func saveMapToFile(filePath string, data map[string]string) error {
	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Serialize the map to JSON and write to the file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode map: %v", err)
	}
	return nil
}
func loadMapFromFile(filePath string) (map[string]string, error) {
	data := make(map[string]string)

	// Open the file for reading

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file does not exist")
		return data, nil
	}
	defer file.Close()

	// Decode the JSON into a map
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode map: %v", err)
	}
	return data, nil
}

func checkWord(word string) (bool, error) {
	if valid, ok := cache[word]; ok {
		return valid == "true", nil
	}
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
			cache[word] = "true"
			return true, nil
		}
		// If not, it's not a valid word
		if resp.StatusCode == http.StatusNotFound {
			cache[word] = "false"
			return false, nil
		}
	}
}

func main() {
	var err error
	// Load the map back
	cache, err = loadMapFromFile(filePath)
	if err != nil {
		fmt.Println("Error loading map:", err)
		return
	}
	// fmt.Println("Loaded data:", cache)

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
	slate := [5]byte{bytes[0], bytes[1], bytes[2], bytes[3], bytes[4]}
	wordle([]byte(*bannedString), slate)
	saveMapToFile(filePath, cache)
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
