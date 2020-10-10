package main

import (
	"fmt"
	"io/ioutil"
)

const SIZE = 15

type hashNode struct {
	Value int
	Next *hashNode
}

type node struct {
	str         string
	value       int
	left, right *node
}

type helper struct {
	sum int
	cnt int
}

type hashTable struct {
	Table map[int]*hashNode
	Size int
}

var root *node
var level int
func traversal(root *node, level int, channels map[int]helper ) {
		if root == nil {
			return
		}
		formatter:=""
		for i := 0;i<level;i++{
			formatter=formatter+"\t"
		}
		fmt.Printf("%s Level %d Val %s %d\n",formatter , level, root.str,root.value)
		if _, ok := channels[level]; ok {
			channels[level] = helper{channels[level].sum + root.value, channels[level].cnt + 1}

		}else{
			channels[level] = helper{sum: root.value,cnt : 1}
		}
		traversal(root.left,level+1,channels)
		traversal( root.right,level+1,channels)

}

func hashFunction(i,size int)int {
return (i%size)
}

func insert ()int {
	index := hashFunction(value,hash.Size)
	element := hashNode (Value:value , Next : hash.Table(index))
	hash.Table[index] = &element
	return index
}

func traverse(hash *hashTable){
	for k := range hash.Table{
		if hash.Table(k) != nil{
			t:=hash.Table(k)
			for t:= nil{
				fmt.Printf("%d -> ",t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}
func hashTraverse(){
	table := make(map[int] *Node,SIZE)
	hash := &HashTable(Table:table ,Size : SIZE)
	fmt.Println("Number of spaces :",hash.Size)
	for i := 0; i < 120; i++{
		insert(hash , i )
	}
	traverse(hash)
}

func initial(){
	root = &node {str :"Root", value : 10, left:nil,right:nil}
	root.left = &node {str :"left", value : 5, left:nil,right:nil}
	root.right = &node {str :"right", value : 15, left:nil,right:nil}
	root.right.left = &node {str :"right-left", value : 95, left:nil,right:nil}
	root.right.right = &node {str :"right-right", value : 35, left:nil,right:nil}
	root.left.left = &node {str :"left-left", value : 95, left:nil,right:nil}
	root.left.right = &node {str :"left-right", value : 35, left:nil,right:nil}

}
func createTree(){
	channels := make(map[int]helper)
	initial()

	traversal(root,level,channels)
	fmt.Printf("Channels %v\n\n",channels)
	for val,help := range channels{
		fmt.Printf("Avg for Level %d is %d \n",val,help.sum/help.cnt)

	}
}

func main (){
//createTree()
hashTraverse()
}