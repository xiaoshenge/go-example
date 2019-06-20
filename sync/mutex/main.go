package main

import (
	"fmt"
	"sync"
)

func main() {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			lock.Lock()
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("done")

}
