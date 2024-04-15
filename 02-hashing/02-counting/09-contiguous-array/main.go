package main

import "fmt"

func main() {
	nums := []int{0, 1}
	// nums := []int{0, 1, 0}

	ans := findMaxLength(nums)
	fmt.Println(ans)
}

func findMaxLength(nums []int) int {
	// maxLen := 0
	// zero, one := 0, 0

	// for left := 0; left < len(nums); left++ {
	// 	for right := left; right < len(nums); right++ {
	// 		if nums[right] == 0 {
	// 			zero += 1
	// 		} else {
	// 			one += 1
	// 		}
	// 		if zero == one {
	// 			size := right - left + 1
	// 			if maxLen < size {
	// 				maxLen = size
	// 			}
	// 		}
	// 	}
	// }

	// return maxLen

	m := map[int]int{ // map[count]index
		0: -1,
	}
	maxLen, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count += 1
		} else {
			count += -1
		}
		if cnt, ok := m[count]; ok {
			if maxLen < (i - cnt) {
				maxLen = i - cnt
			}
		} else {
			m[count] = i
		}
	}

	return maxLen
}
