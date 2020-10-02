package main

import (
	"fmt"
	"sync"
)

type person struct{
	name string
}
var members = map[int]*person{}

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=0; i< 10000;i++{
			addPerson(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i:=0;i< 1000 ; i++ {
			p := getPerson(i)
			_= p
		}
	}()
	wg.Wait()
}

func addPerson(i int)  {
	p := &person{name:fmt.Sprintf("%d",i)}
	members[i] = p
}

func getPerson(i int) *person  {
	p , ok := members[i]
	if ok {
		return p
	}
	return nil
}
