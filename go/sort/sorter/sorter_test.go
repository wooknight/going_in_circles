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
	t.Log("Given the need to test quicksort using slices")
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

func TestBubblesort(t *testing.T) {
	testID := 2
	t.Log("Given the need to test bubblesort")
	{
		inp, sorted := getMillion()
		bubblesort(inp)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}

func TestMergesortslice(t *testing.T) {
	testID := 2
	t.Log("Given the need to test slice mergesort")
	{
		inp, sorted := getMillion()
		arr := MergeSortSlice(inp)
		if reflect.DeepEqual(arr, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}

func TestMergesort(t *testing.T) {
	testID := 2
	t.Log("Given the need to test mergesort")
	{
		inp, sorted := getMillion()
		arr := Mergesort(inp, 0, len(inp))
		if reflect.DeepEqual(arr, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}

func TestInsertionsort(t *testing.T) {
	testID := 2
	t.Log("Given the need to test insertion sort")
	{
		inp, sorted := getMillion()
		insertionSort(inp)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing %v to %v", failed, testID, inp , sorted)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}


func TestSelectionsortslice(t *testing.T) {
	testID := 2
	t.Log("Given the need to test sliced selection sort")
	{
		inp, sorted := getMillion()
		selectionSortSlice(inp)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing %v to %v", failed, testID, inp , sorted)
			}
			t.Logf("\t%s\tTest %d:\tcomparing.", success, testID)
		}
	}
}



func TestSelectionsort(t *testing.T) {
	testID := 2
	t.Log("Given the need to test selection sort")
	{
		inp, sorted := getMillion()
		selectionSort(inp)
		if reflect.DeepEqual(inp, sorted) != true {
			{
				t.Fatalf("\t%s\tTest %d:\t comparing %v to %v", failed, testID, inp , sorted)
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
		quicksort(inp, 0, len(inp)-1)
	}
}

func BenchmarkBubblesort(b *testing.B) {
	inp, _ := getMillion()
	for i := 0; i < b.N; i++ {
		bubblesort(inp)
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
