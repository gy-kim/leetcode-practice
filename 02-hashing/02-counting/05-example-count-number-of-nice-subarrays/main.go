package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3

	ans := numberOfSubarrays(nums, k)
	fmt.Println(ans)
}

func numberOfSubarrays(nums []int, k int) int {
	counts := map[int]int{}
	counts[0] = 1

	ans, curr := 0, 0
	for _, num := range nums {
		curr += (num % 2)
		ans += counts[curr-k]

		cnt := 0
		if v, ok := counts[curr]; ok {
			cnt = v
		}
		cnt += 1
		counts[curr] = cnt
	}

	return ans
}
