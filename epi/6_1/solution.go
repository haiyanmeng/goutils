package main

import (
	"fmt"
	"errors"
)

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func partition(a *[]int, i int) error {
	pivot := (*a)[i]
	var x, y int = 0, len(*a) - 1
	for x<y {
		for (*a)[x] < pivot {
			x++
		}
		for (*a)[y] >= pivot {
			y--
		}
		if x<y {
			swap(&((*a)[x]), &((*a)[y]))
			x++
			y--
		}
	}

	x, y = 0, len(*a) - 1

	for x<y {
		for (*a)[x] <= pivot {
			x++
		}

		for (*a)[y] > pivot {
			y--
		}
		if x<y {
			swap(&((*a)[x]), &((*a)[y]))
		}
		x++
		y--
	}

	return errors.New("error happens")
}

func partitionTest(a *[]int, i int) {
	fmt.Println(*a)
	partition(a, i)
	fmt.Println(*a)
}

func main() {
	a := []int{4, 9, 3, 1, 2}
	partitionTest(&a, 0)
	b := []int{4, 9, 3, 4, -1, 8, 4, 5, 0, 1, 2}
	partitionTest(&b, 0)
}
