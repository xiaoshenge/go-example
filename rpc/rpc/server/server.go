package main

import (
	"net/rpc"
	"net"
	"log"
	"example/rpc/rpc/arith"
	"net/http"
)



func main() {
	a := new(arith.Arith)
	rpc.Register(a)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal("listen error :", e)
	}

	http.Serve(l, nil)
}