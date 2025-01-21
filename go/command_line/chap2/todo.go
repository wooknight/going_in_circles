package main

import (
	"bufio"
	"chap2/list"
	"fmt"
	"os"
	"strings"
	"time"
)

var todos list.List
var filename = "todos.txt"

func main() {
	fmt.Fprintln(os.Stdout, "Welcome to the todo app")
	for {
		command := bufio.NewScanner(os.Stdin)
		command.Scan()
		inp := strings.Split(command.Text(), " ")
		if len(inp) == 0 {
			continue
		}
		switch inp[0] {
		case "add":
			fmt.Println("adding", inp[1:])
			todos.Add(list.Todo{Title: strings.Join(inp[1:], " "), Done: false, CreatedAt: time.Now().Format(time.RFC3339)})
		case "list":
			fmt.Fprintln(os.Stdout, "listing", inp[1:])
			fmt.Fprintln(os.Stdout, todos.List())
		case "delete":
			fmt.Fprintln(os.Stdout, "deleting", inp[1:])
			todos.DeleteByTitle(strings.Join(inp[1:], " "))
		case "done":
			fmt.Fprintln(os.Stdout, "marking", inp[1:], "done")
			todos.CompleteByTitle(strings.Join(inp[1:], " "))
		case "quit":
			fmt.Fprintln(os.Stdout, "quitting", inp[1:])
			return
		case "write":
			todos.Save(filename)
		case "read":
			todos.Load(filename)

		default:
			fmt.Println("invalid command", inp[0])
		}
	}
}
