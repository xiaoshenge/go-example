package main

import (
	"fmt"
)

func towSum(nums []int, target int) []int {
	m := map[int]int{}
	for k,val := range nums{
		diff := target - val
		mk,ok := m[diff]
		m[val] = k

		if ok{
			return []int{mk, k}
		}
	}
	return []int{}
}
func main()  {
	nums := []int{2,7,11,15}
	fmt.Println(towSum(nums, 9))
}