package main

import (
	"fmt"
	"time"
)

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table

	ch := make(chan struct{}, 2)
	ch <- struct{}{}
	ch <- struct{}{}
	close(ch)
	for i := 0; i < 5; i++ {
		ret, ok := <-ch
		fmt.Println(ret)
		fmt.Println(ok)
	}
}

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
