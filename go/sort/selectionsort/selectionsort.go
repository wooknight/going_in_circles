package main

import (
	"fmt"

	"math/rand"
)

//generate arrays of random numbers
func getInput (cnt int) []int {
	arr:= make([]int,cnt)
	// r := rand.New(rand.NewSource(10000))
	for i:= 1;i<  cnt;i++ {
		arr[i] =  rand.Intn(10000)
	}
	return arr
}

func main (){
	arr := getInput(100)
	selectionsort(arr)
	fmt.Printf("Ended")
}

func getMin(values []int) (int,int){
	min:= 1000000000000
	idx:=0
	fmt.Printf("%v\n\n\n",values)
	for i,val:= range values{
		if val<min{
			min=val
			idx=i
		}

	}

	return idx,min
}

func minArr(values []int) []int {
	retArr:= make([]int ,len(values))
	for i,_:= range values{
		_,retArr[i]= getMin(values[i:])
	}
	// fmt.Printf("Array min - %v\n\n",retArr)
	return retArr
}

func selectionsort(values []int){
	for i,_ := range values{
		j,_ := getMin(values[i:])
		if (i != (j+1)){
		values[j+i],values[i]=values[i],values[j+i]
		fmt.Printf("%v\n",values)
		}
	}
	fmt.Printf("Array sorted - %v",values)

}