package main

import "fmt"

func main() {
	// nums := []int{1, 2, 1, 2, 1}
	// k := 3
	nums := []int{1, -1, 1, -1}
	k := 0

	ans := subarraySum(nums, k)
	fmt.Println(ans)
}

func subarraySum(nums []int, k int) int {
	counts := map[int]int{}
	ans, curr := 0, 0
	counts[0] = 1

	for _, num := range nums {
		curr += num
		ans += counts[curr-k]

		cnt := 1
		if count, ok := counts[curr]; ok {
			cnt += count
		}
		counts[curr] = cnt

	}

	return ans
}
