package main

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestConvertSortedArrToBST(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := sortedArrayToBST(arr)

	t.Log("Given the need to test sorting.")
	{
		t.Logf("\tTest 0:\tWhen checking  for sorted item in root")
		{
			if root.Val != 6 {
				t.Fatalf("\t%s\tChecking for Root matching %d", failed, root.Val)
			}
			t.Logf("\t%s\tChecking for root matching", succeed)
		}
	}
	if root.Left.Val != 3 {
		t.Errorf("Root not matching %d", root.Left.Val)
	}

}
