package ch03

import (
	"fmt"
	"sync"
	"testing"
)

func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)
	queue = append(queue, struct{}{})
	queue = append(queue, struct{}{})
	queue = append(queue, struct{}{})
	queue = append(queue, struct{}{})
	queue = append(queue, struct{}{})
	var m sync.WaitGroup
	m.Add(2)
	removeFromQueue := func(m *sync.WaitGroup) {
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue ", len(queue))
		c.L.Unlock()

		c.Signal()

	}

	waitForSignal := func(p int) {
		defer m.Done()
		fmt.Println("Waiting ", p)
		c.L.Lock()
		c.Wait()
		c.L.Unlock()
		fmt.Println("I am freed ", p)
	}
	go removeFromQueue(&m)

	for i := 0; i < 10; i++ {
		go waitForSignal(i)
	}
	go removeFromQueue(&m)
	m.Wait()
	fmt.Println("Final queue count ", len(queue))
}
