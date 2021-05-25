package main

import (
	"testing"
)

func TestConvertSortedArrToBST(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	root := sortedArrayToBST(arr)
	if root.Val != 6 {
		t.Errorf("Root not matching %d", root.Val)
	}
	if root.Left.Val != 3 {
		t.Errorf("Root not matching %d", root.Left.Val)
	}

}
