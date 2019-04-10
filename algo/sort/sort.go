package sort

// BubbleSort 冒泡排序
func BubbleSort(a []int)  {
	n := len(a)
	if n <=1{
		return
	}

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag{
			break
		}
	}
}

// InsertionSort 插入排序
func InsertionSort(a []int)  {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >=0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else{
				break
			}
		}
		a[j+1] = value
	}
}

// SelectionSort 选择排序
func SelectionSort(a []int)  {
	n := len(a)
	if n <= 1{
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i+1; j < n; j++{
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i],a[minIndex] = a[minIndex],a[i]
	}
}