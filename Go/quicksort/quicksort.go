package main

import (
	"fmt"
	"os"
	"runtime"
)

func printStack() {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func quicksort(arr []int) {
	defer func() {
		fmt.Println("Inside quicksort defer")

		if r := recover(); r != nil {
			fmt.Printf("Arr value is %v \n%v", arr, r)
		}
	}()

	if (len(arr) - 1) > 0 {
		pvt := partition(arr) //last element
		if pvt > 0 {
			quicksort(arr[:pvt])
		}
		if pvt < (len(arr) - 1) {
			quicksort(arr[:pvt+1])
		}
	}
}

func partition(arr []int) int {
	pvt := 0
	idx := len(arr) - 1
	for i := range arr {
		if arr[i] < arr[idx] {
			arr[i], arr[pvt] = arr[pvt], arr[i]
			pvt++
		}
	}
	arr[idx], arr[pvt] = arr[pvt], arr[idx]

	if idx != pvt {
		fmt.Printf("Swapped %d in %d to %d in %d\n\n\n%v\n\n", arr[idx], idx, arr[pvt], pvt, arr)
	}
	return pvt
}
