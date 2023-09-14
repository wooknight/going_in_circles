package caching

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("bids.log")
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Exiting")
			break
		}
		fmt.Println(str)
	}
	f.Close()
}
