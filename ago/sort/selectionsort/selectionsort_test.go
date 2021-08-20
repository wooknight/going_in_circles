package main

import (
	"fmt"
	"reflect"
	"testing"
)


func TestSelectionsort(t *testing.T){
	getmins := []struct {
		input []int
		sortedOutput []int
	}{
		{input: []int{0,2,1,8,4} ,sortedOutput: []int{0,1,2,4,8}},
		{input:[]int{12,4,6,18,2} ,sortedOutput: []int{2,4,6,12,18}},

	}
	for _,val:= range getmins {
		arr := val.input
		selectionsort(arr)
		if !reflect.DeepEqual(val.input, val.sortedOutput) {
			t.Errorf("Expected %v ; got %v",val.sortedOutput,arr)
		}
	}

}

func TestGetMin(t *testing.T){
	getmins := []struct {
		haystack []int
		mini int
	}{
		{haystack: []int {1,11,2,3,14,4,21,28,}, mini:1},
		{haystack: []int {21,121,4,78,141,214,231,28,}, mini:4},
	}
	for _,val := range getmins {
		idx, value := getMin(val.haystack)
		// fmt.Printf("\n\nGot from min - %d idx and %d val\n\n",idx, val.haystack[idx])
		if val.mini != val.haystack[idx] {
			t.Errorf("\n\nExpected %d - Got %d\n\n",val.mini,value)
		}
	}
}

func TestMinArr(t *testing.T){
	getArrs:= []struct {
		input []int
		sortedMins []int
	}{
		{input: []int{3,1,4,6,7,8,2},sortedMins: []int{1,1,2,2,2,2,2}},
		{input :[]int{1,2,3,4,5,6,7,8}, sortedMins: []int{1,2,3,4,5,6,7,8} },
	}
	for _,val:= range getArrs {
		sortedVal := minArr(val.input)
		fmt.Printf("\nsorted value - %v - expectedpresorted value - %v\n",sortedVal,val.sortedMins)
		if !reflect.DeepEqual(sortedVal,val.sortedMins){
			t.Errorf("Expected %v - Got %v",val.sortedMins,sortedVal)
		}
	}
}