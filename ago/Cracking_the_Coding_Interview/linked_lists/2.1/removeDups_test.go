package main

import (
	"testing"
)

func TestDeDup(t *testing.T) {
	t.Logf("Given that we are testing for duplicates")
	{
		t.Logf("When we have duplicates")
		{
			retList := deDup(createList([]int{1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1}))
			if compareTwoLists(retList, createList([]int{1, 2})) {
				t.Logf("We have removed duplicates")
			} else {
				t.Errorf("--- failed to remove dups ")
			}
		}
		t.Logf("When we have not have duplicates")
		{
			retList := deDup(createList([]int{1, 2, 3, 7, 9, 11, 12, 121, 21, 41, 18}))
			if compareTwoLists(retList, createList([]int{1, 2, 3, 7, 9, 11, 12, 121, 21, 41, 18})) {
				t.Logf("We have compared non duplicates succesfully")
			} else {
				t.Errorf("--- failed to remove dups ")
			}
		}
		t.Logf("When we have have duplicates but check for wrong items")
		{
			retList := deDup(createList([]int{1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1}))
			if !compareTwoLists(retList, createList([]int{1, 3})) {
				t.Logf("We have compared duplicates succesfully and failed the neggy test")
			} else {
				t.Errorf("--- WHY DID THIS PASS ")
			}
		}
	}
}
