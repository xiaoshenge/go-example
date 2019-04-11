package main

import (
	"math"
)

func Reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}
	x = int(math.Abs(float64(x)))
	var tmp int
	for x > 0 {
		pop := x%10
		x = x/10

		tmp = tmp*10 + pop
	}
	if tmp > math.MaxInt32 || tmp <= math.MinInt32 {
		return 0
	}
	return tmp*sign
}