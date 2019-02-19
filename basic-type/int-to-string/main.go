package main

import (
	"fmt"
	"strconv"
)

func IntToStingWithItoa(i int) string {
	return strconv.Itoa(i)
}
func IntToStringWithSprinf(i int) string {
	return fmt.Sprintf("%d", i)
}

func main()  {
	fmt.Printf("%s\n", IntToStingWithItoa(219))
	fmt.Printf("%s\n", IntToStringWithSprinf(219))
}
