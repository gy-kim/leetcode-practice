package main

import (
	"fmt"
	"strings"
)

func main() {
	// ransomNote := "aa"
	// magazine := "aabb"

	// ransomNote := ""
	// magazine := ""

	// ransomNote := "a"
	// magazine := "b"

	ransomNote := "aa"
	magazine := "ab"

	ans := canConstruct(ransomNote, magazine)
	fmt.Println(ans)
}

func canConstruct(ransomNote string, magazine string) bool {
	rchars := strings.Split(ransomNote, "")
	mchars := strings.Split(magazine, "")

	convertMap := func(arr []string) map[string]int {
		out := map[string]int{}
		for _, v := range arr {
			cnt := out[v]
			cnt += 1
			out[v] = cnt
		}
		return out
	}

	rmap := convertMap(rchars)
	mmap := convertMap(mchars)

	for k, v := range rmap {
		cnt, ok := mmap[k]
		if !ok {
			return false
		}
		if v > cnt {
			return false
		}

	}

	return true
}
