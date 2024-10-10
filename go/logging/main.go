package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func NewLog(w io.Writer, cap int) *Logger {
	lg := Logger{
		ch: make(chan string, cap),
	}
	lg.wg.Add(1)
	go func() {
		defer lg.wg.Done()
		for msg := range lg.ch {
			fmt.Fprintln(w, msg)
		}
	}()
	return &lg
}
func (l *Logger) Write(p []byte) (n int, err error) {
	l.ch <- string(p)
	return len(p), nil
}

func (l *Logger) Shutdown() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(v string) {
	select {
	case l.ch <- v:
		return
	default:
		fmt.Println("Dropping log message")
	}

}

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
	l := NewLog(&d, numCPU)
	for i := 0; i < numCPU; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("log something :%d", id))
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
		//ignoreing data race issues here
		d.problem = !d.problem
	}
}
