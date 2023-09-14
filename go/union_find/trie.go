package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type trie struct {
	root     byte
	prefix   string
	children [26]*trie
}

func main() {

	dat, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal("Could not open file", err)
	}

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanWords)

	root := trie{
		root:   0xff,
		prefix: "",
	}

	//
	for scanner.Scan() {
		addtrie(&root, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read", err)
	}
	printStr(&root)
}

func printStr(root *trie) {
	if root == nil {
		return
	}
	fmt.Println(root.prefix)
	for _, val := range root.children {
		printStr(val)
	}
}

func addtrie(root *trie, word string) {
	if len(word) == 0 {
		return
	}
	val := word[0]
	idx := byte(unicode.ToUpper(rune(val))) - 65
	if root.children[idx] == nil {
		root.children[idx] = &trie{
			root:   val,
			prefix: root.prefix + string(val),
		}
	}
	addtrie(root.children[idx], word[1:])

}
