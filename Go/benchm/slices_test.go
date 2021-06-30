package caching

import "testing"

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
