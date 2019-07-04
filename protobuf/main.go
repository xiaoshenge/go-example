package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/xiaoshenge/go-example/protobuf/demo"
)

func main() {
	test := &demo.Test{
		Title: "hello world!",
		Val:   20190705,
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	newTest := &demo.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newTest.Title)
	fmt.Println(newTest.Val)
}
