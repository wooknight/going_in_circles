package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	Title     string
	Done      bool
	CreatedAt string
}

var todos []Todo

func main() {
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
			todos = append(todos, Todo{Title: strings.Join(inp[1:], " "), Done: false, CreatedAt: time.Now().Format(time.RFC3339)})
		case "list":
			fmt.Println("listing", inp[1:])
			for _, todo := range todos {
				fmt.Println(todo)
			}
		case "delete":
			fmt.Println("deleting", inp[1:])
			for idx, todo := range todos {
				cmp := strings.Join(inp[1:], " ")
				if todo.Title == cmp {
					if idx == len(todos)-1 {
						todos = todos[:idx]
					} else {
						todos = append(todos[:idx], todos[idx+1:]...)
					}
				}
			}
		case "done":
			fmt.Println("marking", inp[1:], "done")
			for idx, todo := range todos {
				cmp := strings.Join(inp[1:], " ")
				if todo.Title == cmp {
					todos[idx].Done = true
				}
			}
		case "quit":
			fmt.Println("quitting", inp[1:])
			os.Exit(0)
		case "write":
			fp, err := os.Create("todos.txt")
			if err != nil {
				fmt.Println("error opening file", err)
				return
			}
			enc := json.NewEncoder(fp)
			err = enc.Encode(todos)
			if err != nil {
				fmt.Println("error getting json encoder ", err)
				return
			}
			fp.Close()
		case "read":
			fp, err := os.Open("todos.txt")
			if err != nil {
				fmt.Println("error opening file", err)
				return
			}
			dec := json.NewDecoder(fp)
			err = dec.Decode(&todos)
			if err != nil {
				fmt.Println("error getting json decoder ", err)
				return
			}
			fp.Close()
		default:
			fmt.Println("invalid command")
		}
	}
}
