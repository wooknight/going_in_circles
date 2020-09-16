package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func inbound(chln chan bool,str string) {
	defer func() {
		fmt.Printf("Inside defer inbound %s\n%s\n",str,myStr)
		chln <- true
	}()
	fmt.Printf("inbound - %s ---- %s",str,myStr)
	sysSig:= make(chan os.Signal,1)
	signal.Notify(sysSig,os.Kill,syscall.SIGINT,syscall.SIGABRT)

	<-sysSig
	fmt.Printf("Crap happendn inside subroutine %s\n",str)
}
func outbound(chln chan bool,str string) {
	defer func() {
		fmt.Printf("Inside defer outbound %s --- %s\n",str,myStr)
		if r:=recover();r!=nil && r=="Bhosada laga re"{
			chln <- true
		}
	}()
	fmt.Println("Outbound")
	panic("pucchi laga re")
}
var myStr string
func main() {

	fmt.Println("Inside main")
	waitUntilErr := make(chan bool)

	myVal:="testy"
	myStr = "fast"
	go inbound(waitUntilErr,myVal)
	myVal="fiesty"
	myStr = "slow"
	go outbound(waitUntilErr,myVal)

	sysError := make(chan os.Signal, 1)
	signal.Notify(sysError, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	select {
	case <-waitUntilErr:
		fmt.Println("Sub routine crashed")

	case <-sysError:
		fmt.Println("System error called\nCrap happendn inside main")
		fmt.Println("System error called\nCrap happendn inside main")
		fmt.Println("System error called\nCrap happendn inside main")
		time.Sleep(1000)
	}
	fmt.Println("This is the end")
}
