package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100

	out := numSubArrayProductLessThanK(nums, k)
	fmt.Println(out)
}

func numSubArrayProductLessThanK(nums []int, k int) int {
	if k <= 0 {
		return 0
	}

	ans, left := 0, 0
	curr := 1

	for right := 0; right < len(nums); right++ {
		curr *= nums[right]
		for curr >= k {
			curr /= nums[left]
			left++
		}
		ans += right - left + 1
	}

	return ans
}
