package main

import (
	"fmt"
	"reflect"
)

type T struct{
	f1 string `test:"f1"`
	f2 string
	f3 string `f three`
	f4, f5 int64 `f four and five`
}

func main()  {
	t := reflect.TypeOf(T{})
	fmt.Printf("%#v\n", t)
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag)
	fmt.Printf("%#v\n", f1)

	fmt.Println("%#v", f1.Tag.Get("test"))

}