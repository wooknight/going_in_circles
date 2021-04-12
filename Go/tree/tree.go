package main

import (
	"errors"
	"fmt"
	"sort"
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
	// printAllNonRepeats(2)
	// permute([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// printAllSubsets(9)

	subsets([]int{9, 0, 3, 5, 7})
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

func printAllNonRepeats(num int) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	pNRHelper("", arr, num)
}

func pNRHelper(slate string, bag []int, num int) {
	if len(bag) == 0 || len(slate) == num {
		fmt.Println(slate)
		return
	}
	for _, val := range bag {
		stackBag := make([]int, 0)
		//add 0-9 in StackBag if not val and not in slate
		for i := 0; i <= 9; i++ {
			if val != i && !strings.ContainsRune(slate, (rune)(i)+'0') {
				stackBag = append(stackBag, i)
			}
		}
		pNRHelper(slate+strconv.Itoa(val), stackBag, num)
	}
}

func printAllSubsets(num int) [][]int {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := make([][]int, 0)
	return pSubsetsHelper(res, arr, 0, []int{})
}

func pSubsetsHelper(res [][]int, data []int, num int, slate []int) [][]int {
	// fmt.Println(bag, slate)
	if len(data) == num {
		fmt.Println(slate)
		return append(res, slate)
	}
	//exclude case
	res = pSubsetsHelper(res, data, num+1, slate)
	//include case
	// newSlate := append(slate[:0:0], slate...)
	// newSlate = append(newSlate, bag[0])
	// pSubsetsHelper(newSlate, bag[1:], num)
	return pSubsetsHelper(res, data, num+1, append(slate, data[num]))
}

//**IK Template*///

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	return permuteHelper(res, nums, 0, []int{})
}

func permuteHelper(res [][]int, nums []int, idx int, slate []int) [][]int {
	if idx == len(nums) {
		fmt.Println(slate)
		return append(res, slate)
	}
	for i := idx; i < len(nums); i++ {

		nums[idx], nums[i] = nums[i], nums[idx]
		res = permuteHelper(res, nums, idx+1, append(slate, nums[idx]))
		nums[idx], nums[i] = nums[i], nums[idx]

	}
	return res
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = subsetsHelper(nums, 0, []int{}, res)
	return res
}

func subsetsHelper(nums []int, i int, slate []int, res [][]int) [][]int {
	if i == len(nums) {
		sort.Ints(slate)
		res = append(res, slate)
		return res
	}
	backupSlate := make([]int, len(slate)+1)
	copy(backupSlate, slate)
	res = subsetsHelper(nums, i+1, slate, res)
	backupSlate = append(slate, nums[i])
	return subsetsHelper(nums, i+1, backupSlate, res)
}
