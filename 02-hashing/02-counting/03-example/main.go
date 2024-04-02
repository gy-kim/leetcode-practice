package main

import "fmt"

func main() {
	// s := "abacbc"
	s := "aaabb"
	out := areOccurencesEqual(s)
	fmt.Println(out)
}

func areOccurencesEqual(s string) bool {
	m := map[rune]int{}

	for _, c := range s {
		cnt := 1
		if v, ok := m[c]; ok {
			cnt += v
		}
		m[c] = cnt
	}

	frequency := 0

	for _, v := range m {
		if frequency == 0 {
			frequency = v
			continue
		}
		if frequency != v {
			return false
		}
	}
	return true
}
