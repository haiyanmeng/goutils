package main

import (
	"fmt"
	"unsafe"
)

func listPrimes(n int) []int {
	var r []int
	var arr []bool= make([]bool, n+1)

	for i:=0; i<=n; i++ {
		arr[i] = true; //true denotes it is a prime; false denotes it is not a prime.
	}

	arr[0] = false
	arr[1] = false
	k := 2
	for k<=n {
		for k <= n && arr[k] == false {
			k++
		}

		if k > n {
			break
		}
			
		r = append(r, k)

		if k <= n {
			i := k
			//i := 2
			for	k * i <= n {
				arr[k*i] = false 
				i++
			}
			k++
		}
	}

	return r
}

func main() {
	a := 18
	fmt.Printf("listPrimes(%d) = %v\n", a, listPrimes(a))
	a = 180
	fmt.Printf("listPrimes(%d) = %v\n", a, listPrimes(a))
}
