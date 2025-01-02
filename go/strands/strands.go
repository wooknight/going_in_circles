package main

import (
	"fmt"
	"net/http"
	"strings"
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
			return true, nil
		}
		// If not, it's not a valid word
		if resp.StatusCode == http.StatusNotFound {
			return false, nil
		}
	}
}

// Helper function to generate permutations
func permute(str string, prefix string, results *[]string) {
	if len(str) == 0 {
		*results = append(*results, prefix)
		return
	}
	for i := 0; i < len(str); i++ {
		// Choose the current character and permute the rest
		permute(str[:i]+str[i+1:], prefix+string(str[i]), results)
	}
}

// Function to get all permutations of a string
func getPermutations(input string) []string {
	var results []string
	permute(input, "", &results)
	return results
}

func reverse(str string) string {
	var result string
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return result

}

func main() {
	// Input string
	input := "LACHIRIOTS"

	// Generate permutations
	permutations := getPermutations(input)
	banned := []string{
		"LI", "LR", "LO", "LT", "LS", "LH",
		"CO", "CT", "CS",
		"HO", "HT", "HS",
		"II",
	}
	// Print permutations

	for _, p := range permutations {
		parStr := strings.ToUpper(p)
		cont := false
		for _, subStr := range banned {
			if strings.Contains(parStr, subStr) {
				cont = true
				break
			}
			if strings.Contains(parStr, reverse(subStr)) {
				cont = true
				break
			}
		}
		if !cont {
			exists, err := checkWord(p)
			if err != nil && exists {
				fmt.Println(p)
			}
		}
	}
}
