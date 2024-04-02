package main

import "fmt"

func main() {
	s := "abcdeda"

	out := repeatedCharacter(s)
	fmt.Println(out)
}

func repeatedCharacter(s string) string {
	m := map[byte]struct{}{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := m[c]; ok {
			return string(c)
		}

		m[c] = struct{}{}
	}

	return ""
}
