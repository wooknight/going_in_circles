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
		if r := recover(); r != nil {
			fmt.Printf("\nInside quicksort defer\nFound an exception %v", r)
			fmt.Printf("Arr value is %v \n%v", arr, r)
		}
	}()

	if (len(arr) - 1) > 0 {
		pvt := partition(arr) //last element
		quicksort(arr[:pvt])
		quicksort(arr[pvt:])
	}
}

func partition(arr []int) int {
	greater := 0
	last := len(arr) - 1
	if last <= 0 {
		return greater
	}
	for i := range arr {
		if arr[i] < arr[last] {
			arr[i], arr[greater] = arr[greater], arr[i]
			greater++
		}
	}
	arr[last], arr[greater] = arr[greater], arr[last]

	// if last != greater {
	// 	fmt.Printf("Swapped %d in %d to %d in %d\n\n\n%v\n\n", arr[last], last, arr[greater], greater, arr)
	// }
	return greater
}
