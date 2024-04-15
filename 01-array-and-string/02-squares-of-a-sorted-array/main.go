package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-7, -3, 2, 3, 11}

	out := sortedSquares(nums)

	fmt.Print(out)
}

// func sortedSquares(nums []int) []int {
// 	out := make([]int, len(nums))
// 	idx := len(nums) - 1

// 	for i, j := 0, len(nums)-1; i <= j; {
// 		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
// 			out[idx] = nums[i] * nums[i]
// 			i++
// 		} else {
// 			out[idx] = nums[j] * nums[j]
// 			j--
// 		}
// 		idx--
// 	}

// 	return out
// }

func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	idx := len(nums) - 1

	left := 0
	right := len(nums) - 1
	for left <= right {
		if math.Abs(float64(nums[left])) > math.Abs(float64(nums[right])) {
			out[idx] = nums[left] * nums[left]
			left++
		} else {
			out[idx] = nums[right] * nums[right]
			right--
		}

		idx--
	}
	return out
}
