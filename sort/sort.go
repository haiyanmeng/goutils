package sort

func merge(a, b []int) []int {
	var p1, p2 int
	result := []int{}
	n1, n2 := len(a), len(b)
	for p1 < n1 && p2 < n2 {
		if a[p1] <= b[p2] {
			result = append(result, a[p1])
			p1++
		} else {
			result = append(result, b[p2])
			p2++
		}
	}
	if p1 < n1 {
		result = append(result, a[p1:]...)
	}
	if p2 < n2 {
		result = append(result, b[p2:]...)
	}
	return result
}

func MergeSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	b1 := MergeSort(a[:n/2])
	b2 := MergeSort(a[n/2:])
	return merge(b1, b2)
}

func QuickSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	i := partition(a)
	QuickSort(a[:i])
	QuickSort(a[(i + 1):])
}

func partition(a []int) int {
	n := len(a)
	pivot := a[0]
	i, j := 1, n-1
	for i < j {
		for i < n && a[i] < pivot {
			i++
		}
		for j > 0 && a[j] >= pivot {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	a[0], a[j] = a[j], a[0]
	return j
}

func BucketSort(a []int) {
	digit := 1
	n := len(a)
	for {
		buckets := make([][]int, 10)
		zeroCount := 0
		for _, elem := range a {
			i := (elem / digit) % 10
			buckets[i] = append(buckets[i], elem)
			if i == 0 {
				zeroCount++
			}
		}
		if zeroCount == n {
			return
		}
		a = a[:0]
		for _, bucket := range buckets {
			a = append(a, bucket...)
		}
		digit *= 10
	}
}
