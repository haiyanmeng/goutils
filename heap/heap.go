package heap

import "fmt"

type Heap struct {
	data []int
	n    int
}

func New(arr []int) *Heap {
	h := &Heap{}
	for _, elem := range arr {
		h.Insert(elem)
	}
	return h
}

func (h *Heap) parentIndex(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func (h *Heap) leftIndex(i int) int {
	return 2*i + 1
}

func (h *Heap) rightIndex(i int) int {
	return 2*i + 2
}

func (h *Heap) Insert(x int) {
	h.data = append(h.data, x)
	parent := h.parentIndex(h.n)
	cur := h.n
	fmt.Printf("the index of the new item is %v; parent index is %v\n", h.n, parent)
	for parent >= 0 {
		if h.data[cur] <= h.data[parent] {
			break
		}
		h.data[parent], h.data[cur] = h.data[cur], h.data[parent]
		cur = parent
		parent = h.parentIndex(parent)
	}
	h.Print()
	h.n++
}

func (h *Heap) Delete() (int, error) {
	if h.n <= 0 {
		return 0, fmt.Errorf("the heap is empty")
	}

	result := h.data[0]
	h.data[0] = h.data[h.n-1]
	h.n--
	h.data = h.data[:h.n]
	cur := 0
	for {
		left, right := h.leftIndex(cur), h.rightIndex(cur)
		if left >= h.n && right >= h.n {
			break
		}
		var biggerChild int
		if left < h.n && right < h.n {
			biggerChild = left
			if h.data[right] > h.data[left] {
				biggerChild = right
			}

		} else if left < h.n && right >= h.n {
			biggerChild = left

		} else {
			biggerChild = right
		}
		if h.data[biggerChild] <= h.data[cur] {
			break
		}
		h.data[biggerChild], h.data[cur] = h.data[cur], h.data[biggerChild]
		cur = biggerChild
	}

	return result, nil
}

func (h *Heap) Print() {
	fmt.Println(h.data)
}

func (h *Heap) Count() int {
	return h.n
}
