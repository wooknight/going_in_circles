package main

import "fmt"

func main(){
	fmt.Printf("Choot ki phuddi\n")
	vars := []int{}
	for i:= 100;i>= 20;i--{
		vars=append(vars,i)
	}
	recursiveSelectionSort(vars)
	fmt.Printf("Sorted array %v\n\n",vars)

}

func permute(vars []int , k int){
	fmt.Printf("%v\t%d\n",vars,k)
	if k >= len(vars)-1 {
		fmt.Printf("%d \n ", vars)
	}else{
		for i := range (vars){
			if i != k{
				vars[i],vars[k]=vars[k],vars[i]
				permute(vars[i:],k+1)
			}
		}
	}
}

func recursiveSelectionSort(arr []int)  {
	if len(arr) <= 0 {
		return
	}
	min:=arr[0]
	start:=0
	for i,val:= range arr{
		if val < min{
			start = i
			min = val
		}
	}
	arr[0],arr[start]=arr[start],arr[0]
	recursiveSelectionSort(arr[1:])
}