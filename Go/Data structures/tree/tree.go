package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func diskUsage(dir string, duInfo os.FileInfo) int {
	var du int
	if duInfo.IsDir() == false {
		return int(duInfo.Size())
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		du += diskUsage(dir+"/"+file.Name()+"/", file)
	}
	fmt.Printf("Size of directory %s - %d\n", dir, du)
	return du
}

func main() {
	dir := "/Users/ramesh/Documents/Repos/"
	fmt.Println("Inside ", dir)
	duInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Error in ", dir, err)
		log.Fatal(err)
	}
	fmt.Println(dir, diskUsage(dir, duInfo))
}
