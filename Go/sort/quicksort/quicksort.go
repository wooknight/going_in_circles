package main

// "fmt"
// "os"
// "runtime"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr)-1 > 0 {
		pvt := partition(arr)
		if pvt > 0 {
			quicksort(arr[:pvt])
		}
		quicksort(arr[pvt:])
	}

	return arr
}

func partition(arr []int) int {
	greaterThanPvt := 0
	if len(arr) <= 1 {
		return greaterThanPvt
	}
	pvt := len(arr) - 1
	for i := range arr {
		if arr[i] < arr[pvt] {
			arr[i], arr[greaterThanPvt] = arr[greaterThanPvt], arr[i]
			greaterThanPvt++
		}
	}
	arr[pvt], arr[greaterThanPvt] = arr[greaterThanPvt], arr[pvt]
	return greaterThanPvt
}
