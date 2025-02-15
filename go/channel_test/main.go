package main

import "fmt"

func main() {

	testFunc := func() (chan int, func()) {
		ch := make(chan int)
		return ch, func() {
			close(ch)
		}
	}

	ch, cancel := testFunc()
	go func(ch chan int, cancel func()) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		cancel()
	}(ch, cancel)
	go func() {
		for i := range ch {
			println(i)
		}
	}()
	// <-ch
	fmt.Println("Program ended")
}
