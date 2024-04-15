package main

import (
	"fmt"
	"sort"
)

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}

	ans := findWinnders(matches)
	fmt.Println(ans)
}

func findWinnders(matches [][]int) [][]int {
	/*
		ansNotLost := []int{}
		ansLostOneMatch := []int{}

		winMatch := map[int]int{}
		lostMatch := map[int]int{}

		for _, match := range matches {
			c := 0
			win := match[0]
			if cnt, ok := winMatch[win]; ok {
				c = cnt
			}
			c++
			winMatch[win] = c

			c = 0
			lost := match[1]
			if cnt, ok := lostMatch[lost]; ok {
				c = cnt
			}

			c++
			lostMatch[lost] = c
		}
		for k, v := range lostMatch {
			if v == 1 {
				ansLostOneMatch = append(ansLostOneMatch, k)
			}
			delete(winMatch, k)
		}

		for k, _ := range winMatch {
			ansNotLost = append(ansNotLost, k)
		}

		sort.Ints(ansNotLost)
		sort.Ints(ansLostOneMatch)

		return [][]int{ansNotLost, ansLostOneMatch}
	*/

	countLosses := map[int]int{}
	ansOnlyWin := []int{}
	ansOneLoss := []int{}
	for _, match := range matches {
		win := match[0]
		lost := match[1]
		if _, ok := countLosses[win]; !ok {
			countLosses[win] = 0
		}

		cnt := 1
		if v, ok := countLosses[lost]; ok {
			cnt += v
		}
		countLosses[lost] = cnt
	}

	for k, v := range countLosses {
		if v == 0 {
			ansOnlyWin = append(ansOnlyWin, k)
		} else if v == 1 {
			ansOneLoss = append(ansOneLoss, k)
		}
	}

	sort.Ints(ansOnlyWin)
	sort.Ints(ansOneLoss)

	return [][]int{ansOnlyWin, ansOneLoss}
}
