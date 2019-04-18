package main

import (
	"unsafe"
	"fmt"
)
var arr1 = []int{1,2}
func main()  {
	arr := []int{1,2,3}
	fmt.Println(unsafe.Pointer(&arr))
	test(arr)
	test4(&arr)
	fmt.Println(arr)
	// 由于arr是slice，参数传递的是指针，对arr进行test操作后，应该是[2,3],但实际还是[1,2,3]
	// 理解：slice里面具体里面的内容是指针，slice还是值传递, test函数里面的arr是个指拷贝，截取操作改变的是拷贝，对原来的slcie不影响
	// 但是对slice里面原始的影响会传递给上层的arr

	test2(arr)
	test(arr)
	fmt.Println(arr)
	// 对slice里面的具体指还是影响到了，所以是[0,2,3]

	test3(arr)
	fmt.Println(arr)
	fmt.Println(arr1)
}
func test(arr []int)  {
	fmt.Println(unsafe.Pointer(&arr))
	arr = arr[1:]
	fmt.Println(arr)
}

func test2(arr []int)  {
	arr[0] = 0
}

func test3(arr []int)  {
	copy(arr1, arr[1:])
}
func test4(arr *[]int)  {
	fmt.Println(unsafe.Pointer(arr))
}