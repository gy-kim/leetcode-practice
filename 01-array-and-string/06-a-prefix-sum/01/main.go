package main

import "fmt"

type Query struct {
	Start int
	End   int
}

func main() {
	nums := []int{1, 6, 3, 2, 7, 2}
	queries := []Query{
		{0, 3},
		{2, 5},
		{2, 4},
	}
	limit := 13

	ans := answerQueries(nums, queries, limit)

	fmt.Println(ans)
}

func answerQueries(nums []int, queries []Query, limit int) []bool {
	prefix := make([]int, len(nums))
	ans := make([]bool, len(queries))

	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	for i := 0; i < len(queries); i++ {
		x := queries[i].Start
		y := queries[i].End

		ans[i] = (prefix[y] - prefix[x] + nums[x]) < limit
	}

	return ans
}
