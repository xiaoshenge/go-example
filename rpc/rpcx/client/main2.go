package main

import (
	"example/rpc/rpcx/arith"
	"github.com/smallnest/rpcx/client"
	"context"
	"fmt"
)



func main() {
	args := &arith.Args{7,8}

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key:"tcp@127.0.0.1:8081"},{Key:"tcp@127.0.0.1:8082"}})
	//d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: "tcp@localhost:8972"}, {Key: "tcp@localhost:8973"}})

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	for i :=0; i < 100000; i++ {
		reply := &arith.Reply{}
		err := xclient.Call(context.Background(),"Mul", args, reply)

		if err != nil{
			fmt.Println(err)
		}

		fmt.Printf("%d * %d = %d\n", args.A, args.B, reply.C)
	}

}
