package main

import "fmt"

type C struct {
	x float64
	int
	string
}
func main() {
	m := map[string]float64{"1":1.0, "pi":3.1415}
	fmt.Println(m)
	delete(m, "1")
	fmt.Println(m)

	c := C{3.5,7, "hellow"}
	fmt.Println(c.string)

	ch := make(chan int)
	v, ok := <-ch
	fmt.Println(v)
	fmt.Println(ok)
}
