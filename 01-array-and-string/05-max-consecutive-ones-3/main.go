package main

import "fmt"

func main() {
	// nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	// k := 3

	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2

	out := longest(nums, k)
	fmt.Println(out)
}

func longest(nums []int, k int) int {
	ans, left := 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}

		if k < 0 {
			k += 1 - nums[left]
			left++
		}
		size := right - left + 1
		if size > ans {
			ans = size
		}
	}

	return ans
}
