package main

import "fmt"

func main() {
	nums := []int{5, 2, 7, 10, 3, 9}
	target := 8

	out := twoSum(nums, target)
	fmt.Println(out)
}

func twoSum(nums []int, target int) []int {
	hmap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		comp := target - num
		if val, ok := hmap[comp]; ok {
			return []int{val, i}
		}

		hmap[num] = i
	}

	return []int{}
}
