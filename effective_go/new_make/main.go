package main

import (
	"fmt"
)

func main(){
	a := make([]int, 10)
	fmt.Println(a)
	b := append(a, 1)
	fmt.Println(b)
}