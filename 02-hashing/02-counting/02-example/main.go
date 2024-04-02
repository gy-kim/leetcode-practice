package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}

	out := intersection(nums)
	fmt.Println(out)
}

func intersection(nums [][]int) []int {
	m := map[int]int{}
	ans := []int{}

	for _, arr := range nums {
		for _, a := range arr {
			cnt := 1
			if v, ok := m[a]; ok {
				cnt += v
			}
			m[a] = cnt
		}
	}

	length := len(nums)

	for k, v := range m {
		if v == length {
			ans = append(ans, k)
		}
	}

	sort.Ints(ans)
	return ans
}
