package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	var bytes []byte
	var wg sync.WaitGroup
	ir, iw := io.Pipe()
	// er, ew := io.Pipe()
	cmd := exec.Command("ls", "-al")
	cmd.Stdout = iw
	wg.Add(1)
	go func(r io.Reader, bytes []byte) {

		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("From Pipe - ", string(bytes))
		os.WriteFile("/tmp/ram", bytes, 0644)
		wg.Done()
	}(ir, bytes)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	iw.Close()
	wg.Wait()

}
