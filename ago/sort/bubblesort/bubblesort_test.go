package main

import (
	"testing"
	"reflect"

)

func TestBubbleSort(t *testing.T){
	data:=[]struct {
		input[]int
		sortedData []int
	}{
		{input:[]int {9,8,7,6,5,4,3,2} , sortedData :[]int {2,3,4,5,6,7,8,9}},
	}
	for _,val:= range (data) {
		sortedData:= bubblesort(val.input)
		if !reflect.DeepEqual(sortedData,val.sortedData){
			t.Errorf("Expected - %v . Got - %v\n\n",val.sortedData,sortedData)
		}
	}
}