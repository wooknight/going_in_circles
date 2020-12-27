package main

import (
	"fmt"
)
type set struct{
	hasher map[int] bool
}

func (p *set) addToSet(val int ) {
	//add to the hash table
	if _,ok:= p.hasher[target]; !ok{
		p.hasher[target]=true
	}
}

func (p *set) getElems() []int{
    keys := make([]int, 0, len(p.hasher))
    for k := range p.hasher {
        keys = append(keys, k)
    }
	return keys
}

func (p *set)Intersection(set1 set) set{

	resultSet = new (set)
	for key, _ := range p.hasher{
		if set1.contains(key){
			resultSet.addToSet(key)
		}
	}
	return resultSet
}

func (p *set) contains(target int){
	if _,ok:= p.hasher[target]; !ok{
		return false
	}
	return true
}

func (p *set)Union(set1 set) set {
	resultSet = new (set)
	for _,_val := range p.getElems(){
		resultSet.addToSet(val)
	}
	for _,_val := range set1.getElems(){
		resultSet.addToSet(val)
	}

	return resultSet
}

func (p *set) Difference(set1 set) set {
	resultSet = new (set)
	for key,_ := range p.hasher(){
		if set1.contains(key)==false{
			resultSet.addToSet(key)
		}
	}
	return resultSet
}