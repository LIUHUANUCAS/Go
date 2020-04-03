package server

import (
// "fmt"
// "log"
// "net"
// "net/http"
// "net/rpc"
// "sync"
// "time"
)

type RpcServer struct {
}

type Args struct {
	A int
	B int
}
type Arith int

func (t *Arith) Add(arg Args, reply *int) error {
	*reply = arg.A + arg.B
	return nil
}

// func main() {
// 	arith := new(Arith)
// 	rpc.Register(arith)
// 	rpc.HandleHTTP()
// 	l, err := net.Listen("tcp", ":1234")
// 	if err != nil {
// 		log.Fatal("err:", err)
// 	}
// 	go http.Serve(l, nil)
// }
