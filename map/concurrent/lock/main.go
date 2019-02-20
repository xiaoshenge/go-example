package main

import (
	"sync"
	"fmt"
)

type RWMap struct{
	sync.RWMutex
	m map[string]int
}

func (r *RWMap)Set(key string, val int)  {
	r.Lock()
	defer r.Unlock()
	r.m[key] = val
}

func main()  {
	c := RWMap{m:map[string]int{}}

	for i := 0; i < 100; i++ {
		go func ()  {
			for j := 0; j < 100000; j++ {
				c.Set(fmt.Sprintf("%d",j), j)
			}
		}()
	}
}