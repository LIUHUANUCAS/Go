package main

import (
	// "fmt"
	"log"
	"net/rpc"
	server "rpc/server"
)

const (
	serverAddress = "127.0.0.1"
)

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
		return
	}
	// Synchronous call
	args := &server.Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Arith: %d+%d=%d\n", args.A, args.B, reply)
}
