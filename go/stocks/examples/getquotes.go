package main

import (
	"fmt"

	"github.com/piquette/finance-go/equity"
)

// "BTC",	"ETH",	"DOGE",

var holdings = []string{
	"AAPL",
	"AMZN",
	"ARCT",
	"ARGYF",
	"ARKK",
	"ARKX",
	"ASAN",
	"ASVDX",
	"BA",
	"BIRD",
	"BLDE",
	"BV",
	"BYND",
	"CARA",
	"COIN",
	"CORE",
	"COST",
	"COUR",
	"CVCO",
	"DBRG",
	"DUOL",
	"EMB",
	"FB",
	"FDRXX",
	"FIZZ",
	"FLCNX",
	"FXAIX",
	"GOOGL",
	"IP",
	"ITRN",
	"JOBY",
	"MSFT",
	"MTTR",
	"MYTE",
	"NFLX",
	"NVDA",
	"OTLY",
	"PLTR",
	"PLUG",
	"PUBM",
	"QUOT",
	"RIVN",
	"ROKU",
	"SHOP",
	"SLVM",
	"SQ",
	"STNE",
	"TA",
	"TSLA",
	"TSM",
	"TU",
	"TWKS",
	"TWTR",
	"UDMY",
	"UPST",
	"VALE",
	"VEA",
	"VEC",
	"VIG",
	"VNQ",
	"VWO",
	"WEJO",
}

func main() {
	for _, val := range holdings {
		q, err := equity.Get(val)
		if err != nil {
			// Uh-oh!
			panic(err)
		}
		// All good.
		fmt.Printf("\n %v\n", q)

	}

}
