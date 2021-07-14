package main

import "math/rand"

// "fmt"
// "os"
// "runtime"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	randIndex := rand.Intn(len(arr))
	pvt := len(arr) - 1 //use the last element
	arr[randIndex], arr[pvt] = arr[pvt], arr[randIndex]
	greaterThanPvt := 0

	for i:=0;i<len(arr);i++ {
		if arr[i] < arr[pvt] {
			arr[i], arr[greaterThanPvt] = arr[greaterThanPvt], arr[i]
			greaterThanPvt++
		}
	}
	arr[pvt], arr[greaterThanPvt] = arr[greaterThanPvt], arr[pvt]
	quicksort(arr[:greaterThanPvt])
	quicksort(arr[greaterThanPvt+1:])

	return arr
}
