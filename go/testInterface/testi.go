package main

import "fmt"

type base struct {
	inti     int
	moth     string
	bearclaw []int
}

func (b *base) changestring() {
	fmt.Println("Love the base", b.inti, b.bearclaw)
}

type inheritor struct {
	base
}

func (i *inheritor) changestring1() {
	fmt.Println("Inherit the wind", i.inti, i.bearclaw)
}

func main() {
	fmt.Println("Starting")
	newbase := base{
		inti:     7,
		moth:     "moth",
		bearclaw: []int{1, 2, 3, 4, 5, 6},
	}
	newbase.changestring()
	inheritor := inheritor{
		base: base{
			inti:     49,
			moth:     "inherited moth",
			bearclaw: []int{1, 4, 9, 16, 25, 36},
		},
	}
	inheritor.changestring()
	inheritor.changestring1()
}
