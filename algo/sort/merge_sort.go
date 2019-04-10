package sort

func MergeSort(a []int)  {
	n := len(a)
	if n <= 1 {
		return
	}
	mergeSort(a, 0, n-1)
}
func mergeSort(a []int, start, end int)  {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(arr []int, start,mid,end int)  {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid+1
	k := 0
	for  ;  k <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++{
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end;j++{
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}