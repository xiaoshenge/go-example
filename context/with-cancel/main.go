package main

import (
	"fmt"
	"context"
	"time"
)

func longMathOp(n int, timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res := n

	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			res += i
			time.Sleep(time.Millisecond)
		}
	}
	return res,nil
}

func main()  {
	res, err := longMathOp(5, time.Millisecond*300)
	fmt.Println("err: ", err)
	fmt.Println(res)

	res, err = longMathOp(5, time.Millisecond*10)
	fmt.Println(err)
	fmt.Println(res)

}