package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diskUsage(dir string, duInfo os.FileInfo) int {
	var du int
	if duInfo.IsDir() == false {
		return int(duInfo.Size())
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		du += diskUsage(dir+"/"+file.Name()+"/", file)
	}
	fmt.Printf("Size of directory %s - %d\n", dir, du)
	return du
}

func sortedArrayToBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	} 
    mid := (len(arr)) / 2
    t := &TreeNode{Val:arr[mid]}
   
	if mid > 0 {
		t.Left = sortedArrayToBST(arr[:mid])
	} 
    if mid < len(arr)    {
        t.Right = sortedArrayToBST(arr[mid+1:])
	}
    
    return t

}

func main() {
	dir := "/Users/ramesh/Documents/Repos/"
	fmt.Println("Inside ", dir)
	duInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Error in ", dir, err)
		log.Fatal(err)
	}
	fmt.Println(dir, diskUsage(dir, duInfo))
}
