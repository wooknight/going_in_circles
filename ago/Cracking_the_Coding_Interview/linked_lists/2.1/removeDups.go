package main

import (
	"fmt"
	"os"
)

type LinkedList struct {
	val  int
	next *LinkedList
}

func main() {
	pth, _ := os.Getwd()
	fmt.Printf("chootiya - %s", pth)

}

func createList(arr []int) *LinkedList {
	head := &LinkedList{val: 0, next: nil}
	start := head
	for _, value := range arr {
		head.next = &LinkedList{val: value, next: nil}
		head = head.next
	}
	return start.next
}

func deDup(root *LinkedList) *LinkedList {
	head1 := root;
	for  head1 != nil{
		ptr1:=head1
		head2 := head1.next;
		for  head2 != nil {
			if head1.val == head2.val {
				ptr1.next = head2.next
				if head2 == nil {
					break
				}
			}else{
				ptr1 = ptr1.next
			}
			head2 = head2.next 
		}
		head1 = head1.next 
	}
	return root
}

func printList(head *LinkedList) {
	for root := head; root != nil; root = root.next {
		fmt.Printf("Val = > %d\n", root.val)
	}
}

func compareTwoLists(list1, list2 *LinkedList) bool {
	head1 := list1
	head2 := list2
	for head1 != nil && head2 != nil {
		if head1.val != head2.val {
			fmt.Printf("Head 1 Val = > %d\tHead 2 Val = > %d\n", head1.val, head2.val)
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}
	if head1 != nil || head2 != nil {
		fmt.Printf("found nil . Head 1 Val = > %v\tHead 2 Val = > %v\n", head1, head2)
		return false
	}
	return true
}
