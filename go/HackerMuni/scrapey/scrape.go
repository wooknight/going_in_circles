package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowURLRevisit(), colly.MaxDepth(2))
	c.IgnoreRobotsTxt = false
	c.OnRequest(func(r *colly.Request) {
		if strings.Contains(r.URL.String(),"facebook") || strings.Contains(r.URL.String(),"google"){
			return
		}
		fmt.Println("Visiting ", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnResponseHeaders(func(r *colly.Response) {
		fmt.Println("Response Headers :", r.Request.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response URL :", r.Request.URL)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	c.Visit("https://onlineservices.cdtfa.ca.gov/?Link=PermitSearch")
	// https://www.colorado.gov/revenueonline/_/#1 - "Dc-c" - POST

	// // authenticate
	// err := c.Post("https://onlineservices.cdtfa.ca.gov/_/", map[string]string{"d-3": "SITSUT", "d-4": "265720192"})
	// if err != nil {
	// 	fmt.Printf("Error - %v\n",err)
	// }

	// <option value="" selected="selected" class="BlankOption watermark"></option>
	// <option value="ACCAJF">Aircraft Jet Fuel Dealer Permit</option>
	// <option value="ACCSUTUSE">Certificate of Registration - Use Tax</option>
	// <option value="ACTCDL">Cigarette &amp; Tobacco Products Distributor License</option>
	// <option value="ACTCWL">Cigarette &amp; Tobacco Products Wholesaler License</option>
	// <option value="ACTCML">Cigarette Manufacturer/Importer License</option>
	// <option value="ACTCRL">Cigarette Retailer License</option>
	// <option value="ACCDDF">Diesel Fuel Supplier License</option>
	// <option value="ACCEWR">E-Waste Account</option>
	// <option value="ACCSMF">Motor Vehicle Fuel Supplier License</option>
	// <option value="SITSUT">Sellers Permit</option>
	// <option value="ACTTML">Tobacco Manufacturer/Importer License</option>
	// <option value="ACCUTF">Underground Storage Tank Account</option>

	// <div class="ViewFieldWrapper"><input type="text" autocomplete="off" name="d-4" id="d-4" class="FieldEnabled FieldRequired DocControlMask FastEvtFieldKeyDown FastFieldEnterEvent FastEvtFieldFocus" value="" data-fast-enter-event="d-4" maxlength="25" 0="" aria-required="true" style="" placeholder="Required   " data-hasqtip="3" oldtitle="Required" title="" aria-describedby="qtip-3">
	// <div id="indicator_d-4" class="FI FieldRequiredIndicator" title="Required"></div></div>

	// Search
	//
}
