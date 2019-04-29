package main

import (
	"fmt"
	"errors"
	"sync"
)

var data = make([]int, 0 , 100000)

func init()  {
	for i:=0;i<100000;i++{
		data = append(data, i)
	}
}

var wg sync.WaitGroup
var mu sync.Mutex
func main()  {

	for i:=0;i<10000;i++{
		wg.Add(1)
		go func ()  {
			_,err := PopData()
			if err !=nil{
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	fmt.Println("end")
}

func PopData() (int,error) {
	mu.Lock()
	defer wg.Done()
	if len(data) > 0 {
		i := data[0]
		data = data[1:]
		mu.Unlock()
		return i, nil
	}
	mu.Unlock()
	return 0, errors.New("nodata")
}