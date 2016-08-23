package main

import "fmt"

type Elem struct {
	x      int
	listID int
}

type Heap struct {
	Data []Elem
	n    int
}

func New() *Heap {
	return &Heap{}
}

func (h *Heap) parentIndex(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func (h *Heap) Insert(d int, listID int) {
	e := Elem{d, listID}
	h.Data = append(h.Data, e)
	h.n++
	cur := h.n - 1
	parent := h.parentIndex(cur)
	for parent >= 0 {
		if h.Data[cur].x >= h.Data[parent].x {
			return
		}
		h.Data[cur], h.Data[parent] = h.Data[parent], h.Data[cur]
		cur = parent
		parent = h.parentIndex(parent)
	}
}

func (h *Heap) Delete() (Elem, error) {
	if h.n == 0 {
		return Elem{}, fmt.Errorf("the heap is empty")
	}

	e := h.Data[0]
	h.Data[0] = h.Data[h.n-1]
	h.n--
	h.Data = h.Data[:h.n]

	cur := 0
	for {
		leftIndex, rightIndex := cur*2+1, cur*2+2
		if leftIndex >= h.n && rightIndex >= h.n {
			break
		}

		var smallerChild int
		if leftIndex < h.n && rightIndex < h.n {
			smallerChild = leftIndex
			if h.Data[rightIndex].x > h.Data[leftIndex].x {
				smallerChild = rightIndex
			}
		} else if leftIndex < h.n && rightIndex >= h.n {
			smallerChild = leftIndex
		} else {
			smallerChild = rightIndex
		}

		if h.Data[smallerChild].x >= h.Data[cur].x {
			break
		}

		h.Data[smallerChild], h.Data[cur] = h.Data[cur], h.Data[smallerChild]
		cur = smallerChild
	}

	return e, nil
}

func mergeLists(lists [][]int) []int {
	h := New()
	listLens := []int{}
	curIndices := []int{}
	for i, list := range lists {
		n := len(list)
		listLens = append(listLens, n)
		curIndices = append(curIndices, 0)
		if n == 0 {
			continue
		}
		h.Insert(list[0], i)
	}

	result := []int{}
	for {
		elem, err := h.Delete()
		if err != nil {
			return result
		}

		result = append(result, elem.x)
		curIndices[elem.listID]++
		if curIndices[elem.listID] >= listLens[elem.listID] {
			continue
		}
		h.Insert(lists[elem.listID][curIndices[elem.listID]], elem.listID)
	}
}

func main() {
	lists := [][]int{
		[]int{3, 5, 7},
		[]int{0, 6, 7, 8},
		[]int{0, 6, 10, 24, 28},
	}
	fmt.Println(lists)
	fmt.Println(mergeLists(lists))
}
