package main

import (
	"fmt"
	"testing"
)

func BenchMarkFib(i int, b *testing.B){
	for n:=0;n<b.N;n++{
		x:=fibonacci (100)
		fmt.Println(x)
	}
}

