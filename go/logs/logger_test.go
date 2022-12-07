package caching

import "testing"

var fa int

func BenchmarkColTraverse(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = ColTraverse()
	}
	fa = a
}

func BenchmarkLinkedListTraverse(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}
	fa = a
}

func BenchmarkRowTraverse(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}
	fa = a
}
