package main

import (
	"fmt"
)

type Person struct{
	Name string
}

func (p Person)WhoSay()  {
	fmt.Printf("%s say.\n", p.Name)
}

func (p *Person)SayWhat()  {
	fmt.Printf("%s say haha.\n", p.Name)
}

func main()  {
	adam := Person{Name:"adam"}
	stephen := new(Person)
	stephen.Name = "stephen"

	adam.SayWhat()
	adam.WhoSay()

	stephen.SayWhat()
	stephen.WhoSay()

	a := &Person{}
	if a != nil {
		fmt.Println("a is not nil")
	} else {
		fmt.Println("a is nil")
	}
}