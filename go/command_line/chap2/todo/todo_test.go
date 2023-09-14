package todo

import (
	"os"
	"testing"
)

const success = "\u2713"
const failed = "\u2717"

func TestAdd(t *testing.T) {
	t.Log("We are testing add")
	{
		t.Log("Adding an item")
		{
			l := NewList()
			l.Add("Testing")
			if l.Count() != 1 {
				t.Fatalf("%s : Count should be 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 1", success)
		}
	}
}

func TestRemove(t *testing.T) {
	t.Log("We are testing remove")
	{
		t.Log("We are testing a successful remove")
		{
			l := NewList()
			l.Add("Testing")
			l.Remove(0)
			if l.Count() != 0 {
				t.Fatalf("%s : Count should be 0. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 0", success)
		}
		t.Log("We are testing a invalid remove")
		{
			l := NewList()
			l.Add("Testing")
			err := l.Remove(-1)
			if err != outOfBoundsError {
				t.Fatalf("%s : Expected out of bounds error. Got %v", failed, err)
			}
			t.Logf("%s : Expected out of bounds error.", success)
			if l.Count() != 1 {
				t.Fatalf("%s : Count should be 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 1.", success)
			err = l.Remove(5)
			if err != outOfBoundsError {
				t.Fatalf("%s : Expected out of bounds error. Got %v", failed, err)
			}
			t.Logf("%s : Expected out of bounds error.", success)
			if l.Count() != 1 {
				t.Fatalf("%s : Count should be 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 1.", success)

		}

	}
}

func TestCompleted(t *testing.T) {
	t.Log("We are testing Completed")
	{
		t.Log("We are testing a successful Completed")
		{
			l := NewList()
			l.Add("Testing")
			l.Complete(0)
			if (*l)[0].Completed != true {
				t.Fatalf("%s : Complete should be set to 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Complete should be set to 1.", success)
			if l.Count() != 1 {
				t.Fatalf("%s : Count should be 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 1", success)
		}
		t.Log("We are testing a invalid Completed")
		{
			l := NewList()
			l.Add("Testing")
			err := l.Complete(-1)
			if err != outOfBoundsError {
				t.Fatalf("%s : Expected out of bounds error. Got %v", failed, err)
			}
			t.Logf("%s : Expected out of bounds error.", success)
			if l.Count() != 1 {
				t.Fatalf("%s : Count should be 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 1.", success)
			err = l.Complete(5)
			if err != outOfBoundsError {
				t.Fatalf("%s : Expected out of bounds error. Got %v", failed, err)
			}
			t.Logf("%s : Expected out of bounds error.", success)
			if l.Count() != 1 {
				t.Fatalf("%s : Count should be 1. Got %d", failed, l.Count())
			}
			t.Logf("%s : Count should be 1.", success)
		}
	}
}

func TestSave(t *testing.T) {

	t.Log("Testing saving a file")
	{
		t.Log("When we are saving a file")
		{
			l := NewList()
			l.Add("Finishing command line program chapter 2")
			l.Add("Finishing calming meditation")

			tmp, err := os.CreateTemp("", "todo.*.json")
			if err != nil {
				t.Fatalf("%s Could not save file", failed)
			}
			l.Save(tmp.Name())
			var m List
			m.Load(tmp.Name())
			for idx, val := range *l {
				val1, _ := m.Get(idx)
				if (*val1).Task != val.Task {
					t.Fatalf("%s l -> %v \n\n m -> %v", failed, val1, val)

				}
			}
			t.Logf("%s write and load suceeded ", success)
		}
	}
}
