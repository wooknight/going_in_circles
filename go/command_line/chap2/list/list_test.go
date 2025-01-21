package list

import (
	"fmt"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	var testtodos List
	t.Log("testing list")
	task := Todo{Title: "test", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task)
	task1 := Todo{Title: "test1", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task1)
	str := testtodos.List()

	if str != `[{"Title":"test","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""},{"Title":"test1","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""}]` {
		log.Println(str)
		t.Fatal("list does not match")

	}
}

func TestAdd(t *testing.T) {
	var testtodos List
	t.Log("testing add")
	task := Todo{Title: "test", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task)
	if len(testtodos) != 1 {
		t.Fatal("add does not match")
	}

}

func TestDelete(t *testing.T) {
	t.Log("testing delete")
	var testtodos List
	t.Log("testing list")
	task := Todo{Title: "test", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task)
	task1 := Todo{Title: "test1", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task1)
	str := testtodos.List()

	if str != `[{"Title":"test","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""},{"Title":"test1","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""}]` {
		log.Println(str)
		t.Fatal("list does not match")

	}
	testtodos.DeleteByTitle("test")
	str = testtodos.List()
	if str != `[{"Title":"test1","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""}]` {
		log.Println(str)
		t.Fatal("deleted list does not match")
	}
}

func TestComplete(t *testing.T) {
	t.Log("testing complete")
	var testtodos List
	t.Log("testing list")
	task := Todo{Title: "test", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task)
	task1 := Todo{Title: "test1", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task1)
	str := testtodos.List()

	if str != `[{"Title":"test","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""},{"Title":"test1","Done":false,"CreatedAt":"2019-02-01T12:00:00Z","CompletedAt":""}]` {
		log.Println(str)
		t.Fatal("list does not match")

	}
	var found bool
	testtodos.CompleteByTitle("test")
	for _, todo := range testtodos {
		if todo.Title == "test" && !todo.Done {
			fmt.Println(testtodos.List())
			t.Fatal("complete does not match")
		} else {
			found = true
		}
	}
	if !found {
		t.Fatal("complete does not match")
	}
}

func TestWrite(t *testing.T) {
	var testtodos List
	t.Log("testing list")
	task := Todo{Title: "test", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task)
	task1 := Todo{Title: "test1", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task1)
	testtodos.Save("test.txt")
	testtodos.Load("test.txt")
	if len(testtodos) != 2 {
		t.Fatal("write does not match")
	}
}

func TestRead(t *testing.T) {
	var testtodos List
	t.Log("testing list")
	task := Todo{Title: "test", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task)
	task1 := Todo{Title: "test1", Done: false, CreatedAt: "2019-02-01T12:00:00Z"}
	testtodos.Add(task1)
	testtodos.Save("test.txt")
	testtodos.Load("test.txt")
	if len(testtodos) != 2 {
		t.Fatal("write does not match")
	}
}
