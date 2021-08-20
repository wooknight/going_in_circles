package main

import "fmt"
import "math/big"

func main() {
	var prob float64 = 1
	var num_days float64 = 365
	coursera()
	
	for i := 1; i < 70; i++ {
		prob = prob * (float64)((num_days-(float64)(i))/num_days)
		// fmt.Printf("for %d of people , the probability is %f\n", i, prob)
	}
	// fmt.Println(prob)

}


func coursera(){
	var prod big.Float
	num1  := big.NewFloat(3141592653589793238462643383279502884197169399375105820974944592)

	num2  := big.NewFloat(2718281828459045235360287471352662497757247093699959574966967627)
	prod.Mul(num1,num2)
	fmt.Printf("Product is %s",prod.String())	
}