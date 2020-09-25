package main

import (
	"reflect"
	"testing"
)

func TestIsValidDate(t *testing.T){
	//func IsValidDate(line string) bool{
	//valida date true etc etc
	testcases := 	[]struct {
		datestr string
		want bool
	} {
		{"2020-07-09T09:21:21",true},
		{"ahgdhbv",false},
	}
	for _,val := range testcases{
		if IsValidDate(val.datestr) != val.want {
			t.Fatalf("Failed date check %v",val)
		}
	}

}

func TestSplitIntoAtoms(t *testing.T){
	testcases1 :=  []struct {
		datestr string
		wantName []string
		wantData []string
	}{
		{datestr:"2020-09-22T07:08:01+00:00 INFO (6): AuctionID = 345842; BidId = 1485783; IsAbs = 0;",
		wantName: []string{ "", "dt", "tm", "level", "data"},
		wantData:[]string { "2020-09-22T07:08:01+00:00 INFO (6): AuctionID = 345842; BidId = 1485783; IsAbs = 0;", "2020-09-22", "07:08:01", "INFO", "AuctionID = 345842; BidId = 1485783; IsAbs = 0;"}},
	}
	for _, val := range testcases1{
		r:= splitIntoAtoms(val.datestr)
		if ! reflect.DeepEqual(r.SubexpNames() ,val.wantName){
			t.Errorf("Names do not match")
		}
		if ! reflect.DeepEqual( r.FindStringSubmatch(val.datestr) , val.wantData){
			t.Errorf("Values do not match \n\n%v\n%v",r.FindStringSubmatch(val.datestr),val.wantData)
		}
	}

}