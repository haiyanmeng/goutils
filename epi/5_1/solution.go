package main

import (
	"fmt"
	"sync"
)

// time complexity O(n), n is the word width in bits
func ComputeParity(x uint64) (r bool) { // the (r bool) must be surrounded by a pair of parenthess
	var i uint
	for i=0; i<64; i++ {
		if x & (0x01<<i) != 0 {
			if r {
				r = false
			} else {
				r = true
			}
		}
	}
	return r
}

// time complexity is O(d), d is the number of the bits which are set to be 1
func ComputeParity1(x uint64) (r bool) {
	var c uint = 0
	for x != 0 {
		x &= (x-1)
		c += 1
	}

	if c%2 == 0 {
		return false
	} else {
		return true
	}
}


// function arguments with the type of channel and all the reference types are passed by pointer
// function arguments with the type of struct and all the non-reference types are passed by value	
func ComputeParityInt16(x uint16, result chan int, wg *sync.WaitGroup) { // sync.WaitGroup is a struct
	defer wg.Done()
	//fmt.Printf("go routine address of wg = %p\n", wg)
	var i uint
	var r int
	for i=0; i<16; i++ {
		if x & (0x01<<i) != 0 {
			r++
		}
	}
	result <- r
}

func ComputeParity2(x uint64) (r bool) {
	var i uint

	// keeping result and wg as local variables allows ComputeParity2 to be called multiple times.
	// if result and wg are global variables, trying to call ComputeParity2 multiple times will have problems. because they all share the same wg and result and get messy.
	var result chan int = make(chan int)
	var wg sync.WaitGroup
	//fmt.Printf("address of wg = %p\n", &wg)

	for i=0; i<4; i++ {
		wg.Add(1)
		go ComputeParityInt16(uint16(x>>(i*16)), result, &wg)
	}

	go func() {
		wg.Wait() // using WaitGroup to wait for all the ComputeParityInt16 goroutines finishing
		close(result) // the channel result must be closed to avoid deadlock.
	}()

	var c int

	// collecting data on the channel result and closing the channel result must be in two different goroutines to avoid deadlock.
	for k := range result {
		c += k
	}

	if c%2 == 0 {
		return false
	} else {
		return true
	}
}

type FuncTemp func(x uint64) (r bool)

func test(f FuncTemp) {
	var x uint64
	x = 0x0b
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))
	
	x = 0x1b
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))

	x = 0x3FFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))

	x = 0x7FFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))

	x = 0xFFFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))
}

func main() {
	test(ComputeParity)
	test(ComputeParity1)
	test(ComputeParity2)
}
