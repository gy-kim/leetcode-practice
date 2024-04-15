package main

import "fmt"

func main() {
	// s := "abcabcbb"
	// s := "bbbbb"
	// s := "pwwkew"
	// s := ""
	// s := "au"
	// s := "pwwkew"
	// s := "aab"
	// s := "dvdf"
	s := "abba"

	ans := lengthOfLongestSubstring(s)
	fmt.Println(ans)
}

func lengthOfLongestSubstring(s string) int {
	m := map[string]int{}
	ans := 0
	left := 0

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for right := 0; right < len(s); right++ {
		char := string(s[right])

		if v, ok := m[char]; ok {
			left = max(left, v)
		}

		size := right - left + 1
		ans = max(size, ans)

		m[char] = right + 1

	}

	return ans
}
