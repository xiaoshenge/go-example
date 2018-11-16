package main

import (
	"fmt"
	"errors"
)

func main() {

	fmt.Println("main enter")
	defer func() {
		if p := recover(); p != nil{
			fmt.Printf("defer panic： %s\n",p)
		}
	}()

	panic(errors.New("something wrong"))
}