package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	log "GoingInCircles/Go/benchm/goroutines/logger/logge"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (in int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}
	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	const grs = 5
	var d device
	l := log.New(&d, 5)
	//500 go routines writing to disk/ screen
	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d : log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	for {
		<-sigchan
		d.problem = !d.problem //toggle
	}

}
