package log

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}
	l.wg.Add(1)
	go func() {
		l.wg.Done()
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}
	}()
	return &l
}

func (l *Logger) Shutdown() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(s string) {
	select {
	case l.ch <- s:
	default:
		fmt.Println("DROP")
	}
}
