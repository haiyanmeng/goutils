package main

import "fmt"

func isAlphaNum(b byte) bool {
	if b >= '0' && b <= '9' ||
		b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + b - 'A'
	}
	return b
}

func isPalindromicity(s string) bool {
	n := len(s)
	i := 0
	j := n - 1
	for i < j {
		for !isAlphaNum(s[i]) {
			i++
		}
		for !isAlphaNum(s[j]) {
			j--
		}

		if i < j {
			if toLower(s[i]) != toLower(s[j]) {
				return false
			}
			i++
			j--
		}

	}
	return true
}

func main() {
	var s string
	s = "Ray a Ray"
	fmt.Printf("isPalindromicity(%s) = %v\n", s, isPalindromicity(s))
	s = "Able was I, ere I saw Elba!!"
	fmt.Printf("isPalindromicity(%s) = %v\n", s, isPalindromicity(s))
	s = "A man, a plan, a canal, Panama."
	fmt.Printf("isPalindromicity(%s) = %v\n", s, isPalindromicity(s))

}
