package main

import (
	// "runtime"
	"fmt"
	"sync"
)

var counter int

func count(lock *sync.Mutex){
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println("counter=", counter)
}

func main()  {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++{
		count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		// runtime.Gosched()
		if c >= 10 {
			break;
		}
	}
}

