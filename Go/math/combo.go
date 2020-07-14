package main

import "fmt"

func main() {
	var prob float64 = 1
	var num_days float64 = 365
	for i := 1; i < 70; i++ {
		prob = prob * (float64)((num_days-(float64)(i))/num_days)
		fmt.Printf("for %d of people , the probability is %f\n", i, prob)
	}
	fmt.Println(prob)
}
