package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"
)

type device struct {
	problem bool
}

// writing the IO writer interface
func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}
	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	numCPU := runtime.NumCPU()
	var d device
	l := log.New(&d, "prefix", 0)
	for i := 0; i < numCPU; i++ {
		go func(id int) {
			for {
				l.Println("log something ", id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	//simulate a problem
	// capture interrrupt signal, e.g., Ctrl+C to toggle the problem
	// use ctrl + z to stop the program
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	for {
		<-sigChan
		d.problem = !d.problem
	}
}
