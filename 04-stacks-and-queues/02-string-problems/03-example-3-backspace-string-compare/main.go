package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ab#c"
	t := "ab#c"

	ans := backspaceCompare(s, t)
	fmt.Println(ans)
}

func backspaceCompare(s string, t string) bool {
	build := func(str string) string {
		stack := []string{}
		for _, r := range str {
			c := string(r)
			if c != "#" {
				stack = append(stack, c)
				continue
			}
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		}

		return strings.Join(stack, "")
	}

	return build(s) == build(t)
}
