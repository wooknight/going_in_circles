package main

import (
	"fmt"
"runtime"
"os"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	fmt.Println("hello world")
}

func printStack() {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func merge_sort(arr []int) []int {
	defer func() {
		fmt.Println("Inside main defer")

		if r := recover(); r != nil && r == "Aaankh lag gaya" {
			printStack()
			fmt.Println("I love kitties ")
		}
	}()

	if len(arr) <= 1 {
		return arr
	}else if len(arr) <= 2{
		if arr[0]>arr[1]{
			return []int{arr[1],arr[0]}
		}
		return arr
	}
	return merge(merge_sort(arr[:len(arr)/2]), merge_sort(arr[len(arr)/2:]))
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

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pvt := partition(arr)
	quicksort(arr[:pvt])
	quicksort(arr[pvt:])

	return arr
}

func partition(arr []int) int {
	greaterThanPvt := 0
	pvt := len(arr) - 1 //use the last element
	for i := range arr {
		if arr[i] < arr[pvt] {
			arr[i], arr[greaterThanPvt] = arr[greaterThanPvt], arr[i]
			greaterThanPvt++
		}
	}
	arr[pvt], arr[greaterThanPvt] = arr[greaterThanPvt], arr[pvt]
	return greaterThanPvt
}
