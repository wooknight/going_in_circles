package main

import (
	"container/heap"
	"fmt"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/equity"
)

type capitalization int64

const (
	micro_cap capitalization = 300000000
	small_cap                = 2000000000
	mid_cap                  = 10000000000
	large_cap                = 200000000000
	mega_cap                 = 3000000000000
)

type Company struct {
	beta                         float64 // measuring volatility ; volatility of a stock / volatility of the market
	ticker                       string
	year                         int
	stkDetails                   *finance.Equity
	pe                           float64
	sales                        float64
	income, net_income, expenses,net_equity float64
	index                        int
	company_type 				 capitalization
	assets                       int64
	liabilities                  int64
	net_worth                    int64
	current_ratio	float64
	quick_ratio float64
	roe , roa , sales_receivables , recievables  float64

}

func NewCompany(c *Company) {
	c.stkDetails, err := equity.Get(c.ticker)
	if err != nil {
		// Uh-oh!
		panic(err)
	}
	c.equity = c.assets - c.liabilities
	c.net_earnings = c.sales - c.expenses
	c.net_income = c.income - c.expenses
	c.company_type  = c.companySize()
	c.current_ratio = c.totalCurrentAssets / c.TotalCurrentLiabilities
	c.quick_ratio = (c.assets - c.Inventory) / c.liabilities
	c.sales_receivables = c.sales / c.recievables
	c.roe = c.net_income/c.net_equity
	c.roa = c.net_income/float64(c.assets)

}

func (c Company ) excitement(){
	//need to get RSS feeds and parse for company news , new products , financial problems, bad earnings , govt problems, liabilities 
}

func (c Company) companySize() capitalization {
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

func (c Company) valuePE() bool {
	return c.pe <= 20
}

func (c Company) growthPE() bool {
	return c.pe <= 40
}

func (c Company) underValued() bool {
	//company is profitable 
	//net income rising by 10%
	//company is in top 10% of the sector
	//sales are up 10%
	//increasing total assets
	//increasing financial assets
	// !(increasing inventory growth but flat sales)
	//increasing equity
	//price to net earnings
	//trailing p/e
	//forward p/e
	//price to sales (1 < tgt < 4)
	return c.stkDetails.PriceToBook < 2 
}

func (c Company) LongTermGrowth() bool {
	//strong brand
	//high barrier to entry aka moat
	//R&D
	//Company Mkt position
	//Industry
	//Economic prospects
	//sales earnings 
	//debt
	//industry
	//Economic prospects
	return false	
}

func (c Company) IsSolvent()  bool{
	return c.AssetsGrowing() && c.LiabilitiesInControl() && \
	c.QuickRatio() && c.DebtToNetEquity() && c.WorkingCapital() 
}

func (c company) IsLiquid() bool {
	return true
}

func (c company) IsCompanyProfitable() bool{
	return c.ReturnOnEquity() >=  5.0 && c.ReturnOnAssets() >= 5 

}

type priorityqueue []*Company

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
	items := []Company{
		{
			pe:       12,
			sales:    2000,
			expenses: 200,
		},
	}
	pq := make(priorityqueue, len(items))
	for _, com := range items {
		pq[i] = &Company{
			pe:    com.pe,
			sales: com.sales,

			expenses: com.expenses,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

}
