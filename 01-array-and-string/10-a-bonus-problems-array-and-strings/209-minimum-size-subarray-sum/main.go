package main

import (
	"fmt"
)

/*
Given an array of positive integers nums and a positive integer target,
return the minimal length of a subarray whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.
*/
func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	ans := minSubArrayLen(target, nums)
	fmt.Println(ans)
}

func minSubArrayLen(target int, nums []int) int {
	ans := 0
	if len(nums) == 0 {
		return ans
	}

	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}

	left := 0
	sum := nums[0]

	if sum >= target {
		return 1
	}

	for right := 1; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if ans == 0 {
				ans = right - left + 1
			}
			ans = min(right-left+1, ans)
			sum -= nums[left]
			left++
		}
	}

	return ans
}
