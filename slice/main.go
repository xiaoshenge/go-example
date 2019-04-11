package main

import (
	"fmt"
)

func main()  {
	src := []int{1,2,3}
	dst := make([]int, len(src))
	fmt.Println(dst)
	copy(dst, src)
	fmt.Printf("%#v\n", src)
	fmt.Printf("%#v\n", dst)
	test(&dst)
	fmt.Printf("%#v\n", dst)
}

func test(arr *[]int)  {
	*arr = (*arr)[0:1]
}