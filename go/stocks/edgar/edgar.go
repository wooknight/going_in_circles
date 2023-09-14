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
	// for key := range stkData {
	// 	fmt.Println("Key =>", key, "Value =>")
	// }
	// fmt.Println("Key => cik ", "Value =>", stkData["cik"])
	dt, ok := stkData["facts"].(map[string]interface{})
	if !ok {
		fmt.Println(" Could not get facts Value")
		return nil
	}
	us_gaap, ok := dt["us-gaap"].(map[string]interface{})
	if !ok {
		fmt.Println(" Could not get facts data us-gaap Value")
		return nil
	}

	// {
	// 	mystkLen, _ := json.Marshal(stkData)
	// 	fmt.Println(len(us_gaap), len(string(mystkLen)))
	// }

	for key, val := range us_gaap {

		if strings.Contains(key, "NetIncomeLoss") {
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
			for idx := 0; idx < len(usd); idx++ {
				val1 := usd[idx]
				val2, _ := val1.(map[string]interface{})
				yr := val2["fy"].(float64)
				if int(yr) < time.Now().Year()-3 || val2["form"] != "10-K" {
					usd = append(usd[:idx], usd[idx+1:]...)
					idx--
				}
			}

		} else {
			delete(us_gaap, key)
		}
	}
	by, err := json.Marshal(stkData)
	if err != nil {
		log.Fatal("Error while marshalling", err)
	}

	// {
	// 	dt, ok := stkData["facts"].(map[string]interface{})
	// 	if !ok {
	// 		fmt.Println(" Could not get facts Value")
	// 		return nil
	// 	}
	// 	us_gaap, ok := dt["us-gaap"].(map[string]interface{})
	// 	if !ok {
	// 		fmt.Println(" Could not get facts data us-gaap Value")
	// 		return nil
	// 	}
	// 	fmt.Println("Len Checking ", len(us_gaap), len(by))
	// }
	fmt.Println(string(by))
	return nil
}

func main() {
	readFile("CIK0001045810.json")
}
