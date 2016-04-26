package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func SwapBits(x int, a uint, b uint) (int, error) {
	w := uint(unsafe.Sizeof(x)) * 8
	if a > (w-1) || b > (w-1) {
		fmt.Printf("The bit index should be [0, %d]!\n", w - 1)
		return x, errors.New("index of bit is out of range")
	}

	if (x>>a & 0x01) !=	(x>>b & 0x01) {
		x ^= 0x01<<a
		x ^= 0x01<<b
	}
	return x, nil
}

func SwapBitsTest(x int, a uint, b uint) {
	if y, err := SwapBits(x, a, b); err == nil {
		fmt.Printf("SwapBits(%b, %d, %d) = %b\n", x, a, b, y)
	}
}

func main() {
	SwapBitsTest(2344354, 3, 32);
	SwapBitsTest(2344354, 1, 2);
	SwapBitsTest(2344354, 1, 200);
}
