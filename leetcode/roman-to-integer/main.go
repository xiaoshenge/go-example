package main

import (
	"fmt"
)

func main()  {
	fmt.Println(RomanToInteger("MCMXCIV"))
}

func RomanToInteger(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var tmp int
	var preStr string
	for _, v := range s{
		key := string(v)
		val, ok := roman[key]
		if !ok {
			return 0
		}
		if (preStr == "I" && (key == "V" || key == "X")) || (preStr == "X" && (key == "L" || key == "C")) || (preStr == "C" && (key == "D" || key == "M")) {
			tmp += (val - 2*roman[preStr])
		}else {
			tmp += val
		}
		
		preStr = key
	}
	return tmp
}