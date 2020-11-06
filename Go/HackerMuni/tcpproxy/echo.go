package main

import (

	"io"
	"log"
	"net"
)

func echo(conn net.Conn){
	defer conn.Close()
	if _,err:=io.Copy(conn,conn);err!=nil{
		log.Fatalf("Unable to read/write data")
	}

}

func main() {
	listener ,err := net.Listen("tcp",":20080")
	if err!=nil{
		log.Fatal("Unable to bind to port")
	}
	log.Println("Listening on localhost:20080")
	for {
		conn,err:= listener.Accept()
		log.Println("Receiver connection")
		if err!=nil{
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}