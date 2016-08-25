package main

import "fmt"

func count(target int, scores []int) int {
	m := len(scores)
	d := make([][]int, 0, m)
	for i, score := range scores {
		d = append(d, make([]int, target+1, target+1))
		for j := 0; j <= target; j++ {
			if i == 0 {
				if j%2 == 0 {
					d[i][j] = 1
				}
			} else {
				remain := j
				for remain >= 0 {
					d[i][j] += d[i-1][remain]
					remain -= score
				}
			}
		}
	}
	return d[m-1][target]
}

func main() {
	fmt.Println(count(12, []int{2, 3, 7}))
	fmt.Println(count(12, []int{7, 2, 3}))
	fmt.Println(count(12, []int{2, 7, 3}))
}
