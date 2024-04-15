package main

import "fmt"

func main() {
	// jewels := ""
	// stones := ""

	// jewels := ""
	// stones := ""

	jewels := ""
	stones := ""

	ans := numJewelsInStones(jewels, stones)
	fmt.Println(ans)
}

func numJewelsInStones(jewels string, stones string) int {
	jmap := map[rune]int{}

	for _, j := range jewels {
		jmap[j] = 1
	}

	ans := 0
	for _, s := range stones {
		ans += jmap[s]
	}

	return ans
}
