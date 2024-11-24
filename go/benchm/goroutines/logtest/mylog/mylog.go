package mylog

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
)

type Logger struct {
	Writer io.Writer
	Logger log.Logger
}

func (l *Logger) MyLog(s string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		l.Logger.SetPrefix(fmt.Sprintf("called from %s on line %d -> ", filepath.Base(file), line))
		l.Logger.Println(s)
	}
	// {
	// 	_, file, line, ok := runtime.Caller(2)
	// 	if ok {
	// 		l.Logger.SetPrefix(fmt.Sprintf(" that in turn called from %s on line %d -> ", filepath.Base(file), line))
	// 		l.Logger.Println(s)
	// 	}
	// }

}

func New(w io.Writer) *Logger {
	l := Logger{
		Logger: *log.New(w, "Custom log -> ", log.LstdFlags|log.Lshortfile),
	}
	return &l
}
