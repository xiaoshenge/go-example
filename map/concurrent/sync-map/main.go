package main

import (
	"sync"
)

func main()  {
	c := sync.Map{}

	for i := 0; i < 100; i++ {
		go func ()  {
			for j := 0; j < 100000; j++ {
				c.Store(j,j)
			}
		}()
	}
}