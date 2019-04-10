package sort

import (
	"fmt"
)

func ExampleBubbleSort()  {
	arr := []int{1,3,5,2,4,5,3}
	BubbleSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 3 4 5 5]
}

func ExampleInsertionSort()  {
	arr := []int{1,3,5,2,4,5,3}
	InsertionSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 3 4 5 5]
}

func ExampleSelectionSort()  {
	arr := []int{1,3,5,2,4,5,3}
	SelectionSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 3 4 5 5]	
}