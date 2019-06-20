package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	a := []byte("ba")
	fmt.Println(len(a), cap(a)) // 2 32
	a1 := append(a, 'd')
	a2 := append(a, 'g')

	fmt.Println(string(a1)) // bag
	fmt.Println(string(a2)) // bag

	n := 43210 // time in seconds
	fmt.Println(n/(60*60), "hours and", n%(60*60), "seconds")

	time.Sleep(time.Second * 100)
	ch := make(chan int)
	ch <- 1
	sync.Map
}
