package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/wooknight/going_in_circles/go/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\nHello %s,This is the Monkey interpreter\n", user)
	fmt.Printf("Start typing commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
