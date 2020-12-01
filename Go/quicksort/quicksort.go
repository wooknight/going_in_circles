package main

// import "fmt"

func quicksort(arr []int,lower , upper int){
//	fmt.Printf("Inside quicksort - %v\n",arr)
	if (upper - lower) > 0 {
	pvt:=partition(arr,len(arr)-1)//last element
	if pvt > 0 {
		quicksort(arr[:pvt-1],0,pvt-1)
	}
		quicksort(arr[:pvt+1],pvt+1,len(arr)-1)
}

}

func partition(arr []int,idx int) int {
	pvt := 0

	for i:= range arr{
		if arr[i] < arr[idx]{
			arr[i],arr[pvt]=arr[pvt],arr[i]
			pvt++
		}
	}
	arr[idx],arr[pvt]=arr[pvt],arr[idx]
	return pvt
}

