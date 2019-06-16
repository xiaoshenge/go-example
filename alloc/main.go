package main

import (
	"fmt"
)

func main() {
	a1 := make(map[string]string)
	a2 := new(map[string]string)
	a3 := map[string]string{}
	var a4 map[string]string
	fmt.Println("a1:", a1)
	if a1 == nil {
		fmt.Println("a1 is nil")
	}
	fmt.Println("a2:", a2)
	if a2 == nil {
		fmt.Println("a2 is nil")
	}
	fmt.Println("a3", a3)
	if a3 == nil {
		fmt.Println("a3 is nil")
	}
	fmt.Println("a4:", a4)
	if a4 == nil {
		fmt.Println("a4 is nil")
	}
	a3["a"] = "a"
	a4["a"] = "a"
}
