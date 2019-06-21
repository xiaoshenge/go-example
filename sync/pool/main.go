package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := pool.Get().(int)
	pool.Put(1)
	b := pool.Get().(int)
	fmt.Println(a, b)

	runtime.GC()
	pool.Put(2)
	runtime.GC()
	c := pool.Get().(int)
	fmt.Println(c)
}
