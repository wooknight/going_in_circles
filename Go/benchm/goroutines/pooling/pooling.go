package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func cook(recv chan string, wg *sync.WaitGroup, num int) {
	// fmt.Println("Inside goroutine")
	defer wg.Done()
	data, ok := <-recv
	if ok {
		fmt.Println(data)
	} else {
		fmt.Println("Closed" + strconv.Itoa(num))
	}
}

func main() {
	var wg sync.WaitGroup
	comm := make(chan string)
	// fmt.Println("Inside main")
	wg.Add(1001)
	// fmt.Println("Inside main")
	go cook(comm, &wg, 0)
	// time.Sleep(1000)
	comm <- "I love Scooty"
	for i := 0; i < 1000; i++ {
		go cook(comm, &wg, i)
		if i < 700 {
			comm <- "Done with " + strconv.Itoa(i)
		}
	}
	close(comm)
	wg.Wait()
	fmt.Println("Starting anew")
	ch := make(chan string)
	numCpu := runtime.NumCPU()
	wg.Add(numCpu)
	for i := 0; i < numCpu; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Inside pool goroutine")
			for p := range ch {
				fmt.Println("Got data ", p)
			}
			fmt.Println("Shutting down")
		}(&wg)
	}
	ch <- "Done with " + strconv.Itoa(numCpu)
	close(ch)
	wg.Wait()
}
