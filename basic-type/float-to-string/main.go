package main

import (
	"fmt"
	"strconv"
)

func FloatToStringWithFromatFloat(f64 float64) string  {
	return strconv.FormatFloat(f64, 'e', -1, 64)
}
func FloatToStringWithSprintf(f64 float64) string {
	return fmt.Sprintf("%f", f64)
}
func main()  {
	fmt.Printf("%s\n", FloatToStringWithFromatFloat(3.1415926))
	fmt.Printf("%s\n", FloatToStringWithSprintf(3.1415926))
}