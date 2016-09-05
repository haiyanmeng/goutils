package main

import "fmt"

/*
1 0 0
0 0 1
0 1 0

0 1 0
0 0 0
1 0 1

0 0 1
1 0 0
0 1 0

0 0 1 0
1 0 0 0
0 0 0 1
0 1 0 0
*/

func AllPlacements(n int) [][][]int {
	if n < 4 {
		return [][][]int{}
	}

	if n == 4 {
		return [][][]int{
			[][]int{
				[]int{0, 1, 0, 0},
				[]int{0, 0, 0, 1},
				[]int{1, 0, 0, 0},
				[]int{0, 0, 1, 0},
			},
			[][]int{
				[]int{0, 0, 1, 0},
				[]int{1, 0, 0, 0},
				[]int{0, 0, 0, 1},
				[]int{0, 1, 0, 0},
			},
		}
	}

	result := [][][]int{}
	r1 := AllPlacements(n - 1)
	for _, placement := range r1 {
		n1 := hasQueenOnDia1(placement)
		n2 := hasQueenOnDia2(placement)
		if !n1 {
			result = append(result, addUpperLeft(placement))
			result = append(result, addLowerRight(placement))
		}

		if !n2 {
			result = append(result, addLowerLeft(placement))
			result = append(result, addUpperRight(placement))
		}
	}
	return result
}

func addLowerRight(a [][]int) [][]int {
	n := len(a)
	b := [][]int{}
	for i := 0; i < n; i++ {
		b = append(b, []int{})
		b[i] = append(b[i], a[i]...)
		b[i] = append(b[i], 0)
	}

	b = append(b, []int{})

	for i := 0; i < n; i++ {
		b[n] = append(b[n], 0)
	}

	b[n] = append(b[n], 1)

	return b
}

func addUpperRight(a [][]int) [][]int {
	n := len(a)
	b := [][]int{}
	b = append(b, []int{})
	for i := 0; i < n; i++ {
		b[0] = append(b[0], 0)

		b = append(b, []int{})
		b[i+1] = append(b[i+1], a[i]...)
		b[i+1] = append(b[i+1], 0)
	}
	b[0] = append(b[0], 1)

	return b
}

func addLowerLeft(a [][]int) [][]int {
	n := len(a)
	b := [][]int{}
	for i := 0; i < n; i++ {
		b = append(b, []int{0})
		b[i] = append(b[i], a[i]...)
	}
	b = append(b, []int{1})
	for i := 0; i < n; i++ {
		b[n] = append(b[n], 0)
	}
	return b
}

func addUpperLeft(a [][]int) [][]int {
	n := len(a)
	b := [][]int{}
	b = append(b, []int{1})
	for i := 0; i < n; i++ {
		b[0] = append(b[0], 0)

		b = append(b, []int{0})
		b[i+1] = append(b[i+1], a[i]...)
	}

	return b
}

func hasQueenOnDia2(a [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i][n-i-1] == 1 {
			return true
		}
	}
	return false
}

func hasQueenOnDia1(a [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i][i] == 1 {
			return true
		}
	}
	return false
}

/*
upper left corner: there are 2*n-1 new locations.
for each locations, we can test whether the Nth queen can be put there or not.
however, among these 2*n-1 new locations, there is only one location which may not cause attacking problems.
how you judge whether putting a queen into a location will cause attacking problems or not.
then we only need the diagonal of each (n-1)*(n-1) placement, whether there is a queen or not.
0 0 1 0
1 0 0 0
0 0 0 1
0 1 0 0
        1

upper right corner:
0 0 1 0
1 0 0 0
0 0 0 1
0 1 0 0

lower left corner:
0 0 1 0
1 0 0 0
0 0 0 1
0 1 0 0

lower right corner:
0 0 1 0
1 0 0 0
0 0 0 1
0 1 0 0


*/

func main() {
	n := 5
	r := AllPlacements(n)
	for i, placement := range r {
		fmt.Printf("The %dth placement:\n", i+1)
		for j := 0; j < n; j++ {
			fmt.Println(placement[j])
		}
	}
}
