package main

import (
	"math/rand"
)

func partition(arr []int) int {
	pivot := rand.Intn(len(arr) - 1) //random pivot
	greatrThan := len(arr) - 1
	// pivot := greatrThan
	lessThan := 0
	for greatrThan > lessThan {
		for arr[lessThan] <= arr[int(pivot)] {
			lessThan++
		}
		for arr[greatrThan] > arr[pivot] {
			greatrThan--
		}
		if greatrThan > lessThan {
			arr[lessThan], arr[greatrThan] = arr[greatrThan], arr[lessThan]
		}
	}
	return pivot
}

func quicksort(arr []int) {
	p := partition(arr)
	quicksort(arr[1:p])
	quicksort(arr[p+1:])
}
