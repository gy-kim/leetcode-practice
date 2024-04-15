package main

import (
	"fmt"
)

func main() {
	// grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	// grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	grid := [][]int{{11, 1}, {1, 11}}
	ans := equalPairs(grid)

	fmt.Println(ans)
}

func equalPairs(grid [][]int) int {
	rowMap := map[string]int{}
	colMap := map[string]int{}

	convertKey := func(arr []int) string {
		key := ""
		for _, v := range arr {
			key += fmt.Sprintf("%d,", v)
		}

		return key
	}

	size := len(grid)
	for row := 0; row < size; row++ {
		key := convertKey(grid[row])
		cnt := rowMap[key]
		cnt += 1
		rowMap[key] = cnt
	}

	for col := 0; col < size; col++ {
		colArr := []int{}
		for row := 0; row < size; row++ {
			colArr = append(colArr, grid[row][col])
		}
		key := convertKey(colArr)
		cnt := colMap[key]
		cnt += 1
		colMap[key] = cnt
	}

	ans := 0
	for k, rowCnt := range rowMap {
		colCnt := colMap[k]
		ans += rowCnt * colCnt
	}

	return ans
}
