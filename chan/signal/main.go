package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{},1)

	go func() {
		fmt.Println("hello")
		time.Sleep(time.Second * 2)
		<-done
	}()

	done <- struct{}{}
	fmt.Println("world!")
}
