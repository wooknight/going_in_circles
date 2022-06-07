package main

import (
	"container/heap"
	"fmt"
	// "fmt"
)

type company struct {
	pe       float64
	sales    int64
	expenses int64
	index    int
}

type priorityqueue []*company

func (pq priorityqueue) Len() int {
	return len(pq)
}

func (pq priorityqueue) Less(i, j int) bool {
	fmt.Println(len(pq), pq, i, j)
	return pq[i].pe < pq[j].pe
}

func (pq priorityqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityqueue) Push(comp any) {
	n := len(*pq)
	item := comp.(*company)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityqueue) Pop() any {
	old := *pq
	n := len(old)
	itm := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	itm.index = -1
	return itm
}

func main() {
	i := 0
	items := []company{
		{
			pe:       12,
			sales:    2000,
			expenses: 200,
		},
	}
	pq := make(priorityqueue, len(items))
	for _, com := range items {
		pq[i] = &company{
			pe:    com.pe,
			sales: com.sales,

			expenses: com.expenses,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

}
