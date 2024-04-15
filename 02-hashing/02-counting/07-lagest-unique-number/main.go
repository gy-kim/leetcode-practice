package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 7, 3, 9, 4, 9, 8, 3, 1}
	// nums := []int{9, 9, 8, 8}
	ans := largestUniqueNumber(nums)
	fmt.Println(ans)
}

func largestUniqueNumber(nums []int) int {
	counts := map[int]int{}
	onlyOnce := map[int]struct{}{}
	for _, num := range nums {
		cnt := 1
		if v, ok := counts[num]; ok {
			cnt += v
			delete(onlyOnce, num)
		} else {
			onlyOnce[num] = struct{}{}
		}

		counts[num] = cnt
	}

	if len(onlyOnce) == 0 {
		return -1
	}

	ans := -1
	for k := range onlyOnce {
		if k > ans {
			ans = k
		}
	}

	return ans
}
