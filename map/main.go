package main

import (
	"fmt"
)

func main()  {
	test := make(map[string]int)
	test["a"] = 1
	test["a"] = 2
	fmt.Println(test)

	if _, ok := test["a"]; ok {
		fmt.Println(ok)
	}
	if _, ok := test["b"]; !ok {
		fmt.Println(ok)
	}
}