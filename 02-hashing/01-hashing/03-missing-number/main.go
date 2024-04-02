package main

import "fmt"

func main() {
	// nums := []int{3, 0, 1}
	// nums := []int{0, 1}
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	out := missingNumber(nums)
	fmt.Println(out)
}

func missingNumber(nums []int) int {
	/*
		m := map[int]struct{}{}

		for _, num := range nums {
			m[num] = struct{}{}
		}

		for i := 0; i <= len(nums); i++ {
			if _, ok := m[i]; !ok {
				return i
			}
		}

		return -1

	*/

	// expectSum := 0
	// for i := 0; i <= len(nums); i++ {
	// 	expectSum += i
	// }
	expectSum := len(nums) * (len(nums) + 1) / 2

	actualSum := 0
	for _, n := range nums {
		actualSum += n
	}

	return expectSum - actualSum
}
