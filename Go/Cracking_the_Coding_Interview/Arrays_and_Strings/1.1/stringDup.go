package main

import (
	"fmt"
	"os"
)

func main() {
	curPath, _ := os.Getwd()
	fmt.Printf("We are in folder %s", curPath)
}

func checkDupsWithMap(str string) bool {
	hasher := make(map[rune]bool)
	for _, val := range str {
		if _, ok := hasher[val]; ok {
			return true
		}
		hasher[val] = true
	}
	return false
}
func checkDupsWith2Ptrs(str string) bool {

	for i, val := range str[:len(str)-1] {
		for _, val1 := range str[i:] {
			if val == val1 {
				return true
			}
		}
	}
	return false
}
