package main

import (
	"container/heap"
	"fmt"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/quote"
)

type capitalization int64

const (
	micro_cap capitalization = 300000000
	small_cap                = 2000000000
	mid_cap                  = 10000000000
	large_cap                = 200000000000
	mega_cap                 = 3000000000000
)

type company struct {
	beta                         float64 // measuring volatility ; volatility of a stock / volatility of the market
	ticker                       string
	year                         int
	stkDetails                   finance.Equity
	pe                           float64
	sales                        int64
	income, net_income, expenses int64
	index                        int
	company_type 				 capitalization
	assets                       int64
	liabilities                  int64
	net_worth                    int64
}

func NewCompany(c *company) {
	c.stkDetails, err := equity.Get(c.ticker)
	if err != nil {
		// Uh-oh!
		panic(err)
	}
	// All good.
	fmt.Printf("%T \n %+v", q, q)

	c.net_worth = c.assets - c.liabilities
	c.net_income = c.income - c.expenses
	c.company_type  = c.companySize()
}

func (c company) companySize() capitalization {
	if capitalization(c.stkDetails.MarketCap) <= micro_cap {
		return micro_cap
	} else if capitalization(c.stkDetails.MarketCap) <= small_cap {
		return small_cap
	} else if capitalization(c.stkDetails.MarketCap) <= mid_cap {
		return mid_cap
	} else if capitalization(c.stkDetails.MarketCap) <= large_cap {
		return large_cap
	}

	return mega_cap
}

func (c company) valuePE() bool {
	return c.pe <= 20
}

func (c company) growthPE() bool {
	return c.pe <= 40
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
