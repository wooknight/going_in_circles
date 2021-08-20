package main

import (
	"fmt"

	"log"
	"net/http"

	_ "net/http/pprof"
	"os"
	"os/exec"
	"time"
)

var envVar string
func main() {
	defer  log.Println("Function end is coming",os.Getenv("WORKER"))

	if(len(os.Args) == 2) && (os.Args[1]=="outbound"){
		log.Println("Inside the outbound process")
		envVar = os.Args[1]
		go getVar()
		time.Sleep(1 * time.Hour)

	}else {
	chln:= make(chan int)
	go func (chln chan int) {
		defer func() { chln <- 1 }()
		//create a channel

		go func(){
			log.Println("Debug ended ",http.ListenAndServe(":5000",nil))
		}()

		cmd := exec.Command(os.Args[0], "outbound")
		cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"WORKER=outbound",    // this value is used
	)
		// err := cmd.Start()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("waiting on Command ", )
		// err = cmd.Wait()
		out,err:= cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Command finished with output: %s", out)

	}(chln)
	envVar = "inbound"
	os.Setenv("WORKER",envVar)
	go func(){
		log.Println("Debug ended ",http.ListenAndServe(":5001",nil))
	}()

	go getVar()
	<-chln
	}
	log.Println("We are coming to the end of thread - ",os.Getenv("WORKER"))
}

func getVar (){
	fmt.Println("my version of the variable is ",envVar)
}