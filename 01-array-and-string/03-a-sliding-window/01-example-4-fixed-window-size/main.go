package main

import "fmt"

func main() {
	nums := []int{3, -1, 4, 12, -8, 5, 6}
	k := 4

	out := findBestSubarray(nums, k)
	fmt.Println(out)
}

func findBestSubarray(nums []int, k int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	curr := 0
	ans := 0

	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans = curr
	for i := k; i < len(nums); i++ {
		curr = curr - nums[i-k] + nums[i]
		ans = max(ans, curr)
	}

	return ans

	// curr := 0
	// for i := 0; i < k; i++ {
	// 	curr += nums[i]
	// }

	// ans := curr
	// for i := k; i < len(nums); i++ {
	// 	curr = curr + nums[i] - nums[i-k]
	// }

	// if curr > ans {
	// 	ans = curr
	// }

	// return ans
}
