package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	//ch <- 1

	//m := <-ch
	//fmt.Println(m)
	//n, ok := <- ch
	//fmt.Println(ok)
	//fmt.Println(n)

	select {
	case <- ch:
		fmt.Println("get ch val")
	//default:
	//	fmt.Println("no")
	case <- time.After(2 * time.Second):
		fmt.Println("timeout")


	}
	fmt.Println("end")

}