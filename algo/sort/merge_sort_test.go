package sort

import (
	"fmt"
)

func ExampleMergeSort()  {
	arr := []int{1,3,5,2,4,5,3}
	MergeSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 3 5 4 5]
}