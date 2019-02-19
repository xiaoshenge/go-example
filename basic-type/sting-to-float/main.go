package main

import (
	"fmt"
	"strconv"
)

func StringToFloatWithParseFloat(str string) (float64, error)  {
	f64, err := strconv.ParseFloat(str, 64)
	return f64, err
}
func StringToFloatWithSscanf(str string) (float64, error)  {
	var f float64
	_, err := fmt.Sscanf(str, "%f", &f)
	return f, err
}

func main()  {
	f, err := StringToFloatWithParseFloat("3.1415926")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%f\n", f)


	f, err = StringToFloatWithSscanf("3.1415926")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%f\n", f)

}