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
	if last <=0 {return greater}
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


func binarySearch(arr []int, target int) int{
	mid:=len(arr)/2
	low:=0
	high := len(arr)-1
	for low<=high{
		if arr[low]==target {return low}
		if arr[high]==target {return high}
		if arr[mid]==target {return mid}
		if (arr[mid]<target){
			low=mid+1
		}else{
			high=mid-1
		}
		mid = low + (high - low)/2
	}
	return -1
}

func selectionSort(arr []int){
	for i,_ := range arr{
		minIdx,_:=findMin(arr[i:])
		arr[i],arr[i+minIdx]=arr[minIdx+i],arr[i]
	}
}

func findMin(arr []int) (int,int){
	idx:=0
	val:=arr[0]
	for idx1,val1:=range arr{
		if val > val1{
			idx=idx1
			val=val1
		}
	}
	// fmt.Printf("Found %d in array %v\n\n",val,arr)
	return idx,val
}

