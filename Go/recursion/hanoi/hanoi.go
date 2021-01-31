package main

import (
	"fmt"

)

type stack struct{
	label string
	values []int
	idx int
}

func (p *stack)push (val int){
	fmt.Printf("%s : Pushing %d onto %v\n",p.label , val,p.values)
	p.values = append(p.values, val)
	p.idx++
}

func (p *stack)pop () int {
	if p.idx >= 0{
		valToRet:= p.values[p.idx-1]
		fmt.Printf("%s: popping %d from %v\n",p.label,valToRet,p.values)
		p.values[p.idx-1] = 0
		p.values = p.values[:p.idx-1]
		p.idx--
		return valToRet
	}
	return -1
}


func main (){
	fmt.Printf("Start of the towers of Hanoi\n\n")
	start := new (stack)
	start.label = "start"
	finish := new (stack)
	finish.label = "finish"
	temp := new (stack)
	temp.label = "temp"
	start.push(27)
	start.push(53)
	start.push(42)

	moveDisks(3,start,finish,temp)
}

func moveDisks(num int, start, finish, temp *stack){
	if num == 1{
		finish.push(start.pop())
	}else {
		moveDisks(num-1,start,temp,finish)
		finish.push(start.pop())
		moveDisks(num-1,temp,finish,start)
	}

}