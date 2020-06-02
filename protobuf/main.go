package main

import (
	"fmt"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
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
	err = gogoproto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newTest.Title)
	fmt.Println(newTest.Val)
}
