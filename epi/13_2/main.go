package main

import "fmt"

/*
hello  good morning

helloo good     morning
count each char in the letter;
and then scan the text, and compare whether each character in the text appears more than in the letter.
*/

func writable(letter, text string) bool {
	map_letter := map[rune]int{}
	for _, c := range letter {
		map_letter[c]++
	}

	for _, c := range text {
		if _, ok := map_letter[c]; !ok {
			continue
		}

		map_letter[c]--

		if map_letter[c] == 0 {
			delete(map_letter, c)
		}

		if len(map_letter) == 0 {
			return true
		}
	}

	return false
}

func main() {
	letter := "hello  good morning !! yyyyy"
	text := "good olleh mornining !! "
	fmt.Printf("`%s` can be writable from `%s`? %v\n", letter, text, writable(letter, text))
}
