package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//binary search tree
type tree struct {
	val   int
	left  *tree
	right *tree
}

func main() {

	// printAllRepeats(7)
	printAllNonRepeats()

}

func searchIterative(root *tree, target int) (*tree, error) {
	if root == nil {
		return nil, nil
	}
	if root.val == target {
		return root, nil
	}
	curr := root
	for curr != nil {
		if curr.val < target {
			//right child
			curr = curr.right
		} else {
			//left tree
			curr = curr.right
		}
	}
	return nil, errors.New("Could not find target")
}

func searchRecursive(root *tree, target int) (*tree, error) {
	if root == nil {
		return nil, nil
	}
	if root.val == target {
		return root, nil
	}
	curr := root
	if curr.val < target {
		//right child
		return searchRecursive(curr.right, target)
	} else if curr.val > target {
		//left tree
		return searchRecursive(curr.left, target)
	}
	return nil, nil
}

// func insert(root *treeval *tree){

// }

// func hashTable (){

// }

func printAllRepeats(num int) {
	phelper("", num)
}

func phelper(slate string, num int) {
	if num == 0 {
		fmt.Println(slate)
		return
	}
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range arr {
		phelper(slate+strconv.Itoa(val), num-1)
	}
}

func printAllNonRepeats() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	pNRHelper("", arr)
}

func pNRHelper(slate string, bag []int) {
	if len(bag) == 0 {
		fmt.Println(slate)
		return
	}
	for _, val := range bag {
		stackBag := make([]int,0)
		//add 0-9 in StackBag if not val and not in slate
		for i:=0;i<=9;i++{
			if val!=i && !strings.ContainsRune(slate, (rune)(i) + '0'){
				stackBag = append(stackBag,i)
			}  
		}
		pNRHelper(slate+strconv.Itoa(val), stackBag)
		}
	}
