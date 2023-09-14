package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/wooknight/going_in_circles/go/command_line/chap2/todo"
)

var todoFileName = ".todo.json"

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}

func main() {
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")

	}
	verbose := flag.Bool("verbose", false, "Print all info")
	unfinished := flag.Bool("unfinished", false, "Print all except finished")
	delete := flag.Int("del", -1, "delete i item in the list")

	add := flag.Bool("add", false, "add task to the ToDo List")
	list := flag.Bool("list", false, "list all tasks")
	complete := flag.Int("complete", 0, "item to be completed")
	flag.Parse()
	l := &todo.List{}
	if err := l.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case *delete > 0:
		l.Remove(*delete)
	case *unfinished:
		fmt.Println(l.Verbose())
	case *verbose:
		fmt.Println(l.Unfinished())
	case *list:
		fmt.Print(l)
	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete >= 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}
