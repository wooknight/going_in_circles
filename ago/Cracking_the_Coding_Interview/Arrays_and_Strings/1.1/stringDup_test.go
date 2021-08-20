package main

import (
	"testing"
)

func TestCheckDupsWithMap(t *testing.T) {
	t.Logf("Given that we are testing for duplicates")
	{
		t.Logf("When we have duplicates")
		{
			str := "abababa"
			isPresent := checkDupsWithMap(str)
			if isPresent {
				t.Logf("We should find duplicates")
			} else {
				t.Errorf("--- Did not find duplicates in %s", str)
			}
		}
		t.Logf("When we have no duplicates")
		{
			str := "abcdefgh"
			isPresent := checkDupsWithMap(str)
			if isPresent {
				t.Errorf("--- We should NOT find duplicates in %s", str)
			} else {
				t.Logf("Did not find duplicates in %s", str)
			}
		}
	}
}

func TestCheckDupsWith2Ptrs(t *testing.T) {
	t.Logf("Given that we are testing for duplicates")
	{
		t.Logf("When we have duplicates")
		{
			str := "abababa"
			isPresent := checkDupsWith2Ptrs(str)
			if isPresent {
				t.Logf("We should find duplicates")
			} else {
				t.Errorf("--- Did not find duplicates in %s", str)
			}
		}
		t.Logf("When we have no duplicates")
		{
			str := "abcdefgh"
			isPresent := checkDupsWithMap(str)
			if isPresent {
				t.Errorf("--- We should NOT find duplicates in %s", str)
			} else {
				t.Logf("Did not find duplicates in %s", str)
			}
		}
	}
}
