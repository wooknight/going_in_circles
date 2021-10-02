package main

import (
	"testing"
)
	
//Success and Failure markers
const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestQuiksort(t *testing.T) {

	t.Log("Given the need to test quicksort")
	{
	inp := getMillion()
	if reflect.DeepEquals( quicksort(inp) , sort.Ints(inp))!= true {
		{
			t.Fatalf("\t%s\tTest %d:\tShould be compare: %v", failed, testID, err)
		}
		t.Logf("\t%s\tTest %d:\tShould be able to compare.", success, testID)
}

func TestSelectionSort() {
	inp := getMillion()
	selection(inp) == sort.Ints(inp)
}

func TestInsertionSort() {
	inp := getMillion()
	insertion(inp) == sort.Ints(inp)
}

func TestHeapSort() {
	inp := getMillion()
	heapsort(inp) == sort.Ints(inp)
}

func TestMergeSort() {
	inp := getMillion()
	mergesort(inp) == sort.Ints(inp)
}
