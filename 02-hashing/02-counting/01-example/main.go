package main

import "fmt"

func main() {
	s := "eceba"
	k := 2
	out := findLongestSubstring(s, k)
	fmt.Println(out)
}

func findLongestSubstring(s string, k int) int {
	left, ans := 0, 0
	m := map[string]int{}

	for right := 0; right < len(s); right++ {
		c := string(s[right])
		if v1, ok := m[c]; ok {
			v1 += 1
			m[c] = v1
		} else {
			m[c] = 1
		}

		for len(m) > k {
			leftChar := string(s[left])
			if v2, ok := m[leftChar]; ok {
				v2 -= 1
				m[leftChar] = v2
				if v2 == 0 {
					delete(m, leftChar)
				}
				left += 1
			}
		}

		curLengh := right - left + 1
		if curLengh > ans {
			ans = curLengh
		}
	}

	return ans
}
