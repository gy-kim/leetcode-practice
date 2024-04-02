package main

import (
	"fmt"
)

func main() {
	// s := "ab-cd"
	// s := "a-bC-dEf-ghIj"
	// s := "Test1ng-Leet=code-Q!"
	s := "7_28]"

	out := reverseOnlyLetters(s)
	fmt.Println(out)
}
func reverseOnlyLetters(s string) string {
	if len(s) <= 1 {
		return s
	}

	i := 0
	j := len(s) - 1

	out := make([]byte, len(s))

	isLetter := func(b byte) bool {
		if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') {
			return false
		}

		return true
	}

	for i <= j {
		for !isLetter(s[i]) {
			out[i] = s[i]
			i++
		}
		for !isLetter(s[j]) {
			out[j] = s[j]
			j--
		}

		if i == j {
			out[i] = s[i]
		}

		out[i] = s[j]
		out[j] = s[i]

		i++
		j--

	}

	return string(out)
}
