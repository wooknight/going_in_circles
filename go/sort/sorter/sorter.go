package main

import (
	"math/rand"
	"sort"
)

const MAX_INT = int(^uint(0) >> 1)

// const MAX_INT = 100

func getMillion() ([]int, []int) {
	LEN := 1024 * 1024
	// LEN = 10
	results := make([]int, LEN)
	sorted := make([]int, LEN)
	for i := 0; i < LEN; i++ {
		results[i] = rand.Intn(MAX_INT)
	}
	copy(sorted, results)
	sort.Ints(sorted)
	return results, sorted
}
func partition(arr []int) int {
	pivot := len(arr) - 1
	lessThan := -1
	for i := 0; i < pivot; i++ {
		if arr[i] <= arr[pivot] {
			lessThan++
			arr[i], arr[lessThan] = arr[lessThan], arr[i]
		}
	}
	arr[lessThan+1], arr[pivot] = arr[pivot], arr[lessThan+1]
	pivot = lessThan + 1
	return pivot
}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := partition(arr)
	// fmt.Println(arr, p)
	if p > 0 {
		quicksort(arr[:p])
	}
	if p < len(arr)-1 {
		quicksort(arr[p+1:])
	}
}
