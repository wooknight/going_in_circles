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
	emps := 1000
	comm := make(chan string)
	// fmt.Println("Inside main")
	wg.Add(emps + 1)
	// fmt.Println("Inside main")
	go cook(comm, &wg, 0)
	// time.Sleep(1000)
	comm <- "I love Scooty"
	numCpu := runtime.NumCPU()
	semaphores := make(chan bool, numCpu)
	for i := 0; i < emps; i++ {
		go func(wg *sync.WaitGroup, num int) {
			defer wg.Done()
			semaphores <- true
			{
				comm <- "sending from go routine" + strconv.Itoa(num)

			}
			<-semaphores
		}(&wg, i)
	}

	for emps > 0 {
		p := <-comm
		fmt.Println(p)
		emps--
	}
	close(comm)

	wg.Wait()

}
