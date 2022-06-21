package models

import (
	"container/heap"
	"fmt"
	"log"

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


type OperatingActivities struct {
	Net_Income float64 // Company.I.Net_Income
	Accounts_Receivable_Increase , Inventory_Increase , Prepaid_Expenses_Increase , Depreciation_Increase , Accounts_Payable_Increase , Income_Tax_Payable_Increase, Cash_Flow_Operations float64
}
type CashFlow struct {
	O OperatingActivities
	Property_Plant_Equip , Intangible_Assets , Cash_Flow_Investing float64
	S_L Notes_Payable
	Cash_Dividends , Addl_Shares float64
	Cash_Flow_Financial float64
	Cash_Flow_Decrease float64
}

type StockHoldersEquity struct {
	Capital_Stock , RetainedEarnings float64 
}
type Notes_Payable struct {
	Short_Term_Notes, Long_Term_Notes float64
}
type Liabilities struct {
	Accts_Payable , Accrued_Expenses_Payable , Income_Tax_Payable   float64
	S_L Notes_Payable
}
type Assets struct {
	Cash , Accts_Recv , Inventory , Prepaid_Exp , Property_Plant_Equip , Accum_Depreciation , Intangible_Assets , Total_Assets float64
}

type BalanceSheet struct {
	A Assets
	L Liabilities
	S StockHoldersEquity
}

type IncomeStmt struct {
	Sales_Revenue , GrossMargin , SellGenAdm , Depreciation , Operating_Earnings , Interest_Exp, EBIT , IT , Net_Income , COGS float64 
}


type Company struct {
	B BalanceSheet
	C CashFlow
	I IncomeStmt
	beta                         float64 // measuring volatility ; volatility of a stock / volatility of the market
	ticker                       string
	year                         int
	stkDetails                   finance.Equity
	pe                           float64
	sales                        float64
	income, expenses,net_equity float64
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
	c.stkDetails, err = equity.Get(c.ticker)
	if err != nil {
		// Uh-oh!
		panic(err)
	}
	// All good.
	fmt.Printf("%T \n %+v", q, q)
	if c.I.Sales_Revenue < c.B.A.Accts_Recv {
		log.Fatal("Sales Revenue cannot be less than Accounts Receivable", c)
	}
	if c.I.SellGenAdm < c.B.A.Prepaid_Exp + c.B.L.Accts_Payable + c.B.L.Accrued_Expenses_Payable{
		log.Fatal("Selling General and Admin Expenses cannot be less than Accts Payable  + Expenses ", c)
	}
	c.net_worth = c.assets - c.liabilities
	c.company_type  = c.companySize()
	c.current_ratio = c.totalCurrentAssets / c.TotalCurrentLiabilities
	c.quick_ratio = (c.assets - c.Inventory) / c.liabilities
	c.sales_receivables = c.sales / c.recievables
	c.roe = c.I.Net_Income/c.net_equity
	c.roa = c.I.Net_Income/float64(c.assets)

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

func (c Company) IsLiquid() bool {
	return true
}

func (c Company) IsCompanyProfitable() bool{
	return c.ReturnOnEquity() >=  5.0 && c.ReturnOnAssets() >= 5 

}