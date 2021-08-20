package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	// "time"
	// "crypto/sha1"
)

type logData struct{
	log string
	datetime string
	count int
}
//our hash table to hold all our data
var logMap map[string]logData

func main(){
	defer func (){
		fmt.Printf("Main Ending")
	}()
	fmt.Printf("Main started")
	logMap = make( map[string] logData)
	r := regexp.MustCompile(`(?P<dt>\d{4}-\d{2}-\d{2})T(?P<tm>\d{2}:\d{2}:\d{2})[+0-9: ]*(?P<level>[A-Za-z]*)[ (0-9)]*[: ]*(?P<data>[a-zA-Z =0-9;]*)`)
	fi,err:=ioutil.ReadDir(".")
	if err != nil{
		fmt.Println("Could not walk like a hooker")
	}
	// h:= sha1.New()
Check:
	for _,file:= range fi {
		// fmt.Println(file.Name())
		fp,err := os.Open(file.Name())
		if err != nil{
			log.Fatalf("Error %v \n File %v",err,file)
		}
		defer fp.Close()
		scanner:= bufio.NewScanner(fp)
		for scanner.Scan(){
//			fmt.Println(scanner.Text())
			if IsValidDate(scanner.Text()){
				values:= r.FindStringSubmatch(scanner.Text())
				// h.Write(([]byte)(values[4]))
				// fmt.Printf("\n%s \n%s \n",values[3],values[4])
				key:= values[3]+"-"+values[4]
				if logMap[key].count == 0 {
					logMap[key] = logData{log:values[4],datetime:values[1],count:1}
				}else{
					cnt := logMap[key].count
					// fmt.Printf("Found duplicates %d",cnt)
					logMap[key] = logData{log:values[4],datetime:values[1],count:cnt+1}
				}
			}else{
				continue Check
			}
		if err:= scanner.Err();err!= nil{
			log.Fatal(err)
		}
	}
		// if date regex found then pick up the next word and use that as the key for the map , the value needs to be a MD5 hash
		//open filec
		// update maps for each line
		// close file
	for _, val := range logMap{
		if val.count > 5{
			fmt.Printf("Occurrio %v\n\n",val)
		}
	}

	}
}

func IsValidDate(line string) bool{
	//2020-09-22T20:22:18+00:00

	re:= regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)
	matched := re.Match( ([]byte)( line))
	if matched == true{
		return true
	}
	return false
}

func splitIntoAtoms(str string) (*regexp.Regexp) {
	//2020-09-22T07:08:01+00:00 INFO (6): AuctionID = 345842; BidId = 1485783; IsAbs = 0;
		// str:= "2020-09-22T07:08:01+00:00 INFO (6): AuctionID = 345842; BidId = 1485783; IsAbs = 0"
		r := regexp.MustCompile(`(?P<dt>\d{4}-\d{2}-\d{2})T(?P<tm>\d{2}:\d{2}:\d{2})[+0-9: ]*(?P<level>[A-Za-z]*)[ (0-9)]*[: ]*(?P<data>[a-zA-Z =0-9;]*)`)
	//r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
		fmt.Printf("%#v\n", r.FindStringSubmatch(str))
		// fmt.Printf("%#v\n", r.SubexpNames())
		return r
	}

