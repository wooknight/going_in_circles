package main

import (
	"reflect"
	"testing"
)

func TestMergesort(t *testing.T) {

	mergeTests := []struct {
		input1 []int
		input2 []int
		result []int
	}{
		{input1: []int{1, 3, 2}, input2: []int{4, 5, 6}, result: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range mergeTests {
		result := mergesort(tt.input1, tt.input2)
		if reflect.DeepEqual(result, tt.result) {
			t.Errorf("Result not matching . Need %v . Got %v", tt.result, result)
		}
	}
}

func TestInsertionSort(t *testing.T){

	insertTests := []struct {
		input1 []int
		result []int
	}{
		{input1: []int{6, 3, 2,5,4, 1}, result: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range insertTests {
		result := insertionSort(tt.input1)
		if reflect.DeepEqual(result, tt.result) {
			t.Errorf("Result not matching . Need %v . Got %v", tt.result, result)
		}
	}

}


//go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
//go tool pprof cpu.prof
func BenchmarkInsertionsort(b *testing.B){
	mergeTests := []struct {
		input1 []int
		result []int
	}{
		{input1: []int{6, 3, 2,4, 5, 1}, result: []int{1, 2, 3, 4, 5, 6}},
	}

		for i := 0; i < b.N; i++ {
		result := insertionSort(mergeTests[0].input1)
		if reflect.DeepEqual(result, mergeTests[0].result) {
			b.Errorf("Result not matching . Need %v . Got %v", mergeTests[0].result, result)
		}
    }
}


//go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
//go tool pprof cpu.prof
func BenchmarkMergesort(b *testing.B){
	mergeTests := []struct {
		input1 []int
		input2 []int
		result []int
	}{
		{input1: []int{1, 3, 2}, input2: []int{4, 5, 6}, result: []int{1, 2, 3, 4, 5, 6}},
	}

	for i := 0; i < b.N; i++ {
		result := mergesort(mergeTests[0].input1, mergeTests[0].input2)
		if reflect.DeepEqual(result, mergeTests[0].result) {
			b.Errorf("Result not matching . Need %v . Got %v", mergeTests[0].result, result)
		}
    }
}