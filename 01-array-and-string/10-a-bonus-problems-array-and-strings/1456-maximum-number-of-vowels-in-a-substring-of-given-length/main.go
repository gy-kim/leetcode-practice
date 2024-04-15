package main

import "fmt"

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.



Example 1:
Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.

Example 2:
Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.

Example 3:
Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.
*/

func main() {
	s := "abciiidef"
	k := 3

	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) int {

	/*
		vowels := map[byte]any{
			'a': nil,
			'e': nil,
			'i': nil,
			'o': nil,
			'u': nil,
		}

		max := 0
		for i := 0; i < k; i++ {
			if (len(s) < k) && i >= k {
				return max
			}
			if _, ok := vowels[s[i]]; ok {
				max++
			}
		}

		curr := max

		for i := k; i < len(s); i++ {
			if _, ok := vowels[s[i]]; ok {
				curr++
			}
			if _, ok := vowels[s[i-k]]; ok {
				curr--
			}

			if curr > max {
				max = curr
			}
		}

		return max
	*/

	vowels := map[byte]any{
		'a': nil,
		'e': nil,
		'i': nil,
		'o': nil,
		'u': nil,
	}

	ans := 0
	curr := 0

	for i := 0; i < len(s); i++ {
		if i < k {
			if _, ok := vowels[s[i]]; ok {
				ans++
			}

			curr = ans

			continue
		}

		if _, ok := vowels[s[i]]; ok {
			curr++
		}
		if _, ok := vowels[s[i-k]]; ok {
			curr--
		}

		if curr > ans {
			ans = curr
		}

	}

	return ans

}
