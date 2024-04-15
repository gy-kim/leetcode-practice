package main

import (
	"fmt"
	"math"
)

func main() {
	// cards := []int{1, 2, 6, 2, 1}
	cards := []int{1, 2, 3, 4, 5}
	ans := minimumCardPickup(cards)
	fmt.Println(ans)
}

func minimumCardPickup(cards []int) int {
	if len(cards) <= 1 {
		return -1
	}

	m := map[int][]int{}
	ans := math.MaxInt

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		var list []int
		ok := false
		if list, ok = m[card]; ok {
			size := i - list[len(list)-1] + 1
			if size < ans {
				ans = size
			}
		}
		list = append(list, i)
		m[card] = list
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}
