package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	squares()
}

func squares() {
	sqHash := make(map[int]int64)
	for i := 1; i < math.MaxInt32; i++ {
		for j := i + 1; j < math.MaxInt32; j++ {
			for k := j + 1; k < math.MaxInt32; k++ {
				if _, iok := sqHash[i]; !iok {
					sqHash[i] = int64(i * i)
				}
				if _, jok := sqHash[j]; !jok {
					sqHash[j] = int64(j * j)
				}
				if _, kok := sqHash[k]; !kok {
					sqHash[k] = int64(k * k)
				}
				if (i+j+k)%10000 > 9995 {
					fmt.Println("Progessing -> ", i, j, k)
				}
				sum := sqHash[i] + sqHash[j] + sqHash[k]
				sqrt := math.Sqrt(float64(sum))
				if val, ok := sqHash[int(sqrt)]; ok {
					if val == sum {
						fmt.Println("Found candidates -> ", i, j, k, "Sum ->", sum, "Square Root ->", sqrt)
					}
				} else if int64(sqrt*sqrt) == sum {
					fmt.Println("Found candidates -> ", i, j, k, "Sum ->", sum, "Square Root ->", sqrt)
				}
			}
		}
	}
}
func cubes() {
	cubeHash := make(map[int]int64)
	max := 0
	for i := 1; i < math.MaxInt32; i++ {
		cubeHash[i] = int64(i * i * i)
		max = i
	}
	fmt.Println("Done ->", max)
}

func coursera() {
	var prod big.Float
	num1 := big.NewFloat(3141592653589793238462643383279502884197169399375105820974944592)

	num2 := big.NewFloat(2718281828459045235360287471352662497757247093699959574966967627)
	prod.Mul(num1, num2)
	fmt.Printf("Product is %s", prod.String())
}
