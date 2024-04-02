package main

import "fmt"

func main() {
	// arr := []int{1, 2, 3}
	arr := []int{1, 1, 3, 3, 5, 5, 7, 7}
	// arr := []int{1, 1, 2, 2}
	out := countElements(arr)
	fmt.Println(out)
}

func countElements(arr []int) int {
	/*
		m := map[int]struct{}{}
		cnt := 0

		for _, n := range arr {
			m[n] = struct{}{}
		}

		for _, n := range arr {
			if _, ok := m[n+1]; ok {
				cnt++
			}
		}

		return cnt
	*/

	// If the arr is sorted, we can use this algorithm.
	cnt := 0
	runLength := 1

	for i := 1; i < len(arr); i++ {
		if arr[i-1] != arr[i] {
			if arr[i-1]+1 == arr[i] {
				cnt += runLength
			}
			runLength = 0
		}

		runLength++
	}

	return cnt
}
