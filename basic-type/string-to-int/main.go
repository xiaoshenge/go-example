package main

import (
	"fmt"
	"strconv"
)

func StringToIntWithAtoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToIntWithSscanf(s string) (int, error)  {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
func main()  {
	s, err := StringToIntWithAtoi("219")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", s)

	s, err = StringToIntWithSscanf("219")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", s)
}