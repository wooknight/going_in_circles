package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var docs []string

func init() {
	rand.Seed(time.Now().UnixNano())
	docs = generateList(1e3)
	fmt.Printf("Processing %d documents using %d goroutines on %d thread(s)\n", len(docs), runtime.NumCPU(), runtime.GOMAXPROCS(0))
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "test", docs)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "test", docs)
	}
}
