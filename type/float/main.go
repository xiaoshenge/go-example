package main

import (
	"fmt"
	"time"
)

func main() {
	var x float32 = 5.21
	a := x * float32(100.00)
	fmt.Println(int64(a))

	b := 1
	switch b {
	case 1, 2, 3:
		fmt.Println("<3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("zz")
	}
	fmt.Println(b)

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	curTime := time.Now()
	begin := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 1, 0, 0, 0, curTime.Location())
	fmt.Println(begin.Hour())

}
