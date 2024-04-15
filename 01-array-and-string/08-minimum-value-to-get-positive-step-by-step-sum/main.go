package main

import "fmt"

func main() {
	// nums := []int{-3, 2, -3, 4, 2}
	// nums := []int{1, 2}
	nums := []int{1, -2, -3}
	out := minStartValue(nums)
	fmt.Println(out)
}
func minStartValue(nums []int) int {
	ans := 1
	sum := 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 1 {
			gap := (sum - 1) * -1
			ans += gap
			sum = 1
		}
	}

	return ans
	// out := 1
	// sum := 1
	// for i := 0; i < len(nums); i++ {
	// 	sum += nums[i]
	// 	if sum < 1 {
	// 		gap := (sum - 1) * -1
	// 		sum = 1
	// 		out += gap
	// 	}
	// }

	// return out
}
