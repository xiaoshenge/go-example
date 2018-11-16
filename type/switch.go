package main

import "fmt"

func main() {
	i := 5
	switch {
	case i < 6:
		fmt.Println("a6"); fallthrough
	case i < 7:
		fmt.Println("a7");fallthrough
	case i <8:
		fmt.Println("a8")


	default:
		fmt.Println("default")
	}
}
