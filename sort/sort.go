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
