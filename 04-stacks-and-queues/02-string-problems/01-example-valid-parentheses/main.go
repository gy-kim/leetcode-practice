package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/
func main() {
	// s := "()"
	s := "{(}"

	ans := isValid(s)
	fmt.Println(ans)
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	chars := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		if _, ok := chars[c]; ok {
			stack = append(stack, c)
			continue
		}

		if chars[stack[len(stack)-1]] == c {
			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	return len(stack) == 0

}
