package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go/command_line/chap3/mdp/markdown"
)

func main() {
	file := flag.String("file", "", "markdown file")
	flag.Parse()
	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *file != "" {
		if err := run(*file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func run(file string) error {
	md := &markdown.MarkDown{}
	dat, err := md.Open(file)
	if err != nil {
		return err
	}
	md.ParseContent(dat)
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	fmt.Println(temp.Name())
	return md.RenderHTML(temp.Name())
}
