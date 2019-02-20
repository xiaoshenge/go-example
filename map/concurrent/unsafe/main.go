package main

import (
	"fmt"
)

func main()  {
	c := map[string]int{}

	for i := 0; i < 100; i++ {
		go func ()  {
			for j := 0; j < 100000; j++ {
				c[fmt.Sprintf("%d", j)] = j
			}
		}()
	}
}