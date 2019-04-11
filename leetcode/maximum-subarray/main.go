package main

import (
	"math"
	"fmt"
)

func maxSubArray(arr []int) int {
	preSum,max := arr[0], arr[0]
	for index, count := 1, len(arr); index < count; index++ {
		preSum = int(math.Max(float64(arr[index]), float64(arr[index]+preSum)))
		max = int(math.Max(float64(preSum), float64(max)))
	}
	return max
}
func main()  {
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(arr))
}