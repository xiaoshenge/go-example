package binarysearch


func BinarySearch(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}
	min := 0
	max := n - 1
	for min <= max {
		mid := min + ((max - min)>>1)
		if a[mid] == v{
			return mid
		} else if a[mid] < v{
			min = mid + 1
		} else {
			max = mid -1
		}
	}
	return -1
}

func BinarySearchRecursion(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	return search(a, v, 0, n-1)
}
func search(a []int, v int, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + ((high-low)>>1)
	if a[mid] == v {
		return mid
	} else if a[mid] < v {
		return search(a, v, mid+1, high)
	} else {
		return search(a, v, low, mid-1)
	}
}