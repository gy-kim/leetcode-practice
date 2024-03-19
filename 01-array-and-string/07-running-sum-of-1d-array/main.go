package main

import "fmt"

func main() {
	// input := []int{1, 2, 3, 4}
	// input := []int{1, 1, 1, 1, 1}
	input := []int{3, 1, 2, 10, 1}
	out := runningSum(input)
	fmt.Println(out)
}

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	out := make([]int, len(nums))
	out[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		out[i] = out[i-1] + nums[i]
	}

	return out
}
