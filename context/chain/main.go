package main

import (
	"time"
	"fmt"
	"context"
)

func main()  {
	ctxa,cancel := context.WithCancel(context.Background())
	defer cancel()
	go work(ctxa, "work1")
	ctxb, _ := context.WithTimeout(ctxa, time.Second * 3)
	go work(ctxb, "work2")
	ctxc := context.WithValue(ctxb, "key", "value_c")
	go workWithValue(ctxc, "work3")

	time.Sleep(time.Second *5)
	cancel()
	time.Sleep(time.Second)
}
func work(ctx context.Context, name string)  {
	for {
		select{
		case <- ctx.Done():
			fmt.Println(name, " get message to quit")
			return
		default:
			fmt.Println(name, " is running")
			time.Sleep(time.Second)
		}
	}
}
func workWithValue(ctx context.Context, name string)  {
	for{
		select{
		case <- ctx.Done():
			fmt.Println(name, " get message to quit")
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Println(name, " is running with value ", value)
			time.Sleep(time.Second)
		}
	}	
}
