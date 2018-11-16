package main

import (
	"net/rpc"
	"log"
	"fmt"
	"example/rpc/rpc/arith"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &arith.Args{A:7, B:8}
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(args)
	fmt.Println(reply)


	args = &arith.Args{7,3}
	replay2 := new(arith.Quotient)


	//divCall := client.Go("Arith.Divide", args, replay2, nil)
	//replayCall := <-divCall.Done
	//fmt.Println(replayCall)
	//fmt.Println(replay2)

	done := make(chan *rpc.Call,1)
	client.Go("Arith.Divide", args, replay2, done)
	fmt.Println(<-done)
	fmt.Println(replay2)
}
