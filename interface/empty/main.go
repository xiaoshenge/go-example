package main

import "fmt"

type Math interface {
	Add(a int, b int) int
}

type mathAdd func(a int, b int) int

func (m mathAdd) Add(a int, b int) int {
	return m(a, b)
}

type DemoMath struct {
	Math
}

// func DemoAdd() mathAdd {
// 	return mathAdd(func(a int, b int) int {

// 	})
// }

func DemoAdd(a int, b int) int {
	return a + b + 1
}

func main() {
	m := DemoMath{mathAdd(DemoAdd)}
	a := m.Add(1, 2)
	fmt.Println(a)
}
