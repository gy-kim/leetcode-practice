package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{18, 43, 36, 13, 7}

	ans := maximumSum(nums)
	fmt.Println(ans)
}

func maximumSum(nums []int) int {
	ans := -1
	m := map[int][]int{}

	sum := func(num int) int {
		digitSum := 0
		for num > 0 {
			digitSum += num % 10
			num /= 10
		}

		return digitSum
	}

	for _, num := range nums {
		digitSum := sum(num)
		var arr []int
		var ok bool
		if arr, ok = m[digitSum]; ok {
			tmpSum := arr[len(arr)-1] + num
			if ans < tmpSum {
				ans = tmpSum
			}
		}

		arr = append(arr, num)
		sort.Ints(arr)
		m[digitSum] = arr
	}

	return ans
}
