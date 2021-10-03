package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// const MAX_INT = int(^uint(0) >> 1)

const MAX_INT = 100

func getMillion() ([]int, []int) {
	LEN := 1024 * 1024
	LEN = 10
	results := make([]int, LEN)
	sorted := make([]int, LEN)
	for i := 0; i < LEN; i++ {
		results[i] = rand.Intn(MAX_INT)
	}
	copy(sorted, results)
	sort.Ints(sorted)
	return results, sorted
}
func partitionSlice(arr []int) int {
	pivot := len(arr) - 1
	lessThan := -1
	for i := 0; i < pivot; i++ {
		if arr[i] <= arr[pivot] {
			lessThan++
			arr[i], arr[lessThan] = arr[lessThan], arr[i]
		}
	}
	arr[lessThan+1], arr[pivot] = arr[pivot], arr[lessThan+1]
	return lessThan + 1
}

func quicksortSlice(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := partitionSlice(arr)
	// fmt.Println(arr, p)
	if p > 0 {
		quicksortSlice(arr[:p])
	}
	if p < len(arr)-1 {
		quicksortSlice(arr[p+1:])
	}
}

func partition(arr []int, start, end int) int {
	pivot := end
	lt := start - 1
	for i := start; i < pivot; i++ {
		if arr[i] <= arr[pivot] {
			lt++
			arr[lt], arr[i] = arr[i], arr[lt]
		}
	}
	arr[lt+1], arr[pivot] = arr[pivot], arr[lt+1]
	// fmt.Println(arr,"Start", start,"End", end,"Original Pvt", pivot,"New Pivot", lt+1, "Pvt value", arr[pivot])
	return lt + 1
}

func quicksort(arr []int, start, end int) {
	if start < end {
		pvt := partition(arr, start, end)
		quicksort(arr, start, pvt-1)
		quicksort(arr, pvt+1, end)
	}
}

func bubblesort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func insertionsortSlice(arr []int) {

}
func insertionSort(arr []int) {

}
func MergeSortSlice(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		return merge(MergeSortSlice(arr[:mid]), MergeSortSlice(arr[mid+1:]))
	}
	fmt.Println(arr)
	return arr
}

func merge(arr1, arr2 []int) []int {
	arr1Idx := 0
	arr2Idx := 0
	totLen := len(arr1)+len(arr2)
	myarr := make([]int, totLen)
	for  k:=0;k<totLen;k++ {
		if arr1Idx < len(arr1) && arr2Idx < len(arr2) {
			if arr1[arr1Idx] < arr2[arr2Idx] {
				myarr[k] = arr1[arr1Idx]
				arr1Idx++
			} else {
				myarr[k] = arr2[arr2Idx]
				arr2Idx++
			}
		} else if arr1Idx < len(arr1) {
			myarr[k] = arr1[arr1Idx]
			arr1Idx++
		} else if arr2Idx < len(arr2) {
			myarr[k] = arr2[arr2Idx]
			arr2Idx++
		}
	}
	fmt.Println(arr1, arr2, myarr)
	return myarr
}

func Mergesort(arr []int, start, end int) []int {
	if len(arr) == 1 || start >= end {
		return arr
	}else if len(arr) == 2{
		if arr[0] > arr[1] {
			return []int{arr[1],arr[0]}
		}
		return arr
	}
	if len(arr) > 1 && start < end {
		mid := start + (end - start / 2)
			return merge(Mergesort(arr, start ,  mid ), Mergesort(arr, mid+1, end))

	}
	fmt.Println(arr)
	return arr
}

func main() {
	fmt.Println("Hello world")
}
