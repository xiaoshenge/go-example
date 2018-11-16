package main

import (
	"github.com/smallnest/rpcx/client"
	"example/rpc/rpcx/arith"
	"context"
	"fmt"
)

func main() {

	args := &arith.Args{7,8}
	reply := &arith.Reply{}
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8081", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	err := xclient.Call(context.Background(), "Mul", args, reply)
	fmt.Println(err)
	fmt.Println(reply)
}
