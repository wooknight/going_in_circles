package heapsort
// 	grep func heapsort.go | awk '{print $2}' | awk -F"(" '{printf "func Test" $1 "(t *testing.T){\n"$1"\n}\n" }'  >> heapsort_test.go

import (
	// "fmt"

)

func heapify(arr []int) *[]int{
	chhot:=make([]int,0)
	for _,val := range arr{
		// fmt.Printf("%v\n",chhot)
		chhot = pqInsert(chhot,val)
	}
	// fmt.Printf("%v\n",chhot)
	return &chhot
}

func heapsort(arr []int) []int{
	choot:= heapify(arr)
	for i := range arr {
		arr[i] = pqExtractMin(choot)
	}
	return arr
}

func pqInsert(heap []int, val int) []int{
	heap = append(heap,val)
	idx:= len(heap)-1
	if idx > 0 {
		bubbleUp(&heap,len(heap)-1)
	}
	return heap
}

func pqExtractMin(heap *[]int ) int{
	val:= (*heap)[0]
	maxIdx := len(*heap)-1
	(*heap)[0] = (*heap)[maxIdx]
	*heap = (*heap)[:maxIdx]
	bubbleDown(heap , 0)
	return val
}

func digitalRoot( num int) int {
	if num < 10 { return num}
	return digitalRoot ((num % 10) + digitalRoot( num /10))

}
func bubbleUp(heap *[]int , idx int){
	if pqParent(*heap ,idx) == -1{
		return
	}
	if (*heap)[pqParent(*heap,idx)] > (*heap)[idx]{
		(*heap)[pqParent(*heap,idx)],(*heap)[idx] = (*heap)[idx],(*heap)[pqParent(*heap,idx)]
		bubbleUp(heap , pqParent(*heap,idx))
	}
}

func bubbleDown(heap *[]int, idx int){
	smallest:=idx
	left := pqLeftChild(idx)
	if left < len(*heap) && (*heap)[left] < (*heap)[smallest]{
		smallest = left
	}
	right := pqRightChild(idx)//get the right child
	if right < len(*heap) && (*heap)[right] < (*heap)[smallest]{
		smallest = right
	}
	if idx != smallest {
		(*heap)[smallest],(*heap)[idx] = (*heap)[idx],(*heap)[smallest]
		bubbleDown(heap , smallest)
	}
}

func pqParent(heap []int, idx int) int{
	if idx == 0 {
		return -1
	}
	return (idx-1)/2
}

func pqLeftChild( idx int) (int) {
	return idx*2+1
}

func pqRightChild( idx int) (int) {
	return idx*2+2
}