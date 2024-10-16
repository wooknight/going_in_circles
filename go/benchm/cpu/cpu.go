package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateList(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}

func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}
func addConcurrentChannel(goroutines int, numbers []int) int {
	var v int64
	sum := func(from, to int, ch chan<- int) {
		var s int
		for _, n := range numbers[from:to] {
			s += n
		}
		ch <- s
	}
	ch := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		g := i * len(numbers) / goroutines
		h := (i + 1) * len(numbers) / goroutines
		go sum(g, h, ch)
	}
	for i := 0; i < goroutines; i++ {
		v += int64(<-ch)
	}
	return int(v)
}
func addConcurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoRoutine := goroutines - 1 // last go routine takes care of the rest
	stride := totalNumbers / goroutines
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoRoutine {
				end = totalNumbers
			}
			var s int
			for _, n := range numbers[start:end] {
				s += n
			}
			atomic.AddInt64(&v, int64(s))
			wg.Done()
		}(g)
	}
	return int(v)
}

func main() {
	numbers := generateList(1e7)
	fmt.Println(add(numbers))
	fmt.Println(addConcurrent(runtime.NumCPU(), numbers))
}
