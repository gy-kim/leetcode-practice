package main

import "fmt"

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3

	// nums := []int{100000}
	// k := 0

	// nums := []int{8}
	// k := 100000

	out := getAverages(nums, k)
	fmt.Println(out)
}

func getAverages(nums []int, k int) []int {
	out := make([]int, len(nums))
	sum := make([]int, len(nums))
	defaultValue := -1
	windowSize := k*2 + 1

	if len(nums) == 0 {
		return out
	}

	out[0] = defaultValue
	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		out[i] = defaultValue
		sum[i] = sum[i-1] + nums[i]
	}

	if len(nums) < windowSize {
		return out
	}

	for i := 0; i < len(nums); i++ {
		if i < k {
			continue
		}

		rightIndex := i + k
		if rightIndex+1 > len(nums) {
			break
		}
		right := sum[rightIndex]

		left := 0
		leftIndex := i - k - 1
		if leftIndex >= 0 {
			left = sum[leftIndex]
		}

		out[i] = (right - left) / windowSize
	}

	return out
}
