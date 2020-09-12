package main

import (
	"fmt"
)


type node struct {
	str         string
	value       int
	left, right *node
}

type helper struct {
	sum int
	cnt int
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

func initial(){
	root = &node {str :"Root", value : 10, left:nil,right:nil}
	root.left = &node {str :"left", value : 5, left:nil,right:nil}
	root.right = &node {str :"right", value : 15, left:nil,right:nil}
	root.right.left = &node {str :"right-left", value : 95, left:nil,right:nil}
	root.right.right = &node {str :"right-right", value : 35, left:nil,right:nil}
	root.left.left = &node {str :"left-left", value : 95, left:nil,right:nil}
	root.left.right = &node {str :"left-right", value : 35, left:nil,right:nil}

}

func main (){
	channels := make(map[int]helper)
	initial()
	
	traversal(root,level,channels)
	fmt.Printf("Channels %v",channels)
}