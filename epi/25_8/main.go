package main

import "fmt"

func justify(words []string, target int) ([]string, error) {
	var result []string
	count := len(words)
	cur := 0   // the current length of the line
	start := 0 // the start index of the first word for the line
	for i, word := range words {
		n := len(word)
		if n > target {
			return result, fmt.Errorf("the word (%s) is longer than the line limit", word)
		}
		if n >= target-cur {
			// the current word should be the start word of a new line
			line := createLine(words, start, i-1, target-cur, false)
			result = append(result, line)
			cur, start = 0, i
		}
		if cur == 0 {
			cur += n
		} else {
			cur += n + 1 // 1 for the blank between two words
		}
	}

	// process the last line
	line := createLine(words, start, count-1, target-cur, true)
	result = append(result, line)
	return result, nil
}

func createLine(words []string, start, end int, extraSpace int, spaceEnd bool) string {
	var line string
	line = words[start]
	if !spaceEnd {
		wordNum := end - start
		totalSpace := extraSpace + wordNum
		x1 := totalSpace / wordNum
		extra := totalSpace % wordNum
		for i := 0; i < extra; i++ {
			line += " "
		}

		spaceSep := ""
		for i := 0; i < x1; i++ {
			spaceSep += " "
		}
		for j := start + 1; j <= end; j++ {
			line = line + spaceSep + words[j]
		}

	} else {
		for j := start + 1; j <= end; j++ {
			line = line + " " + words[j]
		}

		for i := 0; i < extraSpace; i++ {
			line += " "
		}
	}
	return line
}

func main() {
	a := []string{"afdfdfdf", "hellodfdfwrold", "a", "bdd", "thequick", "rown", "fox", "jumped", "over", "the", "lazy", "dogs."}
	r, err := justify(a, 33)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("`%s`\n", r[i])
	}
}
