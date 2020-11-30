package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	fmt.Println("hello world")
}

func mergesort(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	} else if len(arr2) == 0 {
		return arr1
	}
	if len(arr1) <= 3 && len(arr2) <= 3 {
		return merge(arr1, arr2)
	}
	arr1 = mergesort(arr1[:len(arr1)/2], arr1[len(arr1)/2+1:])
	arr2 = mergesort(arr2[:len(arr2)/2], arr2[len(arr2)/2+1:])
	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) []int {
	totLen := len(arr1) + len(arr2)
	sortedArr := make([]int, totLen)
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
