package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abbaca"

	ans := removeDuplicates(s)
	fmt.Println(ans)
}

func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}
	stack := []string{}

	for _, r := range s {
		c := string(r)
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		if stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, c)
	}

	return strings.Join(stack, "")
}
