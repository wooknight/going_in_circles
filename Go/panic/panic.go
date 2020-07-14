package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Inside main defer")
		if r := recover(); r != nil && r == "Bhosada lag gaya" {
			fmt.Println("I love pussy ")
		}
	}()
	fmt.Println("Inside main and getting ready to call bho")
	bho()
}

func bho() {
	defer func() {
		fmt.Println("Inside bho defer\nI love bhosada my dear")
	}()

	fmt.Println("Getting ready to do a panic")
	panic("Bhosada lag hgaya")
}
