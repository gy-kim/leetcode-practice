package main

import "fmt"

func main() {
	s := "1101100111"

	out := findLenth(s)
	fmt.Print(out)
}

func findLenth(s string) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var left, curr, ans int = 0, 0, 0

	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			curr++
		}
		for curr > 1 {
			if s[left] == '0' {
				curr--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
