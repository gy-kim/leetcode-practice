package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	out := reverseWords(s)
	fmt.Println(out)
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	out := make([]string, 0, len(arr))

	for _, s := range arr {
		sub := make([]byte, len(s))
		i := 0
		j := len(s) - 1
		for i < j {
			sub[i] = s[j]
			sub[j] = s[i]
			i++
			j--
			if i == j {
				sub[i] = s[i]
				break
			}
		}

		out = append(out, string(sub))
	}

	return strings.Join(out, " ")
}
