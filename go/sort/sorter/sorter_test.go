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

func TestQuiksortSlice(t *testing.T) {
	testID := 0
	t.Log("Given the need to test quicksort")
	{
		inp, sorted := getMillion()
		quicksortSlice(inp)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}


func TestQuiksort(t *testing.T) {
	testID := 1
	t.Log("Given the need to test quicksort")
	{
		inp, sorted := getMillion()
		quicksort(inp, 0, len(inp)-1)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}

func BenchmarkQuicksortSlice(b *testing.B) {
	inp, _ := getMillion()
	for i := 0; i < b.N; i++ {
		quicksortSlice(inp)
	}
}
func BenchmarkQuicksort(b *testing.B) {
	inp, _ := getMillion()
	for i := 0; i < b.N; i++ {
		quicksort(inp,0,len(inp)-1)
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
