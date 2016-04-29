package main

import (
	"fmt"
	"unsafe"
)

func findFriend(x int) int {
	w := uint(unsafe.Sizeof(x)) * 8
	var i uint
	for i=0; i<(w-1); i++ {
		if (x>>i) & 0x01 != (x>>(i+1)) & 0x01 {
			x ^= (0x01<<i)
			x ^= (0x01<<(i+1))
			return x
		}
	}
	return x
}

func findFriendTest(x int) {
	fmt.Printf("findFriend(%064b) = %064b\n", x, findFriend(x))
}

func main() {
	findFriendTest(0)
	findFriendTest(-1)
	findFriendTest(135465656)
	findFriendTest(0xFFFFF)
}
