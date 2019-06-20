package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var (
		m uint64 = 97
		n uint64 = 1
	)

	const (
		a = 1
	)

	atomic.AddUint64(&m, -n)
	fmt.Println(m)
	atomic.AddUint64(&m, ^uint64(a-1))
	fmt.Println(m)
}
