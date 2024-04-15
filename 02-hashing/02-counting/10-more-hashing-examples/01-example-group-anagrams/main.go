package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)
}

func groupAnagrams(strs []string) any {
	m := map[string][]string{}

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)

		sortedStr := strings.Join(chars, "")
		arr := m[sortedStr]
		arr = append(arr, str)

		m[sortedStr] = arr
	}

	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}
