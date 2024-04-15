package main

import "fmt"

func main() {
	// text := "nlaebolko"
	// text := "loonbalxballpoon"
	// text := "leetcode"
	text := "nlaebolko"
	// text := "balon"

	ans := maxNumberOfBalloons(text)
	fmt.Println(ans)
}

func maxNumberOfBalloons(text string) int {
	if len(text) < 7 {
		return 0
	}
	ans := -1
	m := map[string]int{
		"b": 0,
		"a": 0,
		"l": 0,
		"o": 0,
		"n": 0,
	}

	for _, c := range text {
		char := string(c)
		if cnt, ok := m[char]; ok {
			cnt += 1
			m[char] = cnt
		}
	}

	for k, v := range m {
		if k == "l" || k == "o" {
			v = v / 2
		}

		if ans == -1 || ans > v {
			ans = v
		}
	}

	return ans
}
