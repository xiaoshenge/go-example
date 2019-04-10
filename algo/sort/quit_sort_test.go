package sort

import (
	"fmt"
)

func ExampleQuickSort()  {
	arr := []int{1,3,5,2,4,5,3}
	QuickSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 3 4 5 5]
}