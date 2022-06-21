package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

/*

https://www.sec.gov/edgar/searchedgar/cik.htm

data.sec.gov/api/xbrl/companyfacts/

*/
func readFile(name string) {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Could not open file")
		return
	}
	ProcessJson(content)
}
func ProcessJson(data []byte) error {
	var stkData map[string]interface{}
	err := json.Unmarshal(data, &stkData)
	if err != nil {
		log.Fatal("Could not json unmarshall data", err)
	}
	for key := range stkData {
		fmt.Println("Key =>", key, "Value =>")
	}
	fmt.Println("Key => cik ", "Value =>", stkData["cik"])
	dt, ok := stkData["facts"].(map[string]interface{})
	if !ok {
		fmt.Println(" Could not get facts Value")
	}
	us_gaap, ok := dt["us-gaap"].(map[string]interface{})
	if !ok {
		fmt.Println(" Could not get facts data us-gaap Value")
	}
	for key, val := range us_gaap {

		if strings.Contains(key, "ncome") {
			dt, ok := val.(map[string]interface{})
			if !ok {
				fmt.Println(" Could not get income Value")
			}
			units, ok := dt["units"].(map[string]interface{})
			if !ok {
				fmt.Println(" Could not get units Value")
			}
			usd, ok := units["USD"].([]interface{})
			if !ok {
				fmt.Println(" Could not get usd Value")
			}
			for key, val1 := range usd {
				val2, _ := val1.(map[string]interface{})
				yr := val2["fy"].(float64)
				if int(yr) < time.Now().Year()-3 {
					usd = append(usd[:key], usd[key+1:]...)
				}
			}

		} else {
			delete(us_gaap, key)
		}
	}
	fmt.Println(stkData)
	return nil
}

func main() {
	readFile("C:\\Users\\ram_n\\OneDrive\\Documents\\CIK0001045810.json")
}
