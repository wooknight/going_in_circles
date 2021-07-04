package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

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

func fanoutSemaphore() {

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

func fanoutBounded() {
	work := []string{"papel", "papel", 2000: "paper"}
	g := runtime.NumCPU()
	ch := make(chan string, g)
	wg.Add(g)
	for emp := 0; emp < g; emp++ {
		go func(emp int, wg *sync.WaitGroup) {
			defer wg.Done()
			for p := range ch {
				fmt.Printf("employee %d: recv signal : %s \n", emp, p)
			}
			fmt.Printf("Received shutdown signal : emp %d\n", emp)
		}(emp, &wg)
	}
	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()
}

func drop() {
	work := []string{"papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", "papel", 2000: "paper"}
	ch := make(chan string)
	go func() {
		for p := range ch {
			fmt.Println("Done work ", p)
		}
		fmt.Println("Ending because channel closed")
	}()
	for _, val := range work {
		select {
		case ch <- val:
			fmt.Println("Manager snet data ", val)
		default:
			fmt.Println("Manager dropped data", val)
		}
	}
	close(ch)
}
func main() {
	// fanoutSemaphore()
	// fanoutBounded()
	drop()
}
