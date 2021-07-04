package main

import (
	"fmt"
	"sync"
)

func cook(recv chan string, wg *sync.WaitGroup) {
	// fmt.Println("Inside goroutine")
	defer wg.Done()
	data, ok := <-recv
	if ok {
		fmt.Println(data)
	} else {
		fmt.Println("Closed")
	}
}

// func pool() {
// 	sem := 0
// 	numCpu := runtime.NumCPU()
// 	for i := 0; i < numCpu; i++ {
// 		sem++
// 		go cook("Sending on CPU " + string(i))
// 	}
// }

func main() {
	var wg sync.WaitGroup
	comm := make(chan string, 1)
	// fmt.Println("Inside main")
	comm <- "I love Scooty"
	wg.Add(1001)
	// fmt.Println("Inside main")
	go cook(comm, &wg)
	// time.Sleep(1000)
	comm <- "Done with Bluesy"
	for i := 0; i < 1000; i++ {
		go cook(comm, &wg)
	}
	close(comm)
	wg.Wait()
}
