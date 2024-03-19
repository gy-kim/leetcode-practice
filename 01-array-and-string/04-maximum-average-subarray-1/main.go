package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	out := findMaxAverage(nums, k)
	fmt.Println(out)
}

func findMaxAverage(nums []int, k int) float64 {
	curr := 0

	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans := float64(curr) / float64(k)

	for i := k; i < len(nums); i++ {
		curr = curr + nums[i] - nums[i-k]
		tmp := float64(curr) / float64(k)
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}
