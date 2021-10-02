package main

import (
	"reflect"
	"testing"
)

//Success and Failure markers
const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestQuiksort(t *testing.T) {
	testID := 0
	t.Log("Given the need to test quicksort")
	{
		inp, sorted := getMillion()
		quicksort(inp)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}

// func TestSelectionSort() {
// 	inp := getMillion()
// 	selection(inp) == sort.Ints(inp)
// }

// func TestInsertionSort() {
// 	inp := getMillion()
// 	insertion(inp) == sort.Ints(inp)
// }

// func TestHeapSort() {
// 	inp := getMillion()
// 	heapsort(inp) == sort.Ints(inp)
// }

// func TestMergeSort() {
// 	inp := getMillion()
// 	mergesort(inp) == sort.Ints(inp)
// }
