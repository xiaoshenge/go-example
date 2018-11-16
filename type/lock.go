package main

import (
	"sync"
	"time"
	"fmt"
	"context"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex()*Mutex  {
	mu := &Mutex{ch:make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (mu *Mutex)Lock()  {
	<- mu.ch
}

func (mu *Mutex)Unlock()  {
	select {
	case mu.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (mu *Mutex)TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <- mu.ch:
		timer.Stop()
		return true
	case <- time.After(timeout):
		return false
	}
}

func (mu *Mutex)Trylock(ctx context.Context) bool {
	select {
	case <- ctx.Done():
		return false
	case <- mu.ch:
		return true
	}
}


func main() {
	mu := NewMutex()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("in fun1")
		defer wg.Done()
		//locked := mu.TryLock(time.Second * 3)
		ctx , cancel:= context.WithTimeout(context.Background(), time.Second * 3)
		defer cancel()
		locked := mu.Trylock(ctx)
		if locked {
			defer mu.Unlock()
			fmt.Println("fun1 get lock")
			time.Sleep(time.Second * 10)
		}
		fmt.Println("end fun1")
	}()

	go func() {
		fmt.Println("in fun2")
		defer wg.Done()
		//locked := mu.TryLock(time.Second *3)
		ctx , cancel:= context.WithTimeout(context.Background(), time.Second * 3)
		defer cancel()
		locked := mu.Trylock(ctx)
		if locked {
			defer mu.Unlock()
			fmt.Println("fun2 get lock")
			time.Sleep(time.Second * 10)
		}
		fmt.Println("end fun2")
	}()

	wg.Wait()
}
