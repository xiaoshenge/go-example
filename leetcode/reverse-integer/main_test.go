package main

import (
	"fmt"
)


func ExampleReverse() {
	fmt.Println(Reverse(123))
	fmt.Println(Reverse(-123))
	// Output:
	// 321
	// -321
}