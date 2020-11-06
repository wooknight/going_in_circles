package main

import (
	"fmt"
	"net"
	"sort"
	"sync"

	"honnef.co/go/netdb"
)
var proto *netdb.Protoent

func worker(ports , result chan int){
	for p:= range ports{
		address:=fmt.Sprintf("scanme.nmap.org:%d",p)
		conn,err:=net.Dial("tcp",address)
		if err != nil{
			result<- 0
			continue
		}
		conn.Close()
		result<-p
	}
}


func simpleScan(){

	var wg sync.WaitGroup
	fmt.Printf("I am scanning")
	proto = netdb.GetProtoByName("tcp")
	for i:= 1; i<1024;i++{
		wg.Add(1)
		go func (port int){
			defer wg.Done()
			ToScan:=fmt.Sprintf("scanme.nmap.org:%d",port)
			conn,err:= net.Dial("tcp",ToScan)
			if err != nil {
				// fmt.Printf("Port closed Error %v ",err)
				return
			}
			serv := netdb.GetServByPort(port, proto)
			if serv != nil {
				fmt.Println("Service name : ", serv.Name)
				fmt.Println("Aliases for this service : ")
				for k, v := range serv.Aliases {
						fmt.Println(k, v)
				}
				fmt.Println("Service is using port number : ", serv.Port)
				fmt.Println("######################")
		} else {
			fmt.Printf("Open port found %d ",port)

		}
			conn.Close()
		}(i)

	}
	wg.Wait()
}

func main (){
	ports := make(chan int, 100)
	result:=make (chan int)
	var openports []int
	for i:= 0;i<cap(ports);i++{
		go worker(ports,result)
	}

	go func(){
		for i:=1;i<=1024;i++{
			ports<-i
		}
	}()

	for i:=0;i<1024;i++{
		portNum:= <-result
		if portNum != 0{
			openports=append(openports,portNum)
		}

	}
	proto = netdb.GetProtoByName("tcp")

	close (ports)
	close(result)
	sort.Ints(openports)
	for _,port := range openports{
		serv := netdb.GetServByPort(port, proto)
		if serv != nil {
			fmt.Println("Service name : ", serv.Name)
			fmt.Println("Aliases for this service : ")
			for k, v := range serv.Aliases {
					fmt.Println(k, v)
			}
			fmt.Println("Service is using port number : ", serv.Port)
			fmt.Println("######################")
	} else {
		fmt.Printf("Open port found %d ",port)

	}
	}

}