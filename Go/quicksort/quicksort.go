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
	fmt.Printf("Found %d in array %v\n\n",val,arr)
	return idx,val
}