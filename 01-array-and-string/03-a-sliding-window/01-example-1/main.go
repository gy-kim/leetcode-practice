package main

import (
	"fmt"
)

func main() {
	num := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	k := 8

	out := findLenth(num, k)
	fmt.Print(out)

}

func findLenth(nums []int, k int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	left, curr, ans := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		curr += nums[right]
		for curr > k {
			curr -= nums[left]
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
