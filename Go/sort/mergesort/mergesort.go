package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	fmt.Println("hello world")
}

func merge_sort(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}
	arr1 := merge_sort(arr[:len(arr)/2])
	arr2 := merge_sort(arr[len(arr)/2:len(arr)])
	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int32) []int32 {
	totLen := len(arr1) + len(arr2)
	sortedArr := make([]int32, totLen)
	i := 0
	j := 0
	for k := 0; k < totLen; k++ {
		if i < len(arr1) && j < len(arr2){
			if arr1[i]<arr2[j]{
				sortedArr[k]=arr1[i]
				i++
			}else{
				sortedArr[k]=arr2[j]
				j++
			}
		}else if i < len(arr1) {
			sortedArr[k] = arr1[i]
			i++
		} else if j < len(arr2) {
				sortedArr[k]=arr2[j]
				j++
			}
	}
	return sortedArr
}

func insertionSort(arr []int) []int {
	for key, val := range arr {
		subarr := arr[0 : len(arr)-1]
		for key1, val1 := range subarr {
			if val > val1 {
				arr[key], subarr[key1] = val1, val
			}
		}
	}
	return arr
}
