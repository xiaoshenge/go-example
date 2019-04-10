package sort


func QuickSort(arr []int)  {
	n := len(arr)
	if n <= 1{
		return
	}
	quickSort(arr, 0, n-1)
}
func quickSort(a []int, start,end int)  {
	if start >= end {
		return
	}
	i := partition(a, start, end)
	quickSort(a, start, i-1)
	quickSort(a, i+1, end)
}
func partition(arr []int, start,end int) int {
	pivot := arr[end]
	var i = start
	for j := i; j < end;j++{
		if arr[j] < pivot{
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i],arr[end] = arr[end], arr[i]
	return i
}