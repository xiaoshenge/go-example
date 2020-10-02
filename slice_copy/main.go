package main

import (
	"fmt"
)

var arr = []int{1,2,3}
var arr1 = []int{1,2,3}

func main()  {
	fmt.Println("len arr:",len(arr))
	fmt.Println("cap arr:",cap(arr))
	arr = arr[1:]

	fmt.Println("len arr:",len(arr))
	fmt.Println("cap arr:",cap(arr))

	fmt.Println("len arr1:",len(arr1))
	fmt.Println("cap arr1:",cap(arr1))

	copy(arr1, arr1[1:])
	fmt.Println("len arr1:",len(arr1))
	fmt.Println("cap arr1:",cap(arr1))
	fmt.Println(arr1)

	arr1 = arr1[:len(arr1)-1]
	fmt.Println("len arr1:",len(arr1))
	fmt.Println("cap arr1:",cap(arr1))
	fmt.Println(arr1)

}