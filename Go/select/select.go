package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var netEnd chan int
var stdInEnd chan int
var errorEnd chan int

func main() {
	netEnd = make(chan int, 2)   // interesting that I can change from unbuffered to buffered
	stdInEnd = make(chan int, 2) // interesting that I can change from unbuffered to buffered
	http.HandleFunc("/exit", exitServer)
	http.HandleFunc("/", HelloServer)
	go http.ListenAndServe(":8080", nil)

	go checkStdin()

	select {
	case <-stdInEnd:
		fmt.Printf("Stdin ending")
	case <-netEnd:
		fmt.Printf("Network ending")
	case <-errorEnd:
		fmt.Printf("Error ending")
	}
	fmt.Printf("Select done")
}

func exitServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Exiting!")
	netEnd <- 0
	//signal the end
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ,%s!", r.URL.Path[1:])

}

func checkStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		stdInEnd <- 0
		break
	}
}
