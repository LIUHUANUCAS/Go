package main

import (
	// "fmt"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/server"
	// "time"
)

const (
	port = 1234
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("err:", err)
	}
	log.Printf("listen:%d", port)
	http.Serve(l, nil)
}
