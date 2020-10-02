package main

import (
	"fmt"
	"os"
)

func main()  {
	a := &TestStruct{}
	b := new(TestStruct)
	c := fmt.Sprintf("%#v", b)
	fmt.Println("c: %s", c)
	*b = *a

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)

	b.setStName("b")

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)

	fmt.Println(os.Getenv("adam"))
	fmt.Println(os.Getenv("adam1"))

}

type TestStruct struct {
	st *StructName
}

func (test *TestStruct) setStName(name string)  {
	if test.st == nil {
		test.st = new(StructName)
	}
	test.st.name = name
}

type StructName struct {
	name string
}
