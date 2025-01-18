package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	log.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	cnt := 0
	bufScan := bufio.NewScanner(r)
	bufScan.Split(bufio.ScanWords)
	for bufScan.Scan() {
		cnt++
	}
	return cnt
}
