package leetcode

import "strings"

func countPalindromicSubsequence(s string) int {
	res := 0
	uniqueChars := make(map[rune]bool)
	for _, char := range s {
		uniqueChars[char] = true
	}
	for key, _ := range uniqueChars {
		start := strings.IndexRune(s, key)
		end := strings.LastIndex(s, string(key))

		if start < end {
			uniqueStrings := make(map[rune]bool)

			for _, char := range s[start+1 : end] {
				if !uniqueStrings[char] {
					uniqueStrings[char] = true
					res++
				}
			}
		}
	}
	return res
}
