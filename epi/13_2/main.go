package main

import "fmt"

/*
hello  good morning

helloo good     morning
count each char in the letter;
and then scan the text, and compare whether each character in the text appears more than in the letter.
*/

func writable(letter, text string) bool {
	map_letter, map_text := map[rune]int{}, map[rune]int{}
	for _, c := range letter {
		map_letter[c]++
	}
	for _, c := range text {
		map_text[c]++
	}

	for k, v := range map_letter {
		if map_text[k] < v {
			return false
		}
	}
	return true
}

func main() {
	letter := "hello  good morning !! yyyyy"
	text := "good olleh mornining !! "
	fmt.Printf("`%s` can be writable from `%s`? %v\n", letter, text, writable(letter, text))
}
