package main

import (
	"fmt"
	"math/rand"
)

func getInput (len int) []int{
	arr := make([]int, len)
	for i:=0;i<len;i++{
		arr[i] = rand.Intn(len*100)
	}
	return arr
}


func bubblesort(arr []int) []int{
	arrLen:=len(arr)-1
	for arrLen >0 {
		for i:=0;i<arrLen;i++{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
			}
		}
		arrLen--
	}
	return arr
}

func main(){
	arr := getInput(1000)
	sortedarr:= bubblesort(arr)
	fmt.Printf("created array - %v ; sorted array - %v",arr,sortedarr)
}