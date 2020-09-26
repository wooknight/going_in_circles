package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
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
