package main

import (
	"fmt"
	"unsafe"
)

func reverseBits(x int) int {
	w := uint(unsafe.Sizeof(x)) * 8
	var i uint
	for i=0; i<w/2; i++ {
		// fmt.Println(i, w-i-1, x>>i & 0x01, x>>(w-i-1) & 0x01)
		if (x>>i & 0x01) != (x>>(w-i-1) & 0x01) {
			x ^= 0x01<<i
			x ^= 0x01<<(w-i-1)
			// fmt.Println("  ", i, w-i-1, x>>i & 0x01, x>>(w-i-1) & 0x01)
		}
	}
	return x
}

func reverseBitsTest(x int) {
	fmt.Printf("reverseBits(%064b) = %064b\n", x, reverseBits(x))
}

func main() {
	reverseBitsTest(0x34353436);
	reverseBitsTest(0xFE);
	reverseBitsTest(0xFF);
}
