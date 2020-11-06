package main

import (
	"fmt"
	"io"
	"log"
	"os"
)
type FooReader struct{}

func (foodReader *FooReader) Read(b []byte) (int ,error){
	fmt.Print("in>")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int , error){
	fmt.Print("out>")
	return os.Stdout.Write(b)
}

func main(){
var (
	reader FooReader
	writer FooWriter
)

	// input:=make([]byte,4096)
	// s,err:=reader.Read(input)
	// if err!= nil{
	// 	log.Fatalln("unable to read data")
	// }
	// fmt.Printf("\nRead %d bytes from stdin\n",s)
	// s,err = writer.Write(input)
	// if err != nil{
	// 	log.Fatalln("unable to write data")
	// }

	if _,err:=io.Copy(&writer, &reader);err != nil{
		log.Fatalln("unable to read/write data")
	}
}