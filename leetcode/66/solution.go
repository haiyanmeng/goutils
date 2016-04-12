package main

import (
	"fmt"
	"errors"
)

func plusOne(a []int8) ([]int8, error) {
	fmt.Printf("%v\n", a)
	n := len(a)
	if n == 0 {
		return nil, errors.New("The list is empty!")
	}

	if a[n-1] < 9 {
		a[n-1]++
	} else {
		var i int
		for i=n-1; i>=0; i-- {
			if a[i] == 9 {
				a[i] = 0
			} else {
				a[i]++
			}
		}
		if i < 0 && a[0] == 0 {
			a = append([]int8{1}, a...)		
		}
	}

	return a, nil
}

func plusOneTest(a []int8) {
	if a, err := plusOne(a); err == nil {
		fmt.Printf("%v\n", a)
	} else {
		fmt.Println(err)
	}
	fmt.Println()
}

func main() {
	plusOneTest([]int8{})
	plusOneTest([]int8{4,2,3,9})
	plusOneTest([]int8{3,9,9,9})
	plusOneTest([]int8{9,9,9,9,9})
	plusOneTest([]int8{9})
}
