package main

import "fmt"

type Reboot struct {
	Name *string
}

func (r Reboot)Echo()  {
	fmt.Println(&r.Name)
}

func main() {
	a := Reboot{}
	tmp := "a"
	a.Name = &tmp

	b := a
	tmp1 := "b"
	b.Name = &tmp1

	a.Echo()
	b.Echo()
}