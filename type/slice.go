package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//pase_student()

test()



}
type studet struct {
	Name string
	Age  int
}

func test()  {
	m := make(map[string]*studet)

	stus := []studet{
		{Name:"a", Age:1},
		{Name:"b", Age:2},
	}
	for _, stu := range stus{
		fmt.Println(unsafe.Pointer(&stu))
		m[stu.Name] = &stu
	}

	fmt.Println(m)

	for k,v := range m {
		fmt.Println(k, v)
	}
}


