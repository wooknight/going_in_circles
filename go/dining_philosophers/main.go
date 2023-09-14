package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type philosopher struct {
	name      string
	forks     [2]int
	is_eating bool
	done      bool
}

var philosophers = []philosopher{
	{name: "Plato", forks: [2]int{5, 1}},
	{name: "Aristotle", forks: [2]int{1, 2}},
	{name: "Diogenes", forks: [2]int{2, 3}},
	{name: "Socrates", forks: [2]int{3, 4}},
	{name: "Heraclitus", forks: [2]int{4, 5}},
}

func eating(phil *philosopher, fork1, fork2 *sync.Mutex) {
	fmt.Println(" Sat down at the table  ", (*phil).name)
	fork1.Lock()
	fork2.Lock()
	fmt.Println((*phil).name, " started eating ")
	(*phil).is_eating = true
	time.Sleep(2 * time.Second)
	fmt.Println((*phil).name, " stopped eating ")

	fork2.Unlock()
	fork1.Unlock()
	(*phil).done = true
	fmt.Println(*phil, " left the table ")
}

func main() {
	var forks [6]sync.Mutex
	inter := make(chan os.Signal, 1)
	signal.Notify(inter, os.Interrupt)
	for i := 0; i < len(philosophers); i++ {

		go eating(&philosophers[i], &forks[philosophers[i].forks[0]], &forks[philosophers[i].forks[1]])
	}
	<-inter
	for i := 0; i < len(philosophers); i++ {
		if !philosophers[i].done {
			fmt.Println(philosophers[i], " STARVED ")
		}
	}
}
