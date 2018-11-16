package main

import (
	"github.com/smallnest/rpcx/server"
	"example/rpc/rpcx/arith"
)

func main() {
	srv := server.NewServer()
	srv.RegisterName("Arith", new(arith.Arith2), "")
	srv.Serve("tcp", "127.0.0.1:8082")
}
