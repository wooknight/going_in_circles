package caching

import (
	"fmt"
	"testing"
)

// /go test -run none -bench . -benchtime 3s
var fa int

func BenchmarkLinkListTraversal(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}
	fa = a
}

func BenchmarkColumnTraversal(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}
	fa = a
}

func BenchmarkRowTraversal(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}
	fa = a
}

//go test -bench Sprint -benchtime 3s
var gs string

func BenchmarkSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello ramesh")
	}
	gs = s // so that the compiler does not optimize away the test
}

func BenchmarkSprint(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello ramesh")
	}
	gs = s // so that the compiler does not optimize away the test

}
