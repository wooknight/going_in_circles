package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Inside my defer")
		if r := recover(); r != nil && r == "Bhosada lag hgaya" {
			fmt.Println("I love pussy ")
		}
	}()
	fmt.Println("Getting ready to do a panic")
	panic("Bhosada lag hgaya")

}
