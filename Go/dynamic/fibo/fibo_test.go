package main

import "testing"

var fib int

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

func TestCountsubsetsMemo2D(t *testing.T) {
	fib = countSubsets2DMemo(4, 2)
	if fib != 6 {
		t.Errorf("countSubsets2DMemo(4,2) ! = 6. Got %+v", fib)
	}
}

func BenchmarkCountSubsetsNoMemo(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsets(30, 13)
	}
	fib = val
}

func BenchmarkCountSubsetsMemo(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsetsMemo(30, 13)
	}
	fib = val
}

func BenchmarkCountSubsets2DMemo(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = countSubsets2DMemo(30, 13)
	}
	fib = val
}
