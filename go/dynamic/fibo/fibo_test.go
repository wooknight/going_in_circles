package main

import (
	"testing"
)

var fib int

//go test -bench . -benchtime 3s
func BenchmarkMemofibo(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = memofibo(30)
	}
	fib = n
}

func BenchmarkFibo(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = fibo(30)
	}
	fib = n
}

func TestMin(t *testing.T) {
	x := min(2, 3)
	if x != 2 {
		t.Errorf("min (2,3) failed . Got %d", x)
	}
}

func TestMinChange(t *testing.T) {
	tests := []struct {
		amount    int
		coinArray []int
		result    int
	}{
		{amount: 10, coinArray: []int{1, 5, 7}, result: 2},
		{amount: 10000, coinArray: []int{1, 5, 10, 25}, result: 400},
		{amount: 1000000, coinArray: []int{1, 5, 7}, result: 142858},
	}
	for _, tc := range tests {
		x := coinMinChange(tc.amount, tc.coinArray)
		if x != tc.result {
			t.Errorf("Expected - %d for Input %+v for amount %d. Got %d", tc.result, tc.coinArray, tc.amount, x)
		}
	}
}

func TestMinCost(t *testing.T) {
	x := minCost(10, []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	if x != 6 {
		t.Errorf("minCost(10,[]int{1,100,1,1,1,100,1,1,100,1}) failed. Expected 6 . Got %d", x)
	}

}

func TestCountsubsets(t *testing.T) {
	fib = countSubsets(4, 2)
	if fib != 6 {
		t.Errorf("countSubsets(4,2) ! = 6. Got %+v", fib)
	}
}

func TestCountsubsetsMemo(t *testing.T) {
	fib = countSubsetsMemo(4, 2)
	if fib != 6 {
		t.Errorf("countSubsetsMemo(4,2) ! = 6. Got %+v", fib)
	}
}

func TestCountpaths(t *testing.T) {
	res := countPaths(3, 7)
	if res != 28 {
		t.Errorf("countPaths(3,7) was not 28 . Got %+v", res)
	}
	// fmt.Println(countPaths(3, 3))
	// fmt.Println(countPaths(5, 5))
}

func TestCountsubsetsPregenMemo2D(t *testing.T) {
	fib = countSubsetsPregen2DMemo(4, 2)
	if fib != 6 {
		t.Errorf("countSubsetsPregen2DMemo(4,2) ! = 6. Got %+v", fib)
	}
}

func BenchmarkCountSubsetsNoMemo(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsets(40, 13)
	}
	fib = val
}

func BenchmarkCountSubsetsMemo(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsetsMemo(40, 13)
	}
	fib = val
}

func BenchmarkCountSubsetsMemoMap(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsetsMemoMap(40, 13)
	}
	fib = val
}

func BenchmarkCountSubsetsPregen2DMemo(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsetsPregen2DMemo(40, 13)
	}
	fib = val
}
