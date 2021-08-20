package main

import (
	"container/list"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/profile"

	// "strconv"
	"io"
	// "gopkg.in/Billups/golang-geo.v2"
)

type Bucketlist struct {
	Name     string
	Category string
	// location Point
}
type MyBucketlist struct {
	Bucketlist
	next, prev *MyBucketlist
	// location Point
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func goList() *list.List {
	defer duration(track("goList"))
	bucketlist := list.New()
	f, err := os.Open("myfile.csv")
	if err != nil {
		fmt.Errorf("Critical file error %v", err)
		return bucketlist
	}
	file := csv.NewReader(f)
	if err != nil {
		fmt.Errorf("Critical file error %v", err)
		return bucketlist
	}
	for {
		record, err := file.Read()
		if err == io.EOF {
			// fmt.Printf("EOF reached \n")
			break
		}
		if err != nil {
			fmt.Printf("Unknown error %v", err)
		}
		// fmt.Printf("Goal %s is in Category %s",record[0],record[1])
		bucketlist.PushBack(Bucketlist{Name: record[0], Category: record[1]})
	}
	return bucketlist
}

func LoadMyList() *MyBucketlist {
	defer duration(track("LoadMyList"))
	bucketlist := new(MyBucketlist)
	f, err := os.Open("myfile.csv")
	if err != nil {
		fmt.Errorf("Critical file error %v", err)
		return bucketlist
	}

	file := csv.NewReader(f)
	if err != nil {
		fmt.Errorf("Critical file error %v", err)
		return bucketlist
	}
	ptr := bucketlist
	for {
		record, err := file.Read()
		if err == io.EOF {
			// fmt.Printf("EOF reached \n")
			break
		}
		if err != nil {
			fmt.Printf("Unknown error %v", err)
		}
		// fmt.Printf("Goal %s is in Category %s for ",record[0],record[1])
		curBucket := MyBucketlist{Bucketlist{record[0], record[1]}, nil, ptr}
		ptr.next = &curBucket
		ptr = ptr.next
	}
	return bucketlist
}

func findGoList(goal string, bl *list.List) {
	defer duration(track("findGoList"))
	// var itm Bucketlist
	for e := bl.Front(); e != nil; e = e.Next() {
		// fmt.Printf("findGoList e= %v\n",e)
		itm := e.Value.(Bucketlist)

		if strings.Contains(itm.Name, goal) == true {
			fmt.Printf("findGoList e= %s\n", itm.Name)
		}
	}
}

func findinMyGoList(goal string, bl *MyBucketlist) {
	defer duration(track("findinMyGoList"))
	for e := bl; e != nil; e = e.next {
		// fmt.Printf("findinMyGoList e= %v\n",e)
		if strings.Contains(e.Name, goal) == true {
			fmt.Printf("found in findinMyGoList e= %v\n", e.Name)
		}

	}
}

/*
ramesh@Ramesh-Naidus-MacBook-Pro-2 linked lists (master) $./linked\ lists
2020/10/12 16:39:20 profile: cpu profiling enabled, /var/folders/l2/j6rstcd51ggcc8x8yqx4twz40000gp/T/profile555803327/cpu.pprof
goList: 511.187µs
findGoList: 10.58µs
LoadMyList: 216.942µs
findinMyGoList: 9.477µs
2020/10/12 16:39:20 profile: cpu profiling disabled, /var/folders/l2/j6rstcd51ggcc8x8yqx4twz40000gp/T/profile555803327/cpu.pprof
ramesh@Ramesh-Naidus-MacBook-Pro-2 linked lists (master) $go tool pprof -pdf ./linked
linked lists   linkedlist.go
ramesh@Ramesh-Naidus-MacBook-Pro-2 linked lists (master) $go tool pprof -pdf ./linked
linked lists   linkedlist.go
ramesh@Ramesh-Naidus-MacBook-Pro-2 linked lists (master) $go tool pprof -pdf ./linked
linked lists   linkedlist.go  myfile.csv     runtime
ramesh@Ramesh-Naidus-MacBook-Pro-2 linked lists (master) $go tool pprof -pdf ./linked\ lists /var/folders/l2/j6rstcd51ggcc8x8yqx4twz40000gp/T/profile555803327/cpu.pprof > cgraph.pdf
ramesh@Ramesh-Naidus-MacBook-Pro-2 linked lists (master)


*** go tool pprof  ./linked\ lists /var/folders/l2/j6rstcd51ggcc8x8yqx4twz40000gp/T/profile555803327/cpu.pprof

	grep func heapsort.go | awk '{print $2}' | awk -F"(" '{printf "func Test" $1 "(t *testing.T){\n"$1"\n}\n" }'  >> heapsort_test.go
*/
func main() {
	defer profile.Start().Stop()
	for i := 1; i < 10000; i++ {
		golistVar := goList()
		findGoList("Bhosada", golistVar)
		myListVar := LoadMyList()
		findinMyGoList("Bhosada", myListVar)

	}
}
