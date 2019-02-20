package main

import (
	"fmt"
)

type Person struct{
	Name string
}

func main()  {
	c := new(Person)
	c.Name = "cc"
	fmt.Println(c.Name)

	d := c
	d.Name = "dd"
	fmt.Println(c.Name)

	i := *d
	i.Name="ii"
	fmt.Println(c.Name)
	fmt.Println(d.Name)
	fmt.Println(i.Name)

	fmt.Printf("c type: %T\n", c)
	fmt.Printf("c : %#v\n", c)
	fmt.Printf("c ptr:%p\n", &c)

	fmt.Printf("d type: %T\n", d)
	fmt.Printf("d: %#v\n", d)
	fmt.Printf("d ptr:%p\n", &d)

	fmt.Printf("i type: %T\n", i)
	fmt.Printf("i: %#v\n", i)
	fmt.Printf("i ptr:%p\n", &i )
}