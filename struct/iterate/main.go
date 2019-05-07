package main

import (
	"fmt"

	"github.com/fatih/structs"
)

func main() {
	x := struct {
		Foo string `structs:"foo" json:"foo"`
		Bar int
	}{"foo", 2}

	// v := reflect.ValueOf(x)

	// values := make([]interface{}, v.NumField())

	// for i := 0; i < v.NumField(); i++ {
	//     values[i] = v.Field(i).Interface()
	// }

	// fmt.Println(values)
	m := structs.Map(x)
	fmt.Printf("%#v", m)
}
