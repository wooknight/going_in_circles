package main

import (
	"errors"
	"fmt"
)

//binary search tree
type tree struct {
	val   int
	left  *tree
	right *tree
}

func main() {
	fmt.Printf("Printing treee")

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
}

func insert(root *treeval *tree){

}

func hashTable (){
	
}