package main

import (
	"strings"
	"testing"
)

const success = "\u2713"
const failed = "\u2717"

func TestCount(t *testing.T) {
	t.Log("Given that we are testing byte count")
	{
		{
			t.Log("When we are testing for valid byte count")
			res := count(strings.NewReader("I am a king"), 0)
			if res != 11 {
				t.Fatalf("%s : Expected Byte count was 11 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Byte count was 11", success)
		}
		{
			t.Log("When we are testing for valid byte count")
			res := count(strings.NewReader(""), 0)
			if res != 0 {
				t.Fatalf("%s : Expected Byte count was 0 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Byte count was 0", success)
		}
	}
	t.Log("Given that we are testing rune count")
	{
		{
			t.Log("When we are testing for valid Rune count")
			res := count(strings.NewReader("I am a king"), 1)
			if res != 11 {
				t.Fatalf("%s : Expected Rune count was 11 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Rune count was 11", success)
		}
		{
			t.Log("When we are testing for valid Rune count")
			res := count(strings.NewReader(""), 0)
			if res != 0 {
				t.Fatalf("%s : Expected Rune count was 0 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Rune count was 0", success)
		}
	}
	t.Log("Given that we are testing Word count")
	{
		{
			t.Log("When we are testing for valid Word count")
			res := count(strings.NewReader("I am a king"), 2)
			if res != 4 {
				t.Fatalf("%s : Expected Word count was 4 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Word count was 4", success)
		}
		{
			t.Log("When we are testing for valid Word count")
			res := count(strings.NewReader(""), 0)
			if res != 0 {
				t.Fatalf("%s : Expected Word count was 0 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Word count was 0", success)
		}
	}
	t.Log("Given that we are testing Line count")
	{
		{
			t.Log("When we are testing for valid Line count")
			res := count(strings.NewReader("I am a king"), 3)
			if res != 1 {
				t.Fatalf("%s : Expected Line count was 4 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Line count was 4", success)
		}
		{
			t.Log("When we are testing for valid Line count")
			res := count(strings.NewReader(""), 0)
			if res != 0 {
				t.Fatalf("%s : Expected Line count was 0 . Got %d ", failed, res)
			}
			t.Logf("%s :Expected Line count was 0", success)
		}
	}
}
