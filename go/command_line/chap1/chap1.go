package main

import (
	"bufio"
	"io"
)

const (
	Bytes int = iota
	Runes
	Words
	Lines
)

func count(r io.Reader, parseBy int) int {
	var count int
	scanner := bufio.NewScanner(r)
	switch parseBy {
	case Bytes:
		scanner.Split(bufio.ScanBytes)
	case Runes:
		scanner.Split(bufio.ScanRunes)
	case Words:
		scanner.Split(bufio.ScanWords)
	case Lines:
		scanner.Split(bufio.ScanLines)
	}
	for scanner.Scan() {
		count++
	}
	return count
}
