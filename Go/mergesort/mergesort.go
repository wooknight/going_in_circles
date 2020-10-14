package main

import (
	"fmt"
	"sort"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	fmt.Println("hello world")
}

func mergesort(arr1, arr2 []int) (sortedarr []int) {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}
	if len(arr1) == 1 {
		arr2 = append(arr2, arr1[0])
		sort.Ints(arr2)
		return arr2

	}
	if len(arr2) == 1 {
		arr1 = append(arr1, arr2[0])
		sort.Ints(arr1)
		return arr1
	}
	return mergesort(
		mergesort(arr1[:len(arr1)/2], arr1[len(arr1)/2+1:]),
		mergesort(arr2[:len(arr2)/2], arr2[len(arr2)/2+1:]),
	)
}
