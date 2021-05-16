package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	inFile, err := os.Open("memory.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	in := bufio.NewReader(inFile)
	outFile, err := os.Create("Backup_memory.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	out := bufio.NewWriter(outFile)

	echo(in, out)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
