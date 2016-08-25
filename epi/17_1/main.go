package main

import "fmt"

func count(target int, scores []int) [][]int {
	result := [][]int{}
	n := len(scores)
	if n == 0 {
		return result
	}

	first := scores[0]
	c := 0
	remain := target
	for remain > 0 {
		remain = target - first*c
		combinations := count(remain, scores[1:])
		if len(combinations) == 0 {
			if remain == 0 {
				l := []int{c}
				for i := 0; i < n-1; i++ {
					l = append(l, 0)
				}
				result = append(result, l)
			}
			c++
			continue
		}

		for _, list := range combinations {
			l := []int{c}
			l = append(l, list...)
			result = append(result, l)
		}
		c++
	}

	return result
}

func main() {
	fmt.Println(count(12, []int{2, 3, 7}))
	fmt.Println(count(12, []int{7, 2, 3}))
	fmt.Println(count(12, []int{2, 7, 3}))
}
