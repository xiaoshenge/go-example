package main

import (
	"fmt"
	// "errors"
	"github.com/pkg/errors"
)

var (
)

func main()  {
	err := process()
	fmt.Println(err)
}

func process() error {
	if err := doSomething(); err != nil{
		return errors.Wrap(err, "processing ")
	}
	return nil
}

func doSomething() error {
	return errors.New("something bad happen")
}