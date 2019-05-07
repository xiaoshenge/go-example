package main

import (
	"fmt"
)

func main()  {
	var a interface{}
	var i  = 5
	s := "hello world"

	a = i
	i = a.(int)
	fmt.Printf("%#v\n",i)

	a = s
	s = a.(string)
	fmt.Printf("%#v\n", s)

	params := map[string]interface{}{
		"a": "aaa",
		"b": "97",
	}
	fmt.Println(params["a"].(string))
	fmt.Println(params["b"].(string))
}