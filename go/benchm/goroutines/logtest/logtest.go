package main

import (
	"os"

	"github.com/wooknight/going_in_circles/go/benchm/goroutines/logtest/mylog"
)

func main() {
	l := mylog.New(os.Stdout)
	for i := 0; i < 2; i++ {
		l.MyLog("hello")
	}
	l.MyLog("this is on line 14")
}
